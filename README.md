### Grip Graphql Plugins

grip-graphql is a collection of go plugins designed and implemented for connecting a [Grip](https://github.com/bmeg/grip) server to other microservices in a modified [Gen3](https://gen3.org/) software stack.

gen3_writer directory contains a [Gin](https://github.com/gin-gonic/gin) go server plugin that is used for Writing / Deleting data from Graphs on a grip server

gripgraphql directory contains a graphql based read query plugin that uses a [goja](https://github.com/dop251/goja) engine to read from a static schema defined as a config file to create custom graphql queries that can be used to abstract Grip's complex query language into a more digestible query format for the frontend to use.

graphql_gen3 is a legacy implementation of a reader reader plugin using a more traditional graphql schema builder.

See ./gen3_writer for tests and additional documentation
