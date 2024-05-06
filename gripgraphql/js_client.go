package gripgraphql

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/gripql/inspect"

	//"github.com/bmeg/grip/log"
	gripqljs "github.com/bmeg/grip/gripql/javascript"
	"github.com/dop251/goja"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

type JSClientWrapper struct {
	vm     *goja.Runtime
	client gripql.Client
	query  goja.Callable
	graph  string
}

type JSRenamer struct{}

func (j JSRenamer) FieldName(t reflect.Type, f reflect.StructField) string {
	return f.Name
}

func (j JSRenamer) MethodName(t reflect.Type, m reflect.Method) string {
	if m.Name == "V" || m.Name == "E" {
		return m.Name
	}
	s := m.Name
	return strings.ToLower(s[0:1]) + s[1:]
}

func toInterface(qr *gripql.QueryResult) any {
	if v := qr.GetVertex(); v != nil {
		return v.GetDataMap()
	}
	if e := qr.GetEdge(); e != nil {
		return e.GetDataMap()
	}
	if c := qr.GetCount(); c != 0 {
		return c
	}
	if r := qr.GetRender(); r != nil {
		return r.AsInterface()
	}
	if a := qr.GetAggregations(); a != nil {
		return map[string]any{
			"key":   a.GetKey().AsInterface(),
			"value": a.GetValue(),
			"name":  a.GetName(),
		}
	}
	if p := qr.GetPath(); p != nil {
		pa := []any{}
		for _, c := range p.GetValues() {
			pa = append(pa, c.AsInterface())
		}
		return pa
	}
	return qr
}

func (cw *JSClientWrapper) ToList(args goja.Value) goja.Value {

	//fmt.Printf("ARGS: %s", args.String)
	obj := args.Export()
	//fmt.Printf("obj: %#v", args)

	queryJSON, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	ResourceList := cw.vm.Get("ResourceList").Export().([]any)
	// Testing to make sure passing empty list would filter out everything
	//var ResourceList []interface{} = []interface{}{}

	Header := cw.vm.Get("Header").Export().(any)
	ctx := context.WithValue(context.Background(), "Header", Header)
	ctx = context.WithValue(ctx, "ResourceList", ResourceList)

	// This hasn't gotten connected. Too slow
	query := gripql.GraphQuery{}
	err = protojson.Unmarshal(queryJSON, &query)
	sValue, _ := structpb.NewValue(ResourceList)
	Has_Statement := &gripql.GraphStatement{Statement: &gripql.GraphStatement_Has{
		Has: &gripql.HasExpression{Expression: &gripql.HasExpression_Condition{
			Condition: &gripql.HasCondition{
				Condition: gripql.Condition_WITHIN,
				Key:       "auth_resource_path",
				Value:     sValue,
			},
		}},
	}}

	// gripql.Within("auth_resource_path", ResourceList...)
	// query.Query = append(query.Query, Has_Statement) // .Has

	steps := inspect.PipelineSteps(query.Query)
	query.Graph = cw.graph
	FilteredGS := []*gripql.GraphStatement{}

	for i, v := range query.Query {
		//fmt.Println("V: ", v)
		steps_index, err := strconv.Atoi(steps[i])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return nil
		}
		//fmt.Println("I: ", i, "steps_index: ", steps_index)
		if i == steps_index {
			FilteredGS = append(FilteredGS, v, Has_Statement)
		} else {
			FilteredGS = append(FilteredGS, v)
		}
	}

	fmt.Println("FILTERED GS: ", FilteredGS)
	query.Query = FilteredGS

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	out := []any{}
	res, err := cw.client.Traversal(ctx, &query)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}

	for row := range res {
		out = append(out, cw.vm.ToValue(toInterface(row)))
	}
	fmt.Printf("ToList: %s\n", out)
	return cw.vm.ToValue(out)
}

func (cw *JSClientWrapper) V(args goja.Value) goja.Value {

	//ResourceList := cw.vm.Get("ResourceList")
	gRes, err := cw.query(goja.Undefined(), cw.vm.ToValue(cw))
	if err != nil {
		return goja.Undefined()
	}
	gObj := gRes.ToObject(cw.vm)
	vObj := gObj.Get("V")
	vFunc, _ := goja.AssertFunction(vObj)
	out, _ := vFunc(gObj, args)
	//fmt.Printf("JS LAND: %s %s %s\n", args, gObj, out)
	return out
}

func (cw *JSClientWrapper) toValue() goja.Value {
	return cw.vm.ToValue(cw)
}

func GetJSClient(graph string, client gripql.Client, vm *goja.Runtime) (*JSClientWrapper, error) { // ctx context.Context
	gripqljs, _ := gripqljs.Asset("gripql.js")
	vm.RunString(string(gripqljs))

	qVal := vm.Get("query")
	query, _ := goja.AssertFunction(qVal)

	myWrapper := &JSClientWrapper{vm, client, query, graph}
	return myWrapper, nil
}
