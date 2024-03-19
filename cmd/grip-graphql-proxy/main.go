package main

import (
	"net/http"

	"github.com/bmeg/grip-graphql/gripgraphql"
	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"github.com/bmeg/grip/util/rpc"
)

func main() {
	client, err := gripql.Connect(rpc.ConfigWithDefaults("localhost:8202"), false)
	if err != nil {
		log.Errorf("Error: %s", err)
	}
	httpHandler, err := gripgraphql.NewHTTPHandler(client)
	if err != nil {
		log.Errorf("Error: %s", err)
	}

	http.ListenAndServe(":8080", httpHandler)
}
