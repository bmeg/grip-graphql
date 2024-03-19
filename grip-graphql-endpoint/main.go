package main

import (
	"net/http"

	"github.com/bmeg/grip-graphql/gripgraphql"
	"github.com/bmeg/grip/gripql"
)

func NewHTTPHandler(client gripql.Client) (http.Handler, error) {
	return gripgraphql.NewHTTPHandler(client)
}
