package main

import (
	"net/http"

	"github.com/bmeg/grip-graphql/gripgraphql"
	"github.com/bmeg/grip/gripql"
)

func NewHTTPHandler(client gripql.Client, config map[string]string) (http.Handler, error) {
	return gripgraphql.NewHTTPHandler(client, config)
}
