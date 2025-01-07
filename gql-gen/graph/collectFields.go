package graph

import (
	"context"
	"fmt"
	"os"

	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/gripql/inspect"
	"github.com/vektah/gqlparser/v2/ast"
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

	render := make(map[string]any, len(rt.rFieldPaths)*2)
	for checkpoint, paths := range rt.rFieldPaths {
		checkpointPrefix := "$" + checkpoint + "."
		render[checkpoint+"_gid"] = checkpointPrefix + "id"
		for _, path := range paths.path {
			render[path+"_data"] = checkpointPrefix + path
		}
	}

	q = q.Select("f0")
	fmt.Printf("ARGS: %#v\n", resctx.Args)

	if filter, ok := resctx.Args["filter"]; ok {
		if filter != nil && len(filter.(map[string]any)) > 0 {
			err := rt.applyFilters(&q, filter.(map[string]any))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("filter is specified but filter not populated in variables")
		}
	}

	if first, ok := resctx.Args["first"]; ok {
		firstPtr, _ := first.(*int)
		if firstPtr == nil {
			q = q.Limit(uint32(10))
		} else {
			q = q.Limit(uint32(*firstPtr))
		}
	}
	if offset, ok := resctx.Args["offset"]; ok {
		if offset.(*int) != nil {
			q = q.Skip(uint32(*offset.(*int)))
		}
	}

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
	fmt.Println("QUERY AFTER RENDER: ", q)

	result, err := r.GripDb.Traversal(context.Background(), &gripql.GraphQuery{Graph: "CALIPER", Query: q.Statements})
	if err != nil {
		return nil, fmt.Errorf("Traversal Error: %s", err)
	}

	cachedTree := make(map[string]any, len(rt.rTree))
	buildTreeStructure(cachedTree, rt.rTree)

	out := []any{}
	for r := range result {
		values := r.GetRender().GetStructValue().AsMap()
		data := buildFilteredResponseTree(cachedTree, values)
		out = append(out, data)
	}
	return out, nil
}

func buildTreeStructure(output map[string]any, renderTree map[string]any) {
	for key, val := range renderTree {
		switch v := val.(type) {
		case renderTreePath:
			for _, fieldPath := range v.path {
				segments := strings.Split(fieldPath, ".")
				current := output
				for i := 0; i < len(segments)-1; i++ {
					segment := segments[i]
					if next, exists := current[segment]; exists {
						current = next.(map[string]any)
					} else {
						newMap := make(map[string]any, len(segments)-i-1)
						current[segment] = newMap
						current = newMap
					}
				}
				lastSegment := segments[len(segments)-1]
				current[lastSegment] = nil
			}
		case map[string]any:
			subTree := make(map[string]any)
			buildTreeStructure(subTree, v)
			output[key] = subTree
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
}

func buildFilteredResponseTree(cachedTree, values map[string]any) map[string]any {
	output := make(map[string]any, len(cachedTree))
	for key, val := range cachedTree {
		switch v := val.(type) {
		case map[string]any:
			if filteredSubTree := buildFilteredResponseTree(v, values); len(filteredSubTree) > 0 {
				output[key] = filteredSubTree
			}
		case nil:
			if renderedValue, exists := values[key+"_data"]; exists {
				output[key] = renderedValue
			}
		default:
			output[key] = val
		}
	}
	return output
}
