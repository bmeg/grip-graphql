package main

import (
	"net/http"

	"github.com/bmeg/grip-graphql/grip-js/pluginface"
	"github.com/bmeg/grip-graphql/grip-js/server"
	"github.com/bmeg/grip/gripql"
)

var Plugin pluginface.HTTPHandlerCreator

func NewHTTPHandler(client gripql.Client, config map[string]string) (http.Handler, error) {
	return server.NewHTTPHandler(client, config)
}
