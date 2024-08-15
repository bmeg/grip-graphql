package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/bmeg/grip/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/patrickmn/go-cache"
)

type TokenData struct {
	Expiration   time.Time
	ResourceList []any
}

var jwtCache *cache.Cache

func init() {
	jwtCache = cache.New(0, 0)
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
	return time.Time{}, fmt.Errorf("Expiration field 'exp' type float64 not found in token %s", token)
}

func HandleJWTToken(token string, perm_method string) ([]any, error) {
	cachedData, found := jwtCache.Get(token)

	// If cache hit check expiration and return resourceList
	if found {
		tokenData, ok := cachedData.(TokenData)
		if !ok {
			return nil, &ServerError{StatusCode: 400, Message: fmt.Sprintf("failed to parse token data %#v", cachedData)}
		}

		fmt.Println("expiration:", tokenData.Expiration)

		if tokenData.Expiration.After(time.Now()) {
			log.Infoln("Retrieved Cached token")
			return tokenData.ResourceList, nil
		}
		jwtCache.Delete(token)
		return nil, &ServerError{StatusCode: 401, Message: fmt.Sprintf("token %s has expired %s", token, tokenData.Expiration)}
	}

	// Otherise check expiration, add token to cache and return resourceList
	expiration, err := GetExpiration(token)
	if err != nil {
		return nil, &ServerError{StatusCode: 400, Message: fmt.Sprintf("Failed to get expiration from token %s", token)}
	}

	if expiration.After(time.Now()) {
		resourceList, err := AddJWTToken(token, expiration, perm_method)
		if err != nil {
			return nil, err
		}
		return resourceList, nil
	}
	return nil, &ServerError{StatusCode: 401, Message: fmt.Sprintf("Token validation failed for token: %s", token)}
}

func AddJWTToken(token string, expiration time.Time, method string) ([]any, error) {
	resourceList, err := GetAllowedProjects("http://arborist-service/auth/mapping", token, method)
	if err != nil {
		return nil, err
	}

	tokenData := TokenData{
		Expiration:   expiration,
		ResourceList: resourceList,
	}
	jwtCache.Set(token, tokenData, cache.NoExpiration)
	return resourceList, nil
}
