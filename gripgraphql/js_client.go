package gripgraphql

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/bmeg/grip/gripql"
	gripqljs "github.com/bmeg/grip/gripql/javascript"
	"github.com/dop251/goja"
	"google.golang.org/protobuf/encoding/protojson"
)

type jsClientWrapper struct {
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

func (cw *jsClientWrapper) ToList(args goja.Value) goja.Value {

	obj := args.Export()

	queryJSON, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	fmt.Printf("%s\n", queryJSON)
	query := gripql.GraphQuery{}
	err = protojson.Unmarshal(queryJSON, &query)
	query.Graph = cw.graph
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	out := []any{}
	res, err := cw.client.Traversal(&query)
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

func (cw *jsClientWrapper) V(args goja.Value) goja.Value {
	gRes, err := cw.query(goja.Undefined(), cw.vm.ToValue(cw))
	if err != nil {
		return goja.Undefined()
	}
	gObj := gRes.ToObject(cw.vm)
	vObj := gObj.Get("V")
	vFunc, _ := goja.AssertFunction(vObj)
	out, _ := vFunc(gObj, args)
	//fmt.Printf("%s %s %s\n", args, gObj, out)
	return out
}

func GetJSClient(graph string, client gripql.Client, vm *goja.Runtime) (goja.Value, error) {
	//TODO: more error checking
	gripqljs, _ := gripqljs.Asset("gripql.js")
	vm.RunString(string(gripqljs))

	qVal := vm.Get("query")
	query, _ := goja.AssertFunction(qVal)

	myWrapper := &jsClientWrapper{vm, client, query, graph}
	clientWrapper := vm.ToValue(myWrapper)
	return clientWrapper, nil
}
