package graph

import (
	"fmt"
	"strings"

	"github.com/bmeg/grip/gripql"
)

func applyFilters(query **gripql.Query, args map[string]any) error {
	//Todo: support "accessiblity", "format", "sort":
	if filter, ok := args["filter"]; ok {
		chainedFilter, err := applyJsonFilter(filter.(map[string]any))
		if err != nil {
			return err
		}
		fmt.Printf("CHAINED FILTER: %s\n", chainedFilter.String())
		*query = (*query).Has(chainedFilter)
	}
	if first, ok := args["first"]; ok {
		if first.(*int) == nil {
			*query = (*query).Limit(uint32(10))
		} else {
			*query = (*query).Limit(uint32(*first.(*int)))
		}
	}
	if offset, ok := args["offset"]; ok {
		if offset.(*int) != nil {
			*query = (*query).Skip(uint32(*offset.(*int)))
		}
	}
	return nil
}

func applyJsonFilter(filter map[string]any) (*gripql.HasExpression, error) {
	topLevelOp := ""
	for key := range filter {
		topLevelOp = key
		break
	}
	topLevelOpLowerCase := strings.ToLower(topLevelOp)

	switch topLevelOpLowerCase {
	case "and", "or":
		var expressions []*gripql.HasExpression
		for _, item := range filter[topLevelOp].([]any) {
			itemObj, ok := item.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("invalid nested filter structure")
			}
			subExpr, err := applyJsonFilter(itemObj)
			if err != nil {
				return nil, err
			}
			expressions = append(expressions, subExpr)
		}

		if len(expressions) == 1 {
			return expressions[0], nil
		} else if len(expressions) > 1 {
			if topLevelOpLowerCase == "and" {
				return gripql.And(expressions...), nil
			} else {
				return gripql.Or(expressions...), nil
			}
		} else {
			return nil, fmt.Errorf("no valid expressions for logical operator: %s", topLevelOp)
		}

	default:
		field := ""
		topFilter, ok := filter[topLevelOp].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("Top level filter %s not of type map[string]any", filter[topLevelOp])
		}

		for key := range topFilter {
			field = key
			break
		}

		hasExpr, err := mapGraphQLOperatorToGrip(field, topFilter[field], topLevelOp)
		if err != nil {
			return nil, err
		}

		return hasExpr, nil
	}
}

func mapGraphQLOperatorToGrip(field string, value any, op string) (*gripql.HasExpression, error) {
	switch strings.ToLower(op) {
	case "eq", "=":
		return gripql.Eq(field, value), nil
	case "neq", "!=":
		return gripql.Neq(field, value), nil
	case "lt", "<":
		return gripql.Lt(field, value), nil
	case "gt", ">":
		return gripql.Gt(field, value), nil
	case "gte", ">=":
		return gripql.Gte(field, value), nil
	case "lte", "<=":
		return gripql.Lte(field, value), nil
	case "in":
		return gripql.Within(field, value), nil
	default:
		return nil, fmt.Errorf("Operator %s does not match any of the known operators\n", op)
	}
}
