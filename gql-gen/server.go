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
	router *gin.Engine
	config map[string]string
}

func (gh *Handler) graphqlHandler(client gripql.Client) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		//GripClient: client,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func (gh *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	gh.router.ServeHTTP(writer, request)
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
	r.NoRoute(func(c *gin.Context) {
		log.WithFields(log.Fields{
			"graph":  nil,
			"status": "404",
		}).Info(c.Request.URL.Path + " Not Found")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "404",
			"message": c.Request.URL.Path + " Not Found",
			"data":    nil,
		})
	})
	r.Use(gin.Recovery())

	r.RemoveExtraSlash = true
	h := &Handler{
		router: r,
		config: config,
	}
	r.POST("/query", h.graphqlHandler(client))
	r.GET("/", gin.WrapH(playground.Handler("GraphQL", "/query"))) //r.GET("/", playgroundHandler())
	r.Run()
	return h, nil
}
