package main

import (
	"net/http"

	"github.com/bmeg/grip-graphql/gripgraphql"
	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"github.com/bmeg/grip/util/rpc"

	"github.com/spf13/pflag"
)

func main() {

	pflag.Parse()

	client, err := gripql.Connect(rpc.ConfigWithDefaults("localhost:8202"), false)
	if err != nil {
		log.Errorf("Error: %s", err)
	}
	config := map[string]string{}

	httpHandler, err := gripgraphql.NewHTTPHandler(client, config)
	if err != nil {
		log.Errorf("Error: %s", err)
	}
	http.ListenAndServe(":8080", httpHandler)
}
