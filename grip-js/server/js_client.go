package server

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/gripql/inspect"
	"github.com/google/uuid"

	gripqljs "github.com/bmeg/grip/gripql/javascript"
	"github.com/bmeg/grip/log"
	"github.com/dop251/goja"
	"google.golang.org/protobuf/encoding/protojson"
)

type JSClientWrapper struct {
	vm     *goja.Runtime
	client gripql.Client
	query  goja.Callable
	graph  string

	auth bool
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
	fmt.Println("QUERY RESULT: ", qr)
	if v := qr.GetVertex(); v != nil {
		data := v.GetDataMap()
		data["id"] = v.Id
		return data
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
	obj := args.Export()
	queryJSON, err := json.Marshal(obj)
	if err != nil {
		log.Infof("Error: %s\n", err)
		return nil
	}

	query := gripql.GraphQuery{}

	err = protojson.Unmarshal(queryJSON, &query)
	if err != nil {
		log.Errorf("unmarshal error: %s", err)
	}
	query.Graph = cw.graph

	FilteredGS, RemainingGS, CachedGS := []*gripql.GraphStatement{}, []*gripql.GraphStatement{}, []*gripql.GraphStatement{}
	if cw.auth {
		ResourceList := cw.vm.Get("ResourceList").Export()
		log.Debugln("Resource List: ", ResourceList)
		Header := cw.vm.Get("Header").Export().(any)
		ctx := context.WithValue(context.Background(), "Header", Header)
		ctx = context.WithValue(ctx, "ResourceList", ResourceList)

		Has_Statement := &gripql.GraphStatement{Statement: &gripql.GraphStatement_Has{Has: gripql.Within("auth_resource_path", ResourceList.([]any)...)}}
		steps := inspect.PipelineSteps(query.Query)
		//step_value := 0
		for i, v := range query.Query {
			steps_index, _ := strconv.Atoi(steps[i])
			if i > steps_index {
				RemainingGS = append(RemainingGS, v)
			}
			if i == steps_index {
				FilteredGS = append(FilteredGS, v, Has_Statement)
				CachedGS = append(CachedGS, v, Has_Statement)
			} else {
				if i == 0 {
					CachedGS = append(CachedGS, v)
				}
				FilteredGS = append(FilteredGS, v)
			}
		}
		query.Query = FilteredGS
	}
	/*
		log.Infof("Getting cached job")
		resultChan, err := cw.GetCachedJob(query, CachedGS, RemainingGS)
		if err != nil {
			log.Infof("Error: %s\n", err)
			return nil
		}
		if resultChan != nil {
			cachedOut := []any{}
			for row := range resultChan {
				cachedOut = append(cachedOut, cw.vm.ToValue(toInterface(row)))
			}
			return cw.vm.ToValue(cachedOut)
		}
	*/
	res, err := cw.client.Traversal(context.Background(), &gripql.GraphQuery{Graph: query.Graph, Query: query.Query})
	if err != nil {
		log.Infof("Traversal Error: %s\n", err)
		return nil
	}

	out := []any{}
	for row := range res {
		result := toInterface(row)
		out = append(out, result)
	}
	log.Debugf("Returning value: %#v", out)
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

func (cw *JSClientWrapper) AddVertex(args ...goja.Value) goja.Value {
	log.Infof("addVertex %s", args)

	_id := ""
	if args[0] != nil {
		g := args[0].Export()
		if gstr, ok := g.(string); ok {
			_id = gstr
		}
	}
	if _id == "" {
		_id = uuid.New().String()
	}

	log.Debugf("getting label")
	_label := ""
	l := args[1].Export()
	if lstr, ok := l.(string); ok {
		_label = lstr
	}

	vData := map[string]any{}
	data := jsExport(args[2])
	if data != nil {
		if jData, ok := data.(map[string]any); ok {
			vData = jData
		}
	}

	log.Debugf("ID: %s LABEL: %s", _id, _label)
	vertex := &gripql.Vertex{
		Id:    _id,
		Label: _label,
	}
	vData["id"] = _id
	vertex.SetDataMap(vData)

	log.Debugf("adding vertex: %#v graph: %s", vertex, cw.graph)
	err := cw.client.AddVertex(cw.graph, vertex)
	if err != nil {
		log.Errorf("error adding vertex: %s", err)
	}
	log.Debugf("Added vertex")
	return cw.vm.ToValue(_id)
}

func (cw *JSClientWrapper) toValue() goja.Value {
	return cw.vm.ToValue(cw)
}

func GetJSClient(graph string, client gripql.Client, vm *goja.Runtime, auth bool) (*JSClientWrapper, error) { // ctx context.Context
	gripqljs, _ := gripqljs.Asset("gripql.js")
	vm.RunString(string(gripqljs))

	qVal := vm.Get("query")
	query, _ := goja.AssertFunction(qVal)

	myWrapper := &JSClientWrapper{vm, client, query, graph, auth}
	return myWrapper, nil
}
