package main

import (
	"net/http"

	"github.com/bmeg/grip-graphql/grip-js/server"
	"github.com/bmeg/grip/gripql"
)

func NewHTTPHandler(client gripql.Client, config map[string]string) (http.Handler, error) {
	return server.NewHTTPHandler(client, config)
}
