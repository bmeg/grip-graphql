package main

import (
	"net/http"

	"github.com/bmeg/grip-graphql/gripgraphql"
	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"github.com/bmeg/grip/util/rpc"
	"github.com/spf13/pflag"
)

var gripServer = "localhost:8202"
var serverPort = "8080"
var configFile = "config.js"
var graph = "test-db"
var writable = false

func main() {

	//  <grip server> <server port> <config> <database>

	pflag.StringVar(&gripServer, "grip", gripServer, "GRIP server")
	pflag.StringVar(&serverPort, "port", serverPort, "Proxy port")
	pflag.StringVar(&configFile, "config", configFile, "Config")
	pflag.StringVar(&graph, "graph", graph, "Graph")
	pflag.BoolVar(&writable, "write", writable, "Allow Write")

	pflag.Parse()

	client, err := gripql.Connect(rpc.ConfigWithDefaults(gripServer), writable)
	if err != nil {
		log.Errorf("Error: %s", err)
	}
	config := map[string]string{}

	config["graph"] = graph
	config["config"] = configFile

	httpHandler, err := gripgraphql.NewHTTPHandler(client, config)
	if err != nil {
		log.Errorf("Error: %s", err)
	}
	http.ListenAndServe(":"+serverPort, httpHandler)
}
