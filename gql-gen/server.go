package main

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bmeg/grip-graphql/middleware"

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

func RegError(c *gin.Context, writer http.ResponseWriter, graph string, err error) {
	if ae, ok := err.(*middleware.ServerError); ok {
		log.WithFields(log.Fields{
			"graph":  graph,
			"status": ae.StatusCode,
		}).Info(ae.Message)
		c.AbortWithStatusJSON(ae.StatusCode, gin.H{
			"status":  ae.StatusCode,
			"message": ae.Message,
			"data":    nil,
		})
		return
	}
	log.WithFields(log.Fields{
		"graph":  graph,
		"status": "500",
	}).Info("Internal Server Error")
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"status":  "500",
		"message": "[500] Internal Server Error",
		"data":    nil,
	})
}

func (gh *Handler) graphqlHandler(client gripql.Client, jwtHandler middleware.JWTHandler) gin.HandlerFunc {
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
		requestHeaders := c.Request.Header
		if val, ok := requestHeaders["Authorization"]; ok {
			Token := val[0]
			anyList, err := jwtHandler.HandleJWTToken(Token, "read")
			if err != nil {
				RegError(c, c.Writer, c.Param("graph"), err)
				return
			}
			c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "auth_list", anyList))
			//c.Set("auth_list", anyList)
		} else {
			RegError(c, c.Writer, c.Param("graph"), &middleware.ServerError{StatusCode: 400, Message: "Authorization token not provided"})
			return
		}
		gh.handler.ServeHTTP(c.Writer, c.Request)
	}
}

func (gh *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//log.Infoln("HELLO INSIDE SERVE HTTP ", request)
	gh.router.ServeHTTP(writer, request)
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/graphql/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func NewHTTPHandler(client gripql.Client, config map[string]string) (http.Handler, error) {

	var mware middleware.JWTHandler = &middleware.ProdJWTHandler{}
	if config["test"] == "true" {
		mware = &middleware.MockJWTHandler{}
	}

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
	r.POST("/query", h.graphqlHandler(client, mware))
	r.GET("/", playgroundHandler())
	return h, nil

}
