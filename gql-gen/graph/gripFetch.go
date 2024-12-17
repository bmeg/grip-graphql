package graph

import (
	"fmt"

	"github.com/bmeg/grip/gripql"
)

type Resolver struct {
	GripDb gripql.Client
}

func gripQuery(fields []string, sourceType string) *gripql.Query {
	query := gripql.V().HasLabel(sourceType).As("_" + sourceType)
	for i, field := range fields {
		fields[i] = sourceType + "." + field
	}
	typeGraph, traversal := constructTypeTraversal(fields)
	fmt.Println("TYPE GRAPH: ", typeGraph)
	fmt.Println("Traversal Path", traversal)
	return query
}
