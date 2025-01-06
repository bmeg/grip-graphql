package graph

import (
	"context"
	"fmt"
	"os"

	"reflect"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/gripql/inspect"
	"github.com/vektah/gqlparser/v2/ast"
	//"google.golang.org/protobuf/types/known/structpb"
)

type Resolver struct {
	GripDb gripql.Client
	Schema *ast.Schema
}

type renderTreePath struct {
	path       []string
	unwindPath []string
}

type renderTree struct {
	prevName    string
	moved       bool
	rFieldPaths map[string]renderTreePath
	rTree       map[string]interface{}
}

func (rt *renderTree) NewElement() string {
	rName := fmt.Sprintf("f%d", len(rt.rFieldPaths))
	rt.rFieldPaths[rName] = renderTreePath{path: []string{}, unwindPath: []string{}}
	return rName
}

func queryBuild(query **gripql.Query, selSet ast.SelectionSet, curElement string, rt *renderTree, parentPath string, currentTree map[string]any) {
	// Recursively traverses AST and builds grip query, renders field tree
	for _, s := range selSet {
		switch sel := s.(type) {
		case *ast.Field:
			rt.prevName = sel.Name
			newParentPath := parentPath + "." + sel.Name
			if parentPath == "" {
				newParentPath = sel.Name
			}
			if rt.moved {
				*query = (*query).Select(curElement)
				rt.moved = false
			}
			if sel.SelectionSet == nil {
				// Full render paths can be stored, but all that is needed is the top level key names.
				firstTerm := newParentPath
				if dotIndex := strings.Index(newParentPath, "."); dotIndex != -1 {
					firstTerm = newParentPath[:dotIndex]
				}
				exists := false
				for _, term := range rt.rFieldPaths[curElement].path {
					if term == firstTerm {
						exists = true
						break
					}
				}
				if !exists {
					rPath := rt.rFieldPaths[curElement]
					rPath.path = append(rPath.path, firstTerm)
					rt.rFieldPaths[curElement] = rPath
				}
				currentTree[curElement] = rt.rFieldPaths[curElement]
			} else {

				/*fmt.Printf("OBJ DEF: %#v\n", sel.ObjectDefinition)
				fmt.Printf("DEF TYPE: %#v\n", sel.Definition.Type)
				fmt.Printf("DEF TYPE ELEM: %#v\n", sel.Definition.Type.Elem)
				fmt.Println("PARENT PATH: ", newParentPath)*/
				if sel.Definition.Type.Elem != nil {
					rPath := rt.rFieldPaths[curElement]
					rPath.unwindPath = append(rPath.unwindPath, newParentPath)
					rt.rFieldPaths[curElement] = rPath
				}
				queryBuild(query, sel.SelectionSet, curElement, rt, newParentPath, currentTree)
			}
		case *ast.InlineFragment:
			elem := rt.NewElement()
			if _, exists := currentTree[rt.prevName]; !exists {
				currentTree[rt.prevName] = map[string]any{"__typename": sel.TypeCondition}
			}
			fragmentTree := currentTree[rt.prevName].(map[string]interface{})
			//[sel.TypeCondition].(map[string]interface{})
			*query = (*query).OutNull(rt.prevName + "_" + sel.TypeCondition[:len(sel.TypeCondition)-4]).As(elem)
			queryBuild(query, sel.SelectionSet, elem, rt, "", fragmentTree)
			rt.moved = true
		default:
			panic(fmt.Errorf("unsupported type: %T", sel))
		}
	}
}

func (r *queryResolver) GetSelectedFieldsAst(ctx context.Context, sourceType string) ([]any, error) {
	resctx := graphql.GetFieldContext(ctx)

	rt := &renderTree{
		rFieldPaths: map[string]renderTreePath{"f0": renderTreePath{path: []string{}, unwindPath: []string{}}},
		rTree:       map[string]any{},
	}
	q := gripql.V().HasLabel(sourceType[:len(sourceType)-4]).As("f0")
	queryBuild(&q, resctx.Field.Selections, "f0", rt, "", rt.rTree)

	fmt.Printf("RNAME TREE: %#v\n", rt.rFieldPaths)
	fmt.Printf("R TREE: %#v\n", rt.rTree)

	render := map[string]any{}
	for checkpoint, paths := range rt.rFieldPaths {
		render[checkpoint+"_gid"] = "$" + checkpoint + ".id"
		for _, path := range paths.path {
			render[path+"_data"] = "$" + checkpoint + "." + path
		}
	}
	q = q.Select("f0")
	applyUnwinds(&q, rt)
	q = q.As("f0")
	fmt.Printf("ARGS: %#v\n", resctx.Args)
	err := applyFilters(&q, resctx.Args)
	if err != nil {
		return nil, err
	}
	err = applyRewinds(&q, rt)
	q = q.As("f0")

	if os.Getenv("AUTH_ENABLED") == "true" {
		authList, ok := ctx.Value("auth_list").([]interface{})
		if !ok {
			return nil, fmt.Errorf("auth_list not found or invalid")
		}

		Has_Statement := &gripql.GraphStatement{Statement: &gripql.GraphStatement_Has{gripql.Within("auth_resource_path", authList...)}}
		steps := inspect.PipelineSteps(q.Statements)
		FilteredGS := []*gripql.GraphStatement{}
		for i, v := range q.Statements {
			steps_index, _ := strconv.Atoi(steps[i])
			if i == 0 {
				FilteredGS = append(FilteredGS, v)
				continue
			} else if i == steps_index {
				FilteredGS = append(FilteredGS, v, Has_Statement)
			} else {
				FilteredGS = append(FilteredGS, v)
			}
		}

		q.Statements = FilteredGS
	}
	q = q.Render(render)
	fmt.Println("QUERY AFTER: ", q)

	result, err := r.GripDb.Traversal(context.Background(), &gripql.GraphQuery{Graph: "CALIPER", Query: q.Statements})
	if err != nil {
		fmt.Println("HELLO WE HERE: ", err)
		return nil, fmt.Errorf("Traversal Error: %s", err)
	}

	out := []any{}
	for r := range result {
		values := r.GetRender().GetStructValue().AsMap()
		fmt.Printf("VALUES: %#v\n", values)
		data := buildOutputTree(rt.rTree, values)
		//fmt.Printf("DATA: %#v\n", data)
		out = append(out, data)
	}
	return out, nil
}

func buildOutputTree(renderTree map[string]interface{}, values map[string]interface{}) map[string]interface{} {
	output := map[string]interface{}{}
	for key, val := range renderTree {
		switch v := val.(type) {
		case renderTreePath:
			for _, fieldPath := range v.path {
				segments := strings.Split(fieldPath, ".")
				current := output
				for i := 0; i < len(segments)-1; i++ {
					//fmt.Println("CURRENT: ", current, "SEGMENTS[i]", segments[i])
					//fmt.Println("VALUES: ", values)
					if _, exists := current[segments[i]]; !exists {
						current[segments[i]] = map[string]interface{}{}
					}
					//current = current[segments[i]].(map[string]interface{})
				}
				lastSegment := segments[len(segments)-1]
				fieldKey := fieldPath + "_data"
				if renderedValue, exists := values[fieldKey]; exists {
					current[lastSegment] = renderedValue
					// if rendered value is string and render was not found return nil instead of string.
					if reflect.TypeOf("$f") == reflect.TypeOf(renderedValue) && strings.HasPrefix(renderedValue.(string), "$f") {
						current[lastSegment] = nil
					}
				} else {
					current[lastSegment] = nil
				}
			}
		case map[string]interface{}:
			output[key] = buildOutputTree(v, values)
		case string:
			if key == "__typename" {
				output[key] = val
			} else {
				fmt.Printf("Unexpected type: %T\n", val)
			}
		default:
			fmt.Printf("Unexpected type: %T\n", val)
		}
	}

	return output
}
