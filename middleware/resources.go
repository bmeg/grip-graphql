package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/bmeg/grip/log"
	"github.com/golang-jwt/jwt/v5"
)

var HTTPClient = http.DefaultClient

func getAuthMappings(url string, token string) (any, error) {
	GetRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	GetRequest.Header.Set("Authorization", token)
	GetRequest.Header.Set("Accept", "application/json")
	fetchedData, err := HTTPClient.Do(GetRequest)
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

func hasPermission(permissions []any, method string, service string) bool {
	for _, permission := range permissions {
		permission := permission.(map[string]any)
		if permission["service"] == service &&
			permission["method"] == method {
			return true
		}
	}
	return false
}

func GetAllowedProjects(url string, token string, method string, service string) ([]any, error) {
	var readAccessResources []string
	authMappings, err := getAuthMappings(url, token)
	if err != nil {
		return nil, &ServerError{StatusCode: 400, Message: fmt.Sprintf("%s", err)}
	}

	// Iterate through /auth/mapping resultant dict checking for valid permissions
	for resourcePath, permissions := range authMappings.(map[string]any) {
		if hasPermission(permissions.([]any), method, service) {
			readAccessResources = append(readAccessResources, resourcePath)
		}
	}

	// Filter accessable projects by /programs/{PROGRAM}/projects/{PROJECT} to ensure that only project id resources are filtered
	pattern := regexp.MustCompile(`^/programs/[^/]+/projects/[^/]+$`)
	ProjectIds := filterProjects(readAccessResources, pattern)

	s := make([]any, len(ProjectIds))
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
	slices.Sort(filtered)
	return filtered
}

func IsProjectAllowedOnResource(url, token, method, service, resource string) (bool, error) {
	authMappings, err := getAuthMappings(url, token)
	if err != nil {
		return false, &ServerError{StatusCode: 400, Message: fmt.Sprintf("%s", err)}
	}

	// Iterate through /auth/mapping resultant dict checking for valid permissions
	for resourcePath, permissions := range authMappings.(map[string]any) {
		if resource == resourcePath {
			if hasPermission(permissions.([]any), method, service) {
				return true, nil
			}
		}
	}
	return false, nil
}

func GetExpiration(tokenString string) (time.Time, error) {
	// Also consider trimming the 'Bearer ' prefix too
	tokenString = strings.TrimPrefix(tokenString, "bearer ")
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return time.Time{}, err
	}

	// Parse and convert from float64 epoch time to time.Time
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok {
			temp := int64(exp)
			exp := time.Unix(temp, 0)
			return exp, nil
		}
	}
	return time.Time{}, fmt.Errorf("Expiration field 'exp' type float64 not found in token %v", token)
}
