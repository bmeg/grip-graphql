package graph

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/bmeg/grip/gripql"
)

type Resolver struct {
	GripDb gripql.Client
	Schema *ast.Schema
}

type objectMap struct {
	edgeLabel   map[string]map[string]struct{} // Maps vertex labels to edge names
	edgeDstType map[string]map[string]string   // Maps vertex labels and edge names to destination labels
}

func (rt *renderTree) NewElement() string {
	rName := fmt.Sprintf("f%d", len(rt.rNameTree))
	rt.rNameTree[rName] = []string{}
	return rName
}

type renderTree struct {
	prevName  string
	moved     bool
	rNameTree map[string][]string
}

func queryBuild(query **gripql.Query, selSet ast.SelectionSet, curElement string, rt *renderTree, parentPath string) {
	// Recursively traverses AST and builds grip query and render field tree
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
				rt.rNameTree[curElement] = append(rt.rNameTree[curElement], newParentPath)
			} else {
				queryBuild(query, sel.SelectionSet, curElement, rt, newParentPath)
			}
		case *ast.InlineFragment:
			elem := rt.NewElement()
			*query = (*query).OutNull(rt.prevName + "_" + sel.TypeCondition[:len(sel.TypeCondition)-4]).As(elem)
			queryBuild(query, sel.SelectionSet, elem, rt, "")
			rt.moved = true
		default:
			panic(fmt.Errorf("unsupported type: %T", sel))
		}
	}
}

func (r *queryResolver) GetSelectedFieldsAst(ctx context.Context, sourceType string) {
	resctx := graphql.GetFieldContext(ctx)
	rt := &renderTree{
		rNameTree: map[string][]string{"f0": []string{}},
	}
	q := gripql.V().HasLabel(sourceType[:len(sourceType)-4]).As("f0")

	queryBuild(&q, resctx.Field.Selections, "f0", rt, "")
	fmt.Println("QUERY AFTER: ", q)
	fmt.Printf("RNAME TREE: %#v\n", rt.rNameTree)

	/*render := map[string]any{}
	for _, i := range rt.fields {
		fmt.Println("I: ", i, "FIELDNAME: ", rt.fieldName)
		render[i+"_gid"] = "$" + i + "._gid"
		render[i+"_data"] = "$" + i //+ "._data"
	}

	fmt.Printf("RENDER: %#v\n", render)
	q = q.Render(render)*/

	_, err := r.GripDb.Traversal(context.Background(), &gripql.GraphQuery{Graph: "CALIPER", Query: q.Statements})
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
	}

	//out := []any{}
	//for r := range result {
	//	values := r.GetRender().GetStructValue().AsMap()
	//	fmt.Println("VALUES: ", values)
	//}

	/*data := map[string]map[string]any{}
	for _, r := range rt.fields {
		v := values[r+"_data"]
		if d, ok := v.(map[string]any); ok {
			d["id"] = values[r+"_gid"]
			if d["id"] != "" {
				data[r] = d
			}
		}
	}
	for _, r := range rt.fields {
		fmt.Println("RT PARENT: ", rt.parent, "R: ", r)
		if parent, ok := rt.parent[r]; ok {
			fieldName := rt.fieldName[r]
			if data[r] != nil {
				fmt.Println("HELLO 2 DATA?", data)
				fmt.Println("HELLO 2 PARENT?", parent)
				fmt.Println("HELLO 2 FIELD NAME?", fieldName)
				fmt.Println("HELLO 2 data[r]?", data[r])
				data[parent][fieldName] = []any{data[r]}
			}
		}
	}
	fmt.Println("DATA: ", data)*/
	//out = append(out, data["f0"])

}
