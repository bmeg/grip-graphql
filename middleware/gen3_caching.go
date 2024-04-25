package middleware

import (
	"fmt"
	"time"
    "strings"

	"github.com/patrickmn/go-cache"
    "github.com/golang-jwt/jwt/v5"

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

func HandleJWTToken(token string) ([]any, error) {
	cachedData, found := jwtCache.Get(token)

    // If cache hit check expiration and return resourceList
	if found {
		tokenData, ok := cachedData.(TokenData)
		if !ok {
            return nil, fmt.Errorf("failed to parse token data %#v", cachedData)
		}

        fmt.Println("expiration:", tokenData.Expiration)
        
		if tokenData.Expiration.After(time.Now()) {
			return tokenData.ResourceList, nil
		}
		jwtCache.Delete(token)
        err := &ServerError{StatusCode: 401, Message: fmt.Sprintf("token %s has expired", tokenData.Expiration)}
		return nil, err
	} 

    // Otherise check expiration, add token to cache and return resourceList 
    expiration, err := GetExpiration(token)
    if err != nil{
        return nil, err
    }

    if expiration.After(time.Now()){
        resourceList, err := AddJWTToken(token, expiration)
        if err != nil {
            return nil, err
        }

        return resourceList, err
    }
    return nil, fmt.Errorf("token %s has expired. Expiration: %s", token, expiration)
}

func AddJWTToken(token string, expiration time.Time) ([]any, error){
    resourceList, err := GetAllowedProjects("http://arborist-service/auth/mapping", token)
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

