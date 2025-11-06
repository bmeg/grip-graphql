package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MockJWTHandler struct{}
type ProdJWTHandler struct{}
type JWTHandler interface {
	// This method assumes that you are checking method, service on resource paths of form
	// /programs/[program_name]/projects/[project_name]
	GetAllowedResources(token, method, service string) ([]any, error)
	// This method is for checking exact permission for a specific resource
	CheckResourceServiceAccess(token, method, service, resource string) (bool, error)
}

func (m *MockJWTHandler) CheckResourceServiceAccess(token, method, service, resource string) (bool, error) {
	return false, nil
}
func (m *MockJWTHandler) GetAllowedResources(token string, method string, service string) ([]any, error) {
	expiration, err := GetExpiration(token)
	if err != nil {
		return nil, &ServerError{StatusCode: 400, Message: fmt.Sprintf("%s Failed to get expiration from token %s", err, token)}
	}

	if expiration.After(time.Now()) {
		claims, err := m.decodeToken(token)
		if err != nil {
			return nil, &ServerError{StatusCode: 400, Message: fmt.Sprintf("failed to parse token data %#v", token)}
		}
		subject, err := claims.GetSubject()
		if err != nil {
			return nil, &ServerError{StatusCode: 400, Message: fmt.Sprintf("failed to parse token claims data %#v", claims)}
		}
		fmt.Println("SUBJECT:", subject, "")
		if method == subject || "*" == subject {
			return []any{"/programs/ohsu/projects/test"}, nil
		}
		return []any{}, nil
	}
	return nil, &ServerError{StatusCode: 401, Message: fmt.Sprintf("token %s has expired %s", token, expiration)}
}

func (m *MockJWTHandler) decodeToken(tokenString string) (jwt.MapClaims, error) {
	/* A Mock token decoder designed to decode mock tokens for testing purposes only */
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("foo-bar-signature"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func (m *ProdJWTHandler) CheckResourceServiceAccess(token, method, service, resource string) (bool, error) {
	expiration, err := GetExpiration(token)
	if err != nil {
		return false, &ServerError{StatusCode: 400, Message: fmt.Sprintf("%s Failed to get expiration from token %s", err, token)}
	}
	if !expiration.After(time.Now()) {
		return false, &ServerError{StatusCode: 401, Message: fmt.Sprintf("Token %s has expired %s", token, expiration)}
	}
	return IsProjectAllowedOnResource("http://arborist-service/auth/mapping", token, method, service, resource)
}

func (j *ProdJWTHandler) GetAllowedResources(token string, method string, service string) ([]any, error) {
	expiration, err := GetExpiration(token)
	if err != nil {
		return nil, &ServerError{StatusCode: 400, Message: fmt.Sprintf("%s Failed to get expiration from token %s", err, token)}
	}
	if !expiration.After(time.Now()) {
		return nil, &ServerError{StatusCode: 401, Message: fmt.Sprintf("Token %s has expired %s", token, expiration)}
	}
	return GetAllowedProjects("http://arborist-service/auth/mapping", token, method, service)
}
