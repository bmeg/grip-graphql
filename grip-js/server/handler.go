package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/bmeg/grip-graphql/middleware"
	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"github.com/dop251/goja"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

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

type QueryField struct {
	name    string
	field   *graphql.Field
	handler func(goja.FunctionCall) goja.Value
}

type GraphQLJS struct {
	client    gripql.Client
	gjHandler *handler.Handler
	Pool      sync.Pool
	cw        *JSClientWrapper
	auth      bool
	//Once      sync.Once
}

type Endpoint struct {
	client        gripql.Client
	vm            *goja.Runtime
	cw            *JSClientWrapper
	queryNodes    map[string]QueryField
	mutationNodes map[string]QueryField
}

func (e *Endpoint) Add(x map[string]any) {
	o, err := e.parseJSHandler(x)
	if err == nil {
		e.queryNodes[o.name] = o
		log.Infof("Added GraphQL query node : %s", o.name)
	} else {
		log.Errorf("Query 'add' error: %s", err)
	}
}

func (e *Endpoint) AddMutation(x map[string]any) {
	o, err := e.parseJSHandler(x)
	if err == nil {
		e.mutationNodes[o.name] = o
		log.Infof("Added GraphQL mutation node : %s", o.name)
	} else {
		log.Errorf("Query 'addMutation' error: %s", err)
	}
}

func (e *Endpoint) Build() (*graphql.Schema, error) {
	qf := graphql.Fields{}
	for k, v := range e.queryNodes {
		log.Debugf("build fields: %+v", v.field)
		qf[k] = v.field
	}

	queryObj := graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: qf})
	schemaConfig := graphql.SchemaConfig{
		Query: queryObj,
	}

	if len(e.mutationNodes) > 0 {
		mf := graphql.Fields{}
		for k, v := range e.mutationNodes {
			log.Debugf("build mutation fields: %+v", v.field)
			mf[k] = v.field
		}
		mutationObj := graphql.NewObject(graphql.ObjectConfig{Name: "Mutation", Fields: mf})
		schemaConfig.Mutation = mutationObj
	}

	gqlSchema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, fmt.Errorf("graphql.NewSchema error: %v", err)
	}
	return &gqlSchema, nil
}

func NewGraphQLJS(graph string, auth bool, client gripql.Client, code string) *GraphQLJS {
	vm := goja.New()
	vm.SetFieldNameMapper(JSRenamer{})
	jsClient, err := GetJSClient(graph, client, vm, auth)
	if err != nil {
		log.Infof("js error: %s\n", err)
	}

	e := &Endpoint{queryNodes: map[string]QueryField{}, mutationNodes: map[string]QueryField{}, client: client, vm: vm, cw: jsClient}
	vm.Set("endpoint", map[string]any{
		"add":         e.Add,
		"addMutation": e.AddMutation,
		"String":      "String",
		"Int":         "Int",
		"Float":       "Float",
		"Boolean":     "Boolean",
	})

	vm.Set("print", fmt.Printf) //Adding print statement for debugging. This may need to be removed/updated

	_, err = vm.RunString(code)
	if err != nil {
		log.Errorf("Error running data config %s", err)
	}

	schema, err := e.Build()
	if err != nil {
		log.Errorf("Error building Handler: %s", err)
	}
	var hnd *handler.Handler = handler.New(&handler.Config{
		Schema: schema,
	})
	gh := &GraphQLJS{client: client, gjHandler: hnd, cw: jsClient, auth: auth}
	return gh
}

func NewHTTPHandler(client gripql.Client, config map[string]string) (http.Handler, error) {
	configPath := ""
	graph := ""
	if c, ok := config["config"]; ok {
		configPath = c
	}
	if c, ok := config["graph"]; ok {
		graph = c
	}
	if graph == "" || configPath == "" {
		return nil, fmt.Errorf("graph or config not defined. These args must be specified on startup")
	}

	auth := true
	if c, ok := config["auth"]; ok {
		auth, _ = strconv.ParseBool(c)
	}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	log.Infof("Creating new pool ==============================================================")
	Pool := sync.Pool{
		New: func() any {
			return NewGraphQLJS(graph, auth, client, string(data))
		},
	}
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Debug("Getting graph handler from Sync Pool +++++++++++++++++++++++++++++++++++++++++++++++++++++")
		gh := Pool.Get().(*GraphQLJS)
		defer func() {
			log.Debug("Putting graph handler back to Pool ---------------------------------------------------")
			Pool.Put(gh)
		}()
		gh.ServeHTTP(writer, request)
	}), nil
}

func (gh *GraphQLJS) ServeHTTP(writer http.ResponseWriter, request *http.Request) error {

	if request.URL.Path == "" || request.URL.Path == "/" {
		writer.Write([]byte(sandBox))
		return nil
	}
	if request.URL.Path == "/api" || request.URL.Path == "api" {
		requestHeaders := request.Header
		ctx := context.WithValue(context.Background(), "Header", requestHeaders)
		if gh.auth {
			var jwtHandler middleware.JWTHandler = &middleware.ProdJWTHandler{}
			if gh.cw.graph == "TEST" {
				jwtHandler = &middleware.MockJWTHandler{}
			}
			log.Debug("Request Headers: ", requestHeaders)
			if val, ok := requestHeaders["Authorization"]; ok {
				Token := val[0]
				resourceList, err := jwtHandler.HandleJWTToken(Token, "read")
				log.Debugln("Resource List: ", resourceList)
				if err != nil {
					log.Debugln("HandleJWTToken Err: ", err)
					middleware.HandleError(err, writer)
					return err
				}

				if len(resourceList) == 0 || err != nil {
					if len(resourceList) == 0 {
						err = &middleware.ServerError{StatusCode: http.StatusForbidden, Message: "User does not have access to any projects"}
					}
					middleware.HandleError(err, writer)
					return err
				}
				ctx = context.WithValue(ctx, "ResourceList", resourceList)
			} else {
				err := middleware.HandleError(&middleware.ServerError{StatusCode: http.StatusUnauthorized, Message: "No authorization header provided."}, writer)
				log.Infoln("ERR: ", err)
				return err
			}
		}
		gh.gjHandler.ServeHTTP(writer, request.WithContext(ctx))
	}
	return nil
}
