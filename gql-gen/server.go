package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/bmeg/grip-graphql/gql-gen/generated"
	"github.com/bmeg/grip-graphql/gql-gen/graph"

	"github.com/vektah/gqlparser/v2/ast"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router  *gin.Engine
	config  map[string]string
	handler *handler.Server
	client  gripql.Client
}

func (gh *Handler) graphqlHandler(client gripql.Client) gin.HandlerFunc {
	executableSchema := generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})
	schema := executableSchema.Schema()
	resolvers := &graph.Resolver{
		Schema: schema,
		GripDb: client,
	}
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers}))
	gh.handler = srv

	gh.handler.AddTransport(transport.Options{})
	gh.handler.AddTransport(transport.GET{})
	gh.handler.AddTransport(transport.POST{})
	gh.handler.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	gh.handler.Use(extension.Introspection{})
	gh.handler.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return func(c *gin.Context) {
		gh.handler.ServeHTTP(c.Writer, c.Request)
		//c.Next()
	}
}

func (gh *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Infoln("HELLO INSIDE SERVE HTTP ", request)
	gh.router.ServeHTTP(writer, request)
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/graphql/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func NewHTTPHandler(client gripql.Client, config map[string]string) (http.Handler, error) {
	// Setting up Gin
	r := gin.New()
	logConfig := log.Logger{
		Level:     "info",
		Formatter: "json",
	}
	log.ConfigureLogger(logConfig)
	r.Use(gin.Logger())
	/*r.NoRoute(func(c *gin.Context) {
	log.WithFields(log.Fields{
		"graph":  nil,
		"status": "404",
	}).Info(c.Request.URL.Path + " Not Found")
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"status":  "404",
		"message": c.Request.URL.Path + " Not Found",
		"data":    nil,
	})
	})*/
	r.Use(gin.Recovery())

	r.RemoveExtraSlash = true

	h := &Handler{
		router: r,
		config: config,
		client: client,
	}
	r.POST("/query", h.graphqlHandler(client))
	r.GET("/", playgroundHandler())
	return h, nil

}
