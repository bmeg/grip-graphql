package gripgraphql

import (
	"context"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/gripql/inspect"
	"github.com/google/uuid"

	//"github.com/bmeg/grip-graphql/middleware"
	//"github.com/bmeg/grip/jobstorage"
	gripqljs "github.com/bmeg/grip/gripql/javascript"
	"github.com/bmeg/grip/log"
	"github.com/dop251/goja"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
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

	var ctx context.Context
	FilteredGS, CachedGS, RemainingGS := []*gripql.GraphStatement{}, []*gripql.GraphStatement{}, []*gripql.GraphStatement{}

	if cw.auth {
		ResourceList := cw.vm.Get("ResourceList").Export()
		Header := cw.vm.Get("Header").Export().(any)
		ctx := context.WithValue(context.Background(), "Header", Header)
		ctx = context.WithValue(ctx, "ResourceList", ResourceList)

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
		steps := inspect.PipelineSteps(query.Query)
		for i, v := range query.Query {
			steps_index, err := strconv.Atoi(steps[i])
			if err != nil {
				log.Infof("Error: %s\n", err)
				return nil
			}
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
	} else {
		ctx = context.Background()
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

	log.Infof("Doing traversal")

	res, err := cw.client.Traversal(ctx, &query)
	if err != nil {
		log.Infof("Error: %s\n", err)
		return nil
	}

	out := []any{}
	for row := range res {
		out = append(out, cw.vm.ToValue(toInterface(row)))
	}
	//log.Infof("Returning value: %s\n", out)

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

	gid := ""
	if args[0] != nil {
		g := args[0].Export()
		if gstr, ok := g.(string); ok {
			gid = gstr
		}
	}
	if gid == "" {
		gid = uuid.New().String()
	}

	log.Info("getting label")
	label := ""
	l := args[1].Export()
	if lstr, ok := l.(string); ok {
		label = lstr
	}

	vData := map[string]any{}
	data := jsExport(args[2])
	if data != nil {
		if jData, ok := data.(map[string]any); ok {
			vData = jData
		}
	}

	vertex := &gripql.Vertex{
		Id:    gid,
		Label: label,
	}
	vertex.SetDataMap(vData)

	log.Infof("adding vertex: %s", vertex)
	err := cw.client.AddVertex(cw.graph, vertex)
	if err != nil {
		log.Errorf("error adding vertex: %s", err)
	}
	log.Infof("Added vertex")
	return cw.vm.ToValue(gid)
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
