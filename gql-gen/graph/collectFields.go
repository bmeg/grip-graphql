package graph

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/bmeg/grip/gripql"
	"github.com/vektah/gqlparser/v2/ast"
)

type Resolver struct {
	GripDb gripql.Client
	Schema *ast.Schema
}

type renderTree struct {
	prevName    string
	moved       bool
	rFieldPaths map[string][]string
	rTree       map[string]interface{}
}

func (rt *renderTree) NewElement() string {
	rName := fmt.Sprintf("f%d", len(rt.rFieldPaths))
	rt.rFieldPaths[rName] = []string{}
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
				for _, term := range rt.rFieldPaths[curElement] {
					if term == firstTerm {
						exists = true
						break
					}
				}
				if !exists {
					rt.rFieldPaths[curElement] = append(rt.rFieldPaths[curElement], firstTerm)
				}
				//rt.rFieldPaths[curElement] = append(rt.rFieldPaths[curElement], newParentPath)
				currentTree[curElement] = rt.rFieldPaths[curElement]
			} else {
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
		rFieldPaths: map[string][]string{"f0": []string{}},
		rTree:       map[string]any{},
	}
	q := gripql.V().HasLabel(sourceType[:len(sourceType)-4]).As("f0")
	queryBuild(&q, resctx.Field.Selections, "f0", rt, "", rt.rTree)

	fmt.Println("QUERY AFTER: ", q)
	fmt.Printf("RNAME TREE: %#v\n", rt.rFieldPaths)
	fmt.Printf("R TREE: %#v\n", rt.rTree)

	render := map[string]any{}
	for checkpoint, paths := range rt.rFieldPaths {
		render[checkpoint+"_gid"] = "$" + checkpoint + ".id"
		for _, path := range paths {
			render[path+"_data"] = "$" + checkpoint + "." + path
		}
	}

	//fmt.Printf("RENDER: %#v\n", render)
	q = q.Limit(10).Render(render)

	result, err := r.GripDb.Traversal(context.Background(), &gripql.GraphQuery{Graph: "CALIPER", Query: q.Statements})
	if err != nil {
		return nil, fmt.Errorf("Traversal Error: %s", err)
	}

	out := []any{}
	for r := range result {
		values := r.GetRender().GetStructValue().AsMap()
		//fmt.Printf("VALUES: %#v\n", values)
		data := buildOutputTree(rt.rTree, values)
		fmt.Printf("DATA: %#v\n", data)
		/*if entry, ok := data["focus"]; ok {
		entry.(map[string]any)["__typename"] = "SpecimenType"
		}*/
		out = append(out, data)
	}
	return out, nil
}

func buildOutputTree(renderTree map[string]interface{}, values map[string]interface{}) map[string]interface{} {
	output := map[string]interface{}{}
	for key, val := range renderTree {
		switch v := val.(type) {
		case []string:
			for _, fieldPath := range v {
				segments := strings.Split(fieldPath, ".")
				current := output
				for i := 0; i < len(segments)-1; i++ {
					//fmt.Println("CURRENT: ", current, "SEGMENTS[i]", segments[i])
					//fmt.Println("VALUES: ", values)
					if _, exists := current[segments[i]]; !exists {
						current[segments[i]] = map[string]interface{}{}
					}
					current = current[segments[i]].(map[string]interface{})
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
