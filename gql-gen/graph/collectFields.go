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

type renderTree struct {
	prevName  string
	moved     bool
	fields    []string
	parent    map[string]string
	fieldName map[string]string
}

type objectMap struct {
	edgeLabel   map[string]map[string]struct{} // Maps vertex labels to edge names
	edgeDstType map[string]map[string]string   // Maps vertex labels and edge names to destination labels
}

func (rt *renderTree) NewElement(cur string, fieldName string) string {
	rName := fmt.Sprintf("f%d", len(rt.fields))
	rt.fields = append(rt.fields, rName)
	rt.parent[rName] = cur
	rt.fieldName[rName] = fieldName
	return rName
}

func traversalBuild(reqCtx *graphql.OperationContext, query **gripql.Query, selSet ast.SelectionSet, curElement string, rt *renderTree, visited map[string]bool) []graphql.CollectedField {
	fmt.Printf("\n\n")
	groupedFields := make([]graphql.CollectedField, 0, len(selSet))
	for _, s := range selSet {
		switch sel := s.(type) {
		case *ast.Field:
			if _, ok := visited[sel.Name]; ok {
				continue
			}
			visited[sel.Name] = true

			fmt.Printf("FIELD NAME: %s\n", sel.Name)
			fmt.Printf("FIELD DEF NAME %s\n", sel.ObjectDefinition.Name)
			if rt.moved {
				*query = (*query).Select(curElement)
				rt.moved = false
			}
			rt.prevName = sel.Name
			elem := rt.NewElement(sel.ObjectDefinition.Name, sel.Name)
			for _, childField := range traversalBuild(reqCtx, query, sel.SelectionSet, elem, rt, visited) {
				_ = rt.NewElement(sel.Name, childField.Name)
				fmt.Println("CUR NAME: ", sel.Name, "CHILD NAME: ", childField.Name)
			}
		case *ast.InlineFragment:
			elem := rt.NewElement(rt.prevName, sel.TypeCondition)

			fmt.Printf("INLINE FRAG Type CONDITION: %s\n", sel.TypeCondition)
			fmt.Printf("InlineFragment DEF NAME %s\n", sel.ObjectDefinition.Name)
			typeConditionLen := len(sel.TypeCondition)
			*query = (*query).OutNull(rt.prevName + "_" + sel.TypeCondition[:typeConditionLen-4]).As(elem)
			for _, childField := range traversalBuild(reqCtx, query, sel.SelectionSet, elem, rt, visited) {
				_ = rt.NewElement(rt.prevName, childField.Name)
				fmt.Println("InlineFragment CUR NAME: ", sel.ObjectDefinition.Name, "CHILD NAME: ", childField.Name)
			}
			rt.moved = true

		case *ast.FragmentSpread:
			fmt.Println("FRAG SPREAD: ", sel.Definition.Name)
		default:
			panic(fmt.Errorf("unsupported %T", sel))
		}

	}
	return groupedFields
}

func (r *queryResolver) GetSelectedFieldsAst(ctx context.Context, sourceType string) {
	resctx := graphql.GetFieldContext(ctx)
	opCtx := graphql.GetOperationContext(ctx)
	rt := &renderTree{
		fields:    []string{"f0"},
		parent:    map[string]string{},
		fieldName: map[string]string{},
	}
	q := gripql.V().HasLabel(sourceType[:len(sourceType)-4]).As("f0")
	//for _, field := range resctx.Field.Selections {
	_ = traversalBuild(opCtx, &q, resctx.Field.Selections, "f0", rt, map[string]bool{})
	fmt.Println("QUERY AFTER: ", q)
	fmt.Printf("RENDER TREE FIELDS: %#v\n", rt.fields)
	fmt.Printf("RENDER TREE PARENT: %#v\n", rt.parent)
	fmt.Printf("RENDER TREE FieldName: %#v\n", rt.fieldName)

	render := map[string]any{}
	for _, i := range rt.fields {
		render[i+"_gid"] = "$" + i + "._gid"
		render[i+"_data"] = "$" + i + "._data"
	}

	fmt.Printf("RENDER: %#v\n", render)
	q = q.Render(render)

	result, err := r.GripDb.Traversal(context.Background(), &gripql.GraphQuery{Graph: "CALIPER", Query: q.Statements})
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
	}

	out := []any{}
	for r := range result {
		values := r.GetRender().GetStructValue().AsMap()
		fmt.Println("VALUES: ", values)

		data := map[string]map[string]any{}
		for _, r := range rt.fields {
			v := values[r+"_data"]
			fmt.Println("V:", v)
			if d, ok := v.(map[string]any); ok {
				fmt.Println("HELLO IN HERE")
				d["id"] = values[r+"_gid"]
				fmt.Println("D: ", d)
				if d["id"] != "" {
					data[r] = d
				}
			}
		}
		for _, r := range rt.fields {
			fmt.Println("RT PARENT: ", rt.parent, "R: ", r)
			if parent, ok := rt.parent[r]; ok {
				fieldName := rt.fieldName[r]
				fmt.Println("DATA: ", data)
				if data[r] != nil {
					data[parent][fieldName] = []any{data[r]}
				}
			}
		}
		fmt.Println("DATA: ", data)
		out = append(out, data["f0"])
	}

}
