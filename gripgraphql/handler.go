package gripgraphql

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"github.com/dop251/goja"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type QueryField struct {
	field   *graphql.Field
	handler func(goja.FunctionCall) goja.Value
}

type GraphQLJS struct {
     Client    gripql.Client
     GjHandler *handler.Handler
     GraphName string
     Config    string
}

type Endpoint struct {
	client     gripql.Client
	vm         *goja.Runtime
	gObject    goja.Value
	queryNodes map[string]QueryField
}

func parseField(name string, x any) (*graphql.Field, error) {
	switch v := x.(type) {
	case string:
		if v == "Int" {
			return &graphql.Field{Name: name, Type: graphql.Int}, nil
		}
		if v == "String" {
			return &graphql.Field{Name: name, Type: graphql.String}, nil
		}
		if v == "Boolean" {
			return &graphql.Field{Name: name, Type: graphql.Boolean}, nil
		}
		if v == "Float" {
			return &graphql.Field{Name: name, Type: graphql.Float}, nil
		}
	case []any:
		if len(v) != 1 {
			return nil, fmt.Errorf("incorrect elements in schema array (only 1)")
		}
		if lf, err := parseField(name, v[0]); err == nil {
			log.Infof("Found array\n")
			l := graphql.NewList(lf.Type)
			return &graphql.Field{Name: name, Type: l}, nil
		} else {
			log.Errorf("Error parsing list: %s", err)
		}
	case map[string]any:
		obj, err := parseObject(name, v)
		if err != nil {
			return nil, err
		}
		return &graphql.Field{Name: name, Type: obj}, nil
	}

	return nil, fmt.Errorf("type not found: %#v", x)
}

func parseObject(name string, x map[string]any) (*graphql.Object, error) {
	fields := graphql.Fields{}
	for k, v := range x {
		f, err := parseField(k, v)
		if err == nil {
			fields[k] = f
		} else {
			return nil, err
		}
	}
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	}), nil
}

func (e *Endpoint) Add(x map[string]any) {
	name := ""
	if nameA, ok := x["name"]; ok {
		if nameStr, ok := nameA.(string); ok {
			name = nameStr
		}
	}
	if name == "" {
		log.Errorf("Name not defined")
		return
	}

	var jHandler func(goja.FunctionCall) goja.Value

	if handlerA, ok := x["handler"]; ok {
		if handler, ok := handlerA.(func(goja.FunctionCall) goja.Value); ok {
			jHandler = handler
		} else {
			log.Errorf("Unknown handler type: %#T\n", handlerA)
			return
		}
	}

	defaults := map[string]any{}
	if defaultsA, ok := x["defaults"]; ok {
		if defaultsM, ok := defaultsA.(map[string]any); ok {
			for k, v := range defaultsM {
				defaults[k] = v
			}
		}
	}

	arguments := map[string]any{}
	if argumentsA, ok := x["args"]; ok {
		if argumentsM, ok := argumentsA.(map[string]any); ok {
			for k, v := range argumentsM {
				arguments[k] = v
			}
		}
	}

	log.Infof("Loading query %s", name)
	if schemaA, ok := x["schema"]; ok {
		objField, err := parseField(name, schemaA)
		if err == nil {
			objField.Resolve = func(params graphql.ResolveParams) (interface{}, error) {
				//fmt.Printf("Calling resolver\n")
				uArgs := map[string]any{}
				for k, v := range defaults {
					uArgs[k] = v
				}
				for k, v := range params.Args {
					uArgs[k] = v
				}

				vArgs := e.vm.ToValue(uArgs)
                //fmt.Printf("e.gObject: %#v vArgs: %#v\n", e.gObject, vArgs)
				args := goja.FunctionCall{
					Arguments: []goja.Value{e.gObject, vArgs},
				}
                //fmt.Printf("ARGS: %#v\n", args)
				val := jHandler(args)
				out := jsExport(val)
				fmt.Printf("Handler returned: %#v\n", out)
				return out, nil
			}

			if len(arguments) > 0 {
				args := graphql.FieldConfigArgument{}
				for k, v := range arguments {
					if v == "String" {
						args[k] = &graphql.ArgumentConfig{Type: graphql.String} // DefaultValue: 100,
					}
					if v == "Int" {
						args[k] = &graphql.ArgumentConfig{Type: graphql.Int}
					}
					if v == "Boolean" {
						args[k] = &graphql.ArgumentConfig{Type: graphql.Boolean}
					}
				}
				objField.Args = args
			}
			e.queryNodes[name] = QueryField{
				objField, jHandler,
			}
			log.Infof("Added GraphQL query node : %s", name)
		} else {
			log.Errorf("Parse Error: %s", err)
		}
	} else {
		log.Errorf("Schema not found for %s", name)
	}
}

func jsExport(val goja.Value) any {
	o := val.Export()
	if oList, ok := o.([]any); ok {
		out := []any{}
		for _, i := range oList {
			if ov, ok := i.(goja.Value); ok {
				out = append(out, ov.Export())
			} else {
				out = append(out, i)
			}
		}
		return out
	}
	return o
}

func (e *Endpoint) Build() (*graphql.Schema, error) {

	qf := graphql.Fields{}
	for k, v := range e.queryNodes {
        //log.Infof("fields: %+v", v.field)
		qf[k] = v.field
	}

	queryObj := graphql.NewObject(graphql.ObjectConfig{Name: "Query", Fields: qf})

	schemaConfig := graphql.SchemaConfig{
		Query: queryObj,
	}

	// Setup the GraphQL schema based on the objects there have been created
	gqlSchema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, fmt.Errorf("graphql.NewSchema error: %v", err)
	}
	return &gqlSchema, nil
}

func NewHTTPHandler(client gripql.Client, config map[string]string) (http.Handler, error) {
	configPath := "config.js"
	graph := "testGraph"
	if c, ok := config["config"]; ok {
		configPath = c
	}
	if c, ok := config["graph"]; ok {
		graph = c
	}

	//fmt.Printf("Plugin config: %s\n", config)

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

    //fmt.Printf("DATA: %s", data)

	vm := goja.New()
	vm.SetFieldNameMapper(JSRenamer{})

	jsClient, err := GetJSClient(graph, client, vm)
	if err != nil {
		fmt.Printf("js error: %s\n", err)
	}
	e := Endpoint{queryNodes: map[string]QueryField{}, client: client, vm: vm, gObject: jsClient}

	vm.Set("endpoint", map[string]any{
		"add":     e.Add,
		"String":  "String",
		"Int":     "Int",
		"Float":   "Float",
		"Boolean": "Boolean",
	})

	vm.Set("print", fmt.Printf)

    // fmt.Println("string DATA: ", string(data))
    _, err = vm.RunString(string(data))
	if err != nil {
		return nil, err
	}

	schema, err := e.Build()
	if err != nil {
		log.Errorf("Error building Handler: %s", err)
		return nil, err
	}
	hnd := handler.New(&handler.Config{
		Schema: schema,
	})

    return &GraphQLJS{Client: client, GjHandler: hnd, GraphName: graph, Config: configPath}, nil
}

// Static HTML that links to Apollo GraphQL query editor
var sandBox = `
<div id="sandbox" style="position:absolute;top:0;right:0;bottom:0;left:0"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
 new window.EmbeddedSandbox({
   target: "#sandbox",
   // Pass through your server href if you are embedding on an endpoint.
   // Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
   initialEndpoint: window.location.href,
 });
 // advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
</script>`

func (gh *GraphQLJS) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Infof("Request: %s", request.URL.Path)
	if request.URL.Path == "" || request.URL.Path == "/" {
		writer.Write([]byte(sandBox))
		return
	}
    if gh.Gen3{ 
        data := &middleware.GraphQLJS{
             Client:    gh.Client,
             GjHandler: gh.GjHandler,
             Gen3:      gh.Gen3,
             GraphName: gh.GraphName,
             Config:    gh.Config,
        }
        gql := &middleware.GraphQLJSDATA{
            Data: data,
        }
        err, resourceList := gql.Setup(writer, request)
        //ts, _ := gh.client.GetTimestamp(gh.graph)

    if request.URL.Path == "/api" || request.URL.Path == "api" {
	    gh.GjHandler.ServeHTTP(writer, request)
	}
}
