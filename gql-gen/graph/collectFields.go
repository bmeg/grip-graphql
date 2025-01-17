package graph

import (
	"context"
	"fmt"
	"os"

	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/log"
	"github.com/vektah/gqlparser/v2/ast"
)

type Resolver struct {
	Graph  string
	GripDb gripql.Client
	Schema *ast.Schema
}

type renderTree struct {
	prevName          string
	moved             bool
	rFieldPaths       map[string][]string
	rTree             map[string]interface{}
	rPotentialUnwinds []string
	rActualUnwinds    map[string][]string
}

func (rt *renderTree) NewElement() string {
	rName := fmt.Sprintf("f%d", len(rt.rFieldPaths))
	rt.rFieldPaths[rName] = []string{}
	return rName
}

func containedinSubstr(pUnwinds []string, path string) bool {
	paths := strings.Split(path, ".")

	for _, unwind := range pUnwinds {
		unwindParts := strings.Split(unwind, ".")
		j := 0
		for i := 0; i < len(unwindParts) && j < len(paths); i++ {
			if unwindParts[i] == paths[j] {
				j++
			}
		}
		if j == len(paths) {
			return true
		}
	}
	return false
}

func queryBuild(query **gripql.Query, selSet ast.SelectionSet, curElement string, rt *renderTree, parentPath string, currentTree map[string]any) {
	// Recursively traverses AST and build grip query, renders field tree
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
					rPath := rt.rFieldPaths[curElement]
					rPath = append(rPath, firstTerm)
					rt.rFieldPaths[curElement] = rPath
				}
				currentTree[curElement] = rt.rFieldPaths[curElement]
			} else {
				fmt.Println("UNWINDS: ", rt.rPotentialUnwinds, "NEWPPATH", newParentPath, "NODE: ", len(rt.rFieldPaths))
				if sel.Definition.Type.Elem != nil && containedinSubstr(rt.rPotentialUnwinds, newParentPath) {
					fval := fmt.Sprintf("f%d", len(rt.rFieldPaths)-1)
					aunwinds := rt.rActualUnwinds[fval]
					aunwinds = append(aunwinds, newParentPath)
					rt.rActualUnwinds[fval] = aunwinds
					*query = (*query).Unwind(newParentPath)
					*query = (*query).As(fval)
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
	unwinds := []string{}
	var err error
	if filter, ok := resctx.Args["filter"].(map[string]any); ok && filter != nil {
		unwinds, err = getUnwinds(filter)
		if err != nil {
			return nil, err
		}
	}
	rt := &renderTree{
		rActualUnwinds:    map[string][]string{},
		rPotentialUnwinds: unwinds,
		rFieldPaths:       map[string][]string{},
		rTree:             map[string]any{},
	}
	q := gripql.V().HasLabel(sourceType[:len(sourceType)-4]).As("f0")
	queryBuild(&q, resctx.Field.Selections, "f0", rt, "", rt.rTree)
	fmt.Println("ACTUAL UNWINDS: ", rt.rActualUnwinds)

	log.Infof("RNAME TREE: %#v\n", rt.rFieldPaths)
	log.Infof("R TREE: %#v\n", rt.rTree)

	renderTree := make(map[string]any, len(rt.rTree))
	buildRenderTree(renderTree, rt.rTree)

	log.Infof("ARGS: %#v\n", resctx.Args)
	log.Infof("RENDER: \n", renderTree)
	q = q.Select("f0")

	if filter, ok := resctx.Args["filter"]; ok {
		if filter != nil && len(filter.(map[string]any)) > 0 {
			err := rt.applyFilters(&q, filter.(map[string]any))
			if err != nil {
				return nil, err
			}
		}
	}

	if os.Getenv("AUTH_ENABLED") == "true" {
		authList, ok := ctx.Value("auth_list").([]interface{})
		if !ok {
			return nil, fmt.Errorf("auth_list not found or invalid")
		}
		applyAuthFilters(&q, authList)
	}

	// apply default filters after main filters so that all data can be considered in filter before apply filter statements
	applyDefaultFilters(&q, resctx.Args)

	q = q.Render(renderTree)
	log.Infoln("QUERY AFTER RENDER: ", q)

	result, err := r.GripDb.Traversal(context.Background(), &gripql.GraphQuery{Graph: r.Graph, Query: q.Statements})
	if err != nil {
		return nil, fmt.Errorf("Traversal Error: %s", err)
	}

	log.Infoln("QUERY AFTER TRAVERSAL")

	out := []any{}
	for r := range result {
		out = append(out, r.GetRender().GetStructValue().AsMap())
	}

	log.Infoln("QUERY AFTER RESULT")

	return out, nil
}

func buildRenderTree(output map[string]any, renderTree map[string]any) {
	/* Build the render tree to be used in grip render step */
	for key, val := range renderTree {
		switch v := val.(type) {
		case []string:
			for _, fieldPath := range v {
				current := output
				renderKey := "$" + key + "." + fieldPath
				if next, exists := current[renderKey]; exists {
					current = next.(map[string]any)
				} else {
					current[fieldPath] = renderKey
				}
			}
		case map[string]any:
			subTree := make(map[string]any)
			buildRenderTree(subTree, v)
			output[key] = subTree
		case string:
			output[key] = val
		default:
			fmt.Printf("Unexpected type: %T\n", val)
		}
	}
}
