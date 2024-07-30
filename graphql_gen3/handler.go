/*
GraphQL Web endpoint
*/

package main

import (
	"fmt"
	"net/http"
    "sync"
    "time"
    "io"
    "encoding/json"
    "errors"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"github.com/graphql-go/handler"
)

 type UserAuth struct {
     ExpiresAt time.Time
     AuthorizedResources []any
 }

 type TokenCache struct {
     mu    sync.Mutex
     cache map[string]UserAuth
 }

 func NewTokenCache() *TokenCache {
     return &TokenCache{
         cache: make(map[string]UserAuth),
     }
 }

// handle the graphql queries for a single endpoint
type graphHandler struct {
	graph      string
	gqlHandler *handler.Handler
	timestamp  string
	client     gripql.Client
    tokenCache *TokenCache
	//schema     *gripql.Graph
}

// Handler is a GraphQL endpoint to query the Grip database
type Handler struct {
	handlers map[string]*graphHandler
	client   gripql.Client
}

type ServerError struct {
    StatusCode int
    Message string
}

func (e *ServerError) Error() string {
    return e.Message
}

func getAuthMappings(url string, token string) (any, error) {
     GetRequest, err := http.NewRequest("GET", url, nil)
     if err != nil {
         log.Error(err)
         return nil, err
     }

     GetRequest.Header.Set("Authorization", token)
     GetRequest.Header.Set("Accept", "application/json")
     fetchedData, err := http.DefaultClient.Do(GetRequest)
     if err != nil {
         log.Error(err)
         return nil, err
     }
     defer fetchedData.Body.Close()

     if fetchedData.StatusCode == http.StatusOK {
         bodyBytes, err := io.ReadAll(fetchedData.Body)
         if err != nil {
             log.Error(err)
         }

         var parsedData any
         err = json.Unmarshal(bodyBytes, &parsedData)
         if err != nil {
             log.Error(err)
             return nil, err
         }
         return parsedData, nil

     }
     // code must be nonNull to get here, probably don't want to cache a failed state
     empty_map :=  make(map[string]any)
     err = errors.New("Arborist auth/mapping GET returned a non-200 status code: " + fetchedData.Status)
     return empty_map, err
 }

 func hasPermission(permissions []any) bool {
     for _, permission := range permissions {
         permission := permission.(map[string]any)
         if (permission["service"] == "*" || permission["service"] == "peregrine") &&
             (permission["method"] == "*" || permission["method"] == "read") {
             // fmt.Println("PERMISSIONS: ", permission)
             return true
         }
     }
     return false
 }

 func getAllowedProjects(url string, token string) ([]any, error) {
     var readAccessResources []string
     authMappings, err := getAuthMappings(url, token)
     if err != nil {
         return nil, err
     }

     // Iterate through /auth/mapping resultant dict checking for valid read permissions
     for resourcePath, permissions := range authMappings.(map[string]any) {
         if hasPermission(permissions.([]any)) {
             readAccessResources = append(readAccessResources, resourcePath)
         }
     }

     s := make([]interface{}, len(readAccessResources))
     for i, v := range readAccessResources {
         s[i] = v
     }
     return s, nil
 }

 func handleError(err error, writer http.ResponseWriter) {
    if ae, ok := err.(*ServerError); ok {
        response := ServerError{StatusCode: ae.StatusCode, Message: ae.Message}
        jsonResponse, _ := json.Marshal(response)
        writer.WriteHeader(ae.StatusCode)
        writer.Write(jsonResponse)
    }else {
        response := ServerError{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("General error occured while setting up graphql handler")}
        jsonResponse, _ := json.Marshal(response)
        writer.WriteHeader(http.StatusInternalServerError)
        writer.Write(jsonResponse)
    }
}

// NewClientHTTPHandler initilizes a new GraphQLHandler
func NewHTTPHandler(client gripql.Client, config map[string]string) (http.Handler, error) {
	h := &Handler{
		client:   client,
		handlers: map[string]*graphHandler{},
	}
	return h, nil
}

// Static HTML that links to Apollo GraphQL query editor
var sandBox = `
<div id="sandbox" style="position:absolute;top:0;right:0;bottom:0;left:0"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
 new window.EmbeddedSandbox({
   target: "#sandbox",
   // Pass through your server href if you are embedding on an endpoint.
   // Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
   initialEndpoint: window.location.href,
 });
 // advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
</script>`

// ServeHTTP responds to HTTP graphql requests
func (gh *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//log.Infof("Request for %s", request.URL.Path)
	//If no graph provided, return the Query Editor page
	if request.URL.Path == "" || request.URL.Path == "/" {
		writer.Write([]byte(sandBox))
		return
	}
	//pathRE := regexp.MustCompile("/(.+)$")
	//graphName := pathRE.FindStringSubmatch(request.URL.Path)[1]
	graphName := request.URL.Path
	var handler *graphHandler
	var ok bool
	if handler, ok = gh.handlers[graphName]; ok {
		//Call the setup function. If nothing has changed it will return without doing anything
		err := handler.setup(request.Header)
        if err != nil{
            handleError(err, writer)
            return
        }
	} else {
        tokenCache := NewTokenCache()
		//Graph handler was not found, so we'll need to set it up
		var err error
		handler, err = newGraphHandler(graphName, gh.client, request.Header, tokenCache)
        if err != nil{
            handleError(err, writer)
            return
        }
		gh.handlers[graphName] = handler
	}
	if handler != nil && handler.gqlHandler != nil {
		handler.gqlHandler.ServeHTTP(writer, request)
	} else {
         response := ServerError{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("General error occured while setting up graphql handler")}
         jsonResponse, _ := json.Marshal(response)
         writer.Write(jsonResponse)
	}
}

// newGraphHandler creates a new graphql handler from schema
func newGraphHandler(graph string, client gripql.Client, headers http.Header, userCache *TokenCache) (*graphHandler, error) {
	o := &graphHandler{
		graph:  graph,
		client: client,
        tokenCache: userCache,
	}
	err := o.setup(headers)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// LookupToken looks up a user token in the cache based on the token string.
// but resource lists can change when users are given new permissions so it's probably better not to cache resourceLists
func (tc *TokenCache) LookupResourceList(token string) ([]any) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	auth, _:= tc.cache[token]
    var resourceList []any
    if auth.AuthorizedResources != nil {
        resourceList = auth.AuthorizedResources
    }
	return resourceList
}

// Check timestamp to see if schema needs to be updated or if the access token has changed
// If so rebuild the schema
func (gh *graphHandler) setup(headers http.Header) error {
    // Check if Authorization header is present
    authHeaders, ok := headers["Authorization"]
    if !ok || len(authHeaders) == 0 {
        return &ServerError{StatusCode: http.StatusUnauthorized, Message: "No authorization header provided."}
    }
    authToken := authHeaders[0]

    ts, _ := gh.client.GetTimestamp(gh.graph)
    fmt.Println("HEADERS: ", headers)

    resourceList, err := getAllowedProjects("http://arborist-service/auth/mapping", authToken)
    if err != nil {
        log.WithFields(log.Fields{"graph": gh.graph, "error": err}).Error("auth/mapping fetch and processing step failed")
        return  &ServerError{StatusCode: http.StatusUnauthorized, Message: fmt.Sprintf("%s", err)}
    }

    if ts == nil || ts.Timestamp != gh.timestamp || resourceList != nil {
        fmt.Println("YOU ARE HERE +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++", resourceList)
        log.WithFields(log.Fields{"graph": gh.graph}).Info("Reloading GraphQL schema")
        schema, err := gh.client.GetSchema(gh.graph)
        if err != nil {
            log.WithFields(log.Fields{"graph": gh.graph, "error": err}).Error("GetSchema error")
            return  &ServerError{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("%s", err)}
        }
        gqlSchema, err := buildGraphQLSchema(schema, gh.client, gh.graph, resourceList)
        if err != nil {
            log.WithFields(log.Fields{"graph": gh.graph, "error": err}).Error("GraphQL schema build failed")
            gh.gqlHandler = nil
            gh.timestamp = ""
            return &ServerError{StatusCode: http.StatusInternalServerError, Message: "GraphQL schema build failed"}
        } else {
            log.WithFields(log.Fields{"graph": gh.graph}).Info("Built GraphQL schema")
            gh.gqlHandler = handler.New(&handler.Config{
                Schema: gqlSchema,
            })
            gh.timestamp = ts.Timestamp
        }
    }

    return nil
}
