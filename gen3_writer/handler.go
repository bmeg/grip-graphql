/*
RESTFUL Gin Web endpoint
*/

package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/bmeg/grip-graphql/middleware"
	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"github.com/bmeg/grip/schema"
	"github.com/bmeg/grip/util/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

type Handler struct {
	router *gin.Engine
	client gripql.Client
	config map[string]string
}

func getFields(c *gin.Context) (gin.ResponseWriter, *http.Request, string) {
	return c.Writer, c.Request, c.Param("graph")
}

func convertAnyToStringSlice(anySlice []any) ([]string, error) {
	/* converts []any to []string */
	var stringSlice []string
	for _, v := range anySlice {
		str, ok := v.(string)
		if !ok {
			return nil, &middleware.ServerError{StatusCode: 500, Message: fmt.Sprintf("Element %v is not a string", v)}
		}
		stringSlice = append(stringSlice, str)
	}
	return stringSlice, nil
}

func ParseAccess(c *gin.Context, resourceList []string, resource string, method string) error {
	/*  Iterates through a list of Gen3 resoures and returns true if
	    resource matches the allowable list of resource types for the provided method */

	if len(resourceList) == 0 {
		return &middleware.ServerError{StatusCode: 403, Message: fmt.Sprintf("User is not allowed to %s on any resource path", method)}
	}
	for _, v := range resourceList {
		if resource == v {
			return nil
		}
	}
	return &middleware.ServerError{StatusCode: 403, Message: fmt.Sprintf("User is not allowed to %s on resource path: %s", method, resource)}
}

func TokenAuthMiddleware(jwtHandler middleware.JWTHandler) gin.HandlerFunc {
	/*  Authentication middleware function. Maps HTTP method to expected permssions.
	    If user permissions don't match, abort command and return 401 */

	return func(c *gin.Context) {
		// If request path is open access then no token needed no project_id needed.
		project_split := strings.Split(c.Param("project-id"), "-")
		if len(project_split) != 2 {
			// This error usually occurs when part of the path is correct or a typo in path.
			RegError(c, c.Writer, c.Param("graph"), &middleware.ServerError{StatusCode: 404, Message: fmt.Sprintf("incorrect path %s", c.Request.URL)})
			return
		}

		project_id := "/programs/" + project_split[0] + "/projects/" + project_split[1]
		requestHeaders := c.Request.Header
		if val, ok := requestHeaders["Authorization"]; ok {
			Token := val[0]
			var method string
			if c.Request.Method == http.MethodGet {
				method = "read"
			} else if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodDelete {
				method = "create"
			} else {
				RegError(c, c.Writer, c.Param("graph"), &middleware.ServerError{StatusCode: 405, Message: fmt.Sprintf("Method %s not allowed", c.Request.Method)})
				return
			}

			anyList, err := jwtHandler.HandleJWTToken(Token, method)
			if err != nil {
				RegError(c, c.Writer, c.Param("graph"), err)
				return
			}

			resourceList, convErr := convertAnyToStringSlice(anyList)
			if convErr != nil {
				RegError(c, c.Writer, c.Param("graph"), convErr)
				return
			}

			log.Infof("Resource List for method '%s': %s", method, resourceList)
			err = ParseAccess(c, resourceList, project_id, method)
			if err != nil {
				RegError(c, c.Writer, c.Param("graph"), err)
				return
			}
		} else {
			RegError(c, c.Writer, c.Param("graph"), &middleware.ServerError{StatusCode: 400, Message: "Authorization token not provided"})
			return
		}
		c.Next()
	}
}

func NewHTTPHandler(client gripql.Client, config map[string]string) (http.Handler, error) {

	// Need a way to toggle on mock auth
	var mware middleware.JWTHandler = &middleware.ProdJWTHandler{}
	if config["test"] == "true" {
		mware = &middleware.MockJWTHandler{}
	}

	// Including below line to run in prod mode
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// Since using Grip logging functions, log config needs to be set to properly format JSON data outputs
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

	// Was getting 404s before adding this.
	r.RemoveExtraSlash = true

	h := &Handler{
		router: r,
		client: client,
		config: config,
	}

	r.GET("list-graphs", func(c *gin.Context) {
		h.ListGraphs(c, c.Writer)
	})
	r.GET("_status", func(c *gin.Context) {
		Response(c, c.Writer, "", 200, 200, fmt.Sprintf("[200] healthy _status"))
	})
	r.POST("/add-graph/:graph/:project-id", TokenAuthMiddleware(mware), func(c *gin.Context) {
		h.AddGraph(c)
	})

	g := r.Group(":graph")
	g.Use(TokenAuthMiddleware(mware))
	g.POST("/add-vertex/:project-id", func(c *gin.Context) {
		h.WriteVertex(c)
	})
	g.POST("/add-edge/:project-id", func(c *gin.Context) {
		h.WriteEdge(c)
	})
	g.POST("/bulk-load/:project-id", func(c *gin.Context) {
		h.BulkStream(c)
	})
	g.POST("/bulk-load-raw/:project-id", func(c *gin.Context) {
		h.BulkStreamRaw(c)
	})
	g.POST("/add-grip-schema/:project-id", func(c *gin.Context) {
		h.AddGripSchema(c)
	})
	g.POST("/add-json-schema/:project-id", func(c *gin.Context) {
		h.AddJsonSchema(c)
	})
	g.DELETE("/del-graph/:project-id", func(c *gin.Context) {
		h.DeleteGraph(c)
	})
	g.DELETE("/del-edge/:edge-id/:project-id", func(c *gin.Context) {
		h.DeleteEdge(c, c.Param("edge-id"))
	})
	g.DELETE("/del-vertex/:vertex-id/:project-id", func(c *gin.Context) {
		h.DeleteVertex(c, c.Param("vertex-id"))
	})
	g.DELETE("/bulk-delete/:project-id", func(c *gin.Context) {
		h.BulkDelete(c)
	})
	g.DELETE("/proj-delete/:project-id", func(c *gin.Context) {
		h.ProjectDelete(c)
	})
	g.GET("/list-labels", func(c *gin.Context) {
		h.ListLabels(c)
	})
	g.GET("/get-schema/:project-id", func(c *gin.Context) {
		h.GetSchema(c)
	})
	g.GET("/get-graph/:project-id", func(c *gin.Context) {
		h.GetGraph(c)
	})
	g.GET("/get-vertex/:vertex-id/:project-id", func(c *gin.Context) {
		h.GetVertex(c, c.Param("vertex-id"))
	})
	g.GET("/get-vertices/:project-id", func(c *gin.Context) {
		h.GetProjectVertices(c)
	})
	return h, nil
}

// ServeHTTP responds to HTTP graphql requests
func (gh *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	gh.router.ServeHTTP(writer, request)
}

func GetInternalServerErr(err error) *middleware.ServerError {
	return &middleware.ServerError{StatusCode: 500, Message: fmt.Sprintf("%s", err)}
}

func Response(c *gin.Context, writer http.ResponseWriter, graph string, data any, statusCode int, message string) {
	log.WithFields(log.Fields{
		"graph":  graph,
		"status": statusCode,
		"data":   data,
	}).Info(message)
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": message,
		"data":    data,
	})
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

func (gh *Handler) ListLabels(c *gin.Context) {
	writer, _, graph := getFields(c)
	labels, err := gh.client.ListLabels(graph)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, labels, 200, fmt.Sprintf("[200] list-labels on graph %s", graph))
}

func (gh *Handler) GetSchema(c *gin.Context) {
	writer, _, graph := getFields(c)
	schema, err := gh.client.GetSchema(graph)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, schema, 200, fmt.Sprintf("[200] get-schema on graph %s", graph))
}

func StartMultipartForm(c *gin.Context, writer gin.ResponseWriter, request *http.Request, graph string) (error, gripql.Client, *bytes.Buffer) {
	err := request.ParseMultipartForm(1024 * 1024 * 1024) // 10 GB limit
	if err != nil {
		RegError(c, writer, graph, &middleware.ServerError{StatusCode: 400, Message: fmt.Sprintf("Error parsing form: %s", err)})
		return err, gripql.Client{}, nil
	}
	file, _, err := request.FormFile("file")
	if err != nil {
		RegError(c, writer, graph, &middleware.ServerError{StatusCode: 400, Message: fmt.Sprintf("failed to parse attached file: %s", err)})
		return err, gripql.Client{}, nil
	}
	file.Close()

	conn, err := gripql.Connect(rpc.ConfigWithDefaults("localhost:8202"), true)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return err, gripql.Client{}, nil
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return err, gripql.Client{}, nil
	}
	return nil, conn, buf
}

func (gh *Handler) AddJsonSchema(c *gin.Context) {
	/* This function assumes that Json schema of json format will be submitted and must be converted into grip scheam format */
	writer, request, graph := getFields(c)
	err, conn, buf := StartMultipartForm(c, writer, request, graph)
	if err != nil {
		return
	}
	graphs, err := schema.ParseJSchema(buf.Bytes(), graph)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(fmt.Errorf("json parse error: %s", err)))
		return
	}
	for _, g := range graphs {
		err := conn.AddSchema(g)
		if err != nil {
			RegError(c, writer, graph, GetInternalServerErr(err))
			return
		}
	}
	Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] add-schema on graph %s", graph))
}

func (gh *Handler) AddGripSchema(c *gin.Context) {
	/*  This function assumes that json blob will be submitted already in the form of a grip.Schema structure so that all
	that is needed is an unmarhsal into the structure */
	writer, request, graph := getFields(c)
	var graphs []*gripql.Graph
	err, conn, buf := StartMultipartForm(c, writer, request, graph)
	if err != nil {
		return
	}

	graphs, err = schema.ParseJSONGraphs(buf.Bytes())
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(fmt.Errorf("json parse error: %s", err)))
		return
	}
	for _, g := range graphs {
		err := conn.AddSchema(g)
		if err != nil {
			RegError(c, writer, graph, GetInternalServerErr(err))
			return
		}
	}
	Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] add-schema on graph %s", graph))
}

func (gh *Handler) GetGraph(c *gin.Context) {
	writer, _, graph := getFields(c)
	graph_data, err := gh.client.GetMapping(graph)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, graph_data, 200, "[200] get-graph")
}

func (gh *Handler) ListGraphs(c *gin.Context, writer http.ResponseWriter) {
	graphs, err := gh.client.ListGraphs()
	if err != nil {
		RegError(c, writer, "", GetInternalServerErr(err))
		return
	}
	Response(c, writer, "", graphs, 200, "[200] list-graphs")
}

func (gh *Handler) AddGraph(c *gin.Context) {
	writer, _, graph := getFields(c)
	err := gh.client.AddGraph(graph)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] add-graph added: %s", graph))
}

func (gh *Handler) DeleteGraph(c *gin.Context) {
	writer, _, graph := getFields(c)
	err := gh.client.DeleteGraph(graph)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] delete-graph deleted: %s", graph))
}

func (gh *Handler) GetVertex(c *gin.Context, vertex string) {
	writer, _, graph := getFields(c)
	gql_vertex, err := gh.client.GetVertex(graph, vertex)
	if err != nil {
		if strings.Contains(err.Error(), "rpc error: code = NotFound") {
			RegError(c, writer, graph, &middleware.ServerError{StatusCode: 404, Message: fmt.Sprintf("%s", err)})
			return
		}
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, gql_vertex, 200, fmt.Sprintf("[200] get-vertex: %s", gql_vertex))
}

func (gh *Handler) GetEdge(c *gin.Context, edge string) {
	writer, _, graph := getFields(c)
	gql_edge, err := gh.client.GetEdge(graph, edge)
	if err != nil {
		if strings.Contains(err.Error(), "rpc error: code = NotFound") {
			RegError(c, writer, graph, &middleware.ServerError{StatusCode: 404, Message: fmt.Sprintf("%s", err)})
			return
		}
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, gql_edge, 200, fmt.Sprintf("[200] get-edge: %s", gql_edge))
}

func (gh *Handler) DeleteEdge(c *gin.Context, edgeId string) {
	writer, _, graph := getFields(c)
	if _, err := gh.client.GetEdge(graph, edgeId); err == nil {
		err := gh.client.DeleteEdge(graph, edgeId)
		if err != nil {
			RegError(c, writer, graph, GetInternalServerErr(err))
			return
		}
		Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] delete-edge: %s", edgeId))
	} else {
		RegError(c, writer, graph, GetInternalServerErr(err))
	}
}

func (gh *Handler) DeleteVertex(c *gin.Context, vertexId string) {
	writer, _, graph := getFields(c)
	_, err := gh.client.GetVertex(graph, vertexId)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	err = gh.client.DeleteVertex(graph, vertexId)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] delete-vertex: %s", vertexId))
}

func (gh *Handler) BulkDelete(c *gin.Context) {
	writer, request, graph := getFields(c)

	var body []byte
	var err error
	delData := &gripql.DeleteData{}
	if body, err = io.ReadAll(request.Body); err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	if body == nil {
		RegError(c, writer, graph, err)
		return
	} else {
		if err := protojson.Unmarshal([]byte(body), delData); err != nil {
			RegError(c, writer, graph, err)
			return
		}
	}

	if err := gh.client.BulkDelete(delData); err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] bulk-delete on graph %s", graph))
}

func (gh *Handler) GetProjectVertices(c *gin.Context) {
	ctx := context.Background()

	writer, _, graph := getFields(c)
	project_id := c.Param("project-id")
	str_split := strings.Split(project_id, "-")
	project := "/programs/" + str_split[0] + "/projects/" + str_split[1]

	Vquery := gripql.V().Has(gripql.Eq("auth_resource_path", project))
	query := &gripql.GraphQuery{Graph: c.Param("graph"), Query: Vquery.Statements}
	result, err := gh.client.Traversal(ctx, query)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}

	flusher, ok := writer.(http.Flusher)
	if !ok {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}

	fm := gripql.NewFlattenMarshaler()
	for i := range result {
		rowString, _ := fm.Marshal(i.GetVertex())
		_, err := writer.Write(append(rowString, '\n'))
		if err != nil {
			RegError(c, writer, graph, GetInternalServerErr(err))
			return
		}
		flusher.Flush()
	}
}

func (gh *Handler) ProjectDelete(c *gin.Context) {
	ctx := context.Background()
	var delVs []string

	writer, _, graph := getFields(c)
	project_id := c.Param("project-id")
	str_split := strings.Split(project_id, "-")
	project := "/programs/" + str_split[0] + "/projects/" + str_split[1]

	Vquery := gripql.V().Has(gripql.Eq("auth_resource_path", project)).Render("_id")
	query := &gripql.GraphQuery{Graph: c.Param("graph"), Query: Vquery.Statements}
	result, err := gh.client.Traversal(ctx, query)
	if err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}

	/* Running bulkDelete with only the vertex ids specified will remove all of the disconnected edges caused by removing
	   All vertices in the graph. */
	for i := range result {
		v := i.ToInterface().(map[string]any)["render"].(string)
		delVs = append(delVs, v)
	}

	delData := &gripql.DeleteData{Graph: graph, Vertices: delVs, Edges: []string{}}

	if err := gh.client.BulkDelete(delData); err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] project-delete on project %s", project_id))
}

func (gh *Handler) WriteVertex(c *gin.Context) {
	writer, request, graph := getFields(c)

	var body []byte
	var err error
	v := &gripql.Vertex{}

	if body, err = io.ReadAll(request.Body); err != nil {
		RegError(c, writer, graph, err)
		return
	}
	if body == nil {
		RegError(c, writer, graph, err)
		return
	} else {
		if err := protojson.Unmarshal([]byte(body), v); err != nil {
			RegError(c, writer, graph, err)
			return
		}
	}
	if err := gh.client.AddVertex(graph, v); err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] write-vertex: %s", v.GetId()))
}

func (gh *Handler) WriteEdge(c *gin.Context) {
	writer, request, graph := getFields(c)

	var body []byte
	var err error
	e := &gripql.Edge{}

	if body, err = io.ReadAll(request.Body); err != nil {
		RegError(c, writer, graph, err)
		return
	}
	if body == nil {
		RegError(c, writer, graph, &middleware.ServerError{StatusCode: 400, Message: "Request body empty. Cannot parse request"})
		return
	} else {
		if err := protojson.Unmarshal([]byte(body), e); err != nil {
			RegError(c, writer, graph, err)
			return
		}
	}
	if err := gh.client.AddEdge(graph, e); err != nil {
		RegError(c, writer, graph, GetInternalServerErr(err))
		return
	}
	Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] write-edge: %s", e.GetId()))
}

func (gh *Handler) BulkStreamRaw(c *gin.Context) {
	writer, request, graph := getFields(c)
	project_id := c.Param("project-id")
	mapExtraArgs := map[string]any{}
	proj_split := strings.Split(project_id, "-")
	if len(proj_split) != 2 {
		RegError(c, writer, graph, &middleware.ServerError{StatusCode: 400, Message: fmt.Sprintf("Error parsing project_id: %s. Must be of form PROGRAM-PROJECT", project_id)})
		return
	}
	mapExtraArgs["auth_resource_path"] = "/programs/" + proj_split[0] + "/projects/" + proj_split[1]

	var Hostname string = c.Request.Host
	if strings.Contains(c.Request.Host, ":") {
		var err error
		Hostname, _, err = net.SplitHostPort(c.Request.Host)
		if err != nil {
			RegError(c, writer, graph, GetInternalServerErr(err))
			return
		}
	}

	mapExtraArgs["namespace"] = Hostname
	log.Infof("Using hostname: %s", Hostname)

	host := "localhost:8202"
	var err error
	var res *gripql.BulkJsonEditResult
	var warnings []string

	err = request.ParseMultipartForm(1024 * 1024 * 1024) // 10 GB limit
	if err != nil {
		RegError(c, writer, graph, &middleware.ServerError{StatusCode: 400, Message: fmt.Sprintf("Error parsing form: %s", err)})
		return
	}
	file, _, err := request.FormFile("file")
	if err != nil {
		RegError(c, writer, graph, &middleware.ServerError{StatusCode: 400, Message: fmt.Sprintf("failed to parse attached file: %s", err)})
		return
	}
	defer file.Close()
	var reader io.Reader = file

	conn, err := gripql.Connect(rpc.ConfigWithDefaults(host), true)
	wait := make(chan bool)

	VertChan, warnChan := streamJsonFromReader(reader, graph, mapExtraArgs, 5)

	go func() {
		for warning := range warnChan {
			warnings = append(warnings, warning)
		}
	}()

	go func() {
		err, res = conn.BulkAddRaw(VertChan)
		if err != nil {
			log.Errorf("Internal Server Error: %v", err)
		}
		wait <- false
	}()
	<-wait

	if err != nil {
		RegError(c, writer, graph, &middleware.ServerError{StatusCode: 500, Message: fmt.Sprintf("BulkAddRaw error: %v", err)})
		return
	}

	nonLoadedEdges := append(res.Errors, warnings...)
	if len(nonLoadedEdges) > 0 {
		if res.InsertCount == 0 {
			RegError(c, writer, graph, &middleware.ServerError{StatusCode: 500, Message: fmt.Sprintln("[500] bulk-load-raw", nonLoadedEdges)})
			return
		}
		c.AbortWithStatusJSON(206, gin.H{
			"status":  206,
			"message": nonLoadedEdges,
			"data":    nil,
		})
		return
	}
	Response(c, writer, graph, nil, 200, "[200] bulk-load-raw")
}

func (gh *Handler) BulkStream(c *gin.Context) {
	writer, request, graph := getFields(c)
	host := "localhost:8202"
	var logRate = 10000

	err := request.ParseMultipartForm(1024 * 1024 * 1024) // 10 GB limit
	if err != nil {
		RegError(c, writer, graph, &middleware.ServerError{StatusCode: 400, Message: fmt.Sprintf("Error parsing form: %s", err)})
		return
	}

	types, ok := request.MultipartForm.Value["types"]
	if !ok || len(types) == 0 {
		RegError(c, writer, graph, &middleware.ServerError{StatusCode: 400, Message: fmt.Sprintf("types field not found in form: %s", err)})
		return
	}
	request_type := types[0]

	// Get the file from the form data
	file, handler, err := request.FormFile("file")
	if err != nil {
		RegError(c, writer, graph, &middleware.ServerError{StatusCode: 400, Message: fmt.Sprintf("failed to parse attached file: %s", err)})
		return
	}
	defer file.Close()
	var reader io.Reader = file

	if strings.HasSuffix(handler.Filename, ".gz") {
		gzReader, err := gzip.NewReader(file)
		if err != nil {
			RegError(c, writer, graph, &middleware.ServerError{StatusCode: 500, Message: fmt.Sprintf("Unable to create gzip reader for file %s: err: %s", handler.Filename, err)})
			return
		}
		defer gzReader.Close()
		reader = gzReader
	}

	conn, err := gripql.Connect(rpc.ConfigWithDefaults(host), true)
	elemChan := make(chan *gripql.GraphElement)
	wait := make(chan bool)
	go func() {
		if err := conn.BulkAdd(elemChan); err != nil {
			log.Errorf("bulk add error: %v", err)
		}
		wait <- false
	}()

	if request_type == "vertex" {
		log.Infof("Loading vertex file: %s", handler.Filename)
		VertChan, err := StreamVerticesFromReader(reader, 5)
		if err != nil {
			RegError(c, writer, graph, GetInternalServerErr(err))
			return
		}
		count := 0
		for v := range VertChan {
			count++
			if count%logRate == 0 {
				log.Infof("Loaded %d vertices", count)
			}
			elemChan <- &gripql.GraphElement{Graph: graph, Vertex: v}
		}
		log.Infof("Loaded total of %d vertices", count)
	}
	if request_type == "edge" {
		log.Infof("Loading edge file: %s", handler.Filename)
		EdgeChan, err := StreamEdgesFromReader(reader, 5)
		if err != nil {
			RegError(c, writer, graph, GetInternalServerErr(err))
			return
		}
		count := 0
		for e := range EdgeChan {
			count++
			if count%logRate == 0 {
				log.Infof("Loaded %d vertices", count)
			}
			elemChan <- &gripql.GraphElement{Graph: graph, Edge: e}
		}
		log.Infof("Loaded total of %d edges", count)
	}

	close(elemChan)
	<-wait

	Response(c, writer, graph, nil, 200, fmt.Sprintf("[200] bulk-stream on file: %s", handler.Filename))
}
