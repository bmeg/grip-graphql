package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/bmeg/grip/log"

	"github.com/bmeg/grip/gripql"
	"github.com/graphql-go/handler"
)

type GraphQLJS struct {
	Client    gripql.Client
	GjHandler *handler.Handler
	Gen3      bool
	GraphName string
	Config    string
}

type GraphQLInterface interface {
	Setup()
}

type GraphQLJSDATA struct {
	Data *GraphQLJS
}

type Gen3ServerError struct {
	StatusCode int
	Message    string
}

func (e *Gen3ServerError) Error() string {
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
	empty_map := make(map[string]any)
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

func GetAllowedProjects(url string, token string) ([]any, error) {
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

	// Filter accessable projects by /programs/{PROGRAM}/projects/{PROJECT} to ensure that only project id resources are filtered
	pattern := regexp.MustCompile(`^/programs/[^/]+/projects/[^/]+$`)
	ProjectIds := filterProjects(readAccessResources, pattern)

	s := make([]interface{}, len(ProjectIds))
	for i, v := range ProjectIds {
		s[i] = v
	}
	return s, nil
}

func filterProjects(input []string, pattern *regexp.Regexp) []string {
	var filtered []string
	for _, str := range input {
		if pattern.MatchString(str) {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

func (gh *GraphQLJSDATA) Setup(writer http.ResponseWriter, request *http.Request) (error, []any) {
	authHeaders, ok := request.Header["Authorization"]
	if !ok || len(authHeaders) == 0 {
		return &Gen3ServerError{StatusCode: http.StatusUnauthorized, Message: "No authorization header provided."}, nil
	}
	authToken := authHeaders[0]
	//ts, _ := gh.client.GetTimestamp(gh.graph)
	fmt.Println("HEADERS: ", authHeaders)

	resourceList, err := GetAllowedProjects("http://arborist-service/auth/mapping", authToken)
	if err != nil {
		log.WithFields(log.Fields{"graph": gh.Data.GraphName, "error": err}).Error("auth/mapping fetch and processing step failed")
		return &Gen3ServerError{StatusCode: http.StatusUnauthorized, Message: fmt.Sprintf("%s", err)}, nil
	}
	return nil, resourceList

}
