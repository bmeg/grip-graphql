package graph

import (
	"context"
	"fmt"
	"strings"

	"github.com/bmeg/grip/gripql"
)

type Resolver struct {
	GripDb gripql.Client
}

func (r *queryResolver) gripQuery(fields []string, sourceType string) []any {
	sourceTypeLen := len(sourceType)
	fmt.Println("SOURCE TYPE: ", sourceType[:sourceTypeLen-4])
	query := gripql.V().HasLabel(sourceType[:sourceTypeLen-4]).As(sourceType)
	for i, field := range fields {
		fields[i] = sourceType + "." + field
	}
	traversal, typeGraph := constructTypeTraversal(fields)

	// Construct query from template traversal string
	for _, value := range traversal {
		if strings.HasPrefix(value, "OUTNULL_") {
			splitStr := strings.Split(value, "_")
			//fmt.Println("SPLIT STR: ", splitStr)
			query = query.OutNull(splitStr[1] + "_" + splitStr[2][:len(splitStr[2])-4]).As(splitStr[2])
		} else if strings.HasPrefix(value, "SELECT_") {
			//fmt.Println("VALUE SELECT : ", value[7:])
			query = query.Select(value[7:])
		}
	}

	render := map[string]any{}
	for typeKey, typeVals := range typeGraph {
		for _, val := range typeVals {
			render[typeKey+"_"+val] = "$" + typeKey + "." + val
		}
	}
	query = query.Render(render)
	fmt.Println("RENDER: ", render)

	fmt.Println("STATEMENTS: ", query.Statements)
	result, err := r.GripDb.Traversal(context.Background(), &gripql.GraphQuery{Graph: "CALIPER", Query: query.Statements})
	if err != nil {
		return nil
	}

	out := []any{}
	for r := range result {
		out = append(out, r.GetRender().GetStructValue().AsMap())
	}

	return out
}
