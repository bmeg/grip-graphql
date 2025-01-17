package graph

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/gripql/inspect"
)

func applyAuthFilters(q **gripql.Query, authList []any) {
	Has_Statement := &gripql.GraphStatement{Statement: &gripql.GraphStatement_Has{gripql.Within("auth_resource_path", authList...)}}
	steps := inspect.PipelineSteps((*q).Statements)
	FilteredGS := []*gripql.GraphStatement{}
	step_value := 0

	for i, v := range (*q).Statements {
		steps_index, _ := strconv.Atoi(steps[i])
		if i == 0 || steps_index <= step_value {
			FilteredGS = append(FilteredGS, v)
		} else {
			FilteredGS = append(FilteredGS, v, Has_Statement)
			step_value++
		}
	}
	(*q).Statements = FilteredGS
}

func applyDefaultFilters(q **gripql.Query, args map[string]any) {
	if first, ok := args["first"]; ok {
		firstPtr, _ := first.(*int)
		if firstPtr == nil {
			*q = (*q).Limit(uint32(10))
		} else {
			*q = (*q).Limit(uint32(*firstPtr))
		}
	}
	if offset, ok := args["offset"]; ok {
		if offset.(*int) != nil {
			*q = (*q).Skip(uint32(*offset.(*int)))
		}
	}
}

func (rt *renderTree) applyRewinds(query **gripql.Query) {
	for f, paths := range rt.rActualUnwinds {
		*query = (*query).Select(f)
		// sort fields so that toType operations are done first then groups
		sort.Slice(paths, func(i, j int) bool {
			return len(strings.Split(paths[i], "."))-len(strings.Split(paths[j], ".")) > 0
		})
		for _, path := range paths {
			jsonPath := "$" + f + "." + path
			if !strings.Contains(path, ".") {
				*query = (*query).Group(map[string]string{path: jsonPath})
			} else {
				*query = (*query).ToType(jsonPath, "list")
			}
		}
		*query = (*query).As(f)
	}
}

func (rt *renderTree) applyFilters(query **gripql.Query, filter map[string]any) error {
	//Todo: support "sort" operations
	chainedFilter, err := applyJsonFilter(filter)
	if err != nil {
		return err
	}

	*query = (*query).Has(chainedFilter)
	rt.applyRewinds(query)

	return nil
}

func getUnwinds(filter map[string]any) ([]string, error) {
	/* Returns a list of fields that may need to be unwound so that query builder can unwind as it builds the query */
	topLevelOp := ""
	for key := range filter {
		topLevelOp = key
		break
	}
	topLevelOpLowerCase := strings.ToLower(topLevelOp)

	switch topLevelOpLowerCase {
	case "and", "or":
		var fieldPaths []string
		for _, item := range filter[topLevelOp].([]any) {
			itemObj, ok := item.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("invalid nested filter structure")
			}
			subFieldPaths, err := getUnwinds(itemObj)
			if err != nil {
				return nil, err
			}
			fieldPaths = append(fieldPaths, subFieldPaths...)
		}
		return fieldPaths, nil

	default:
		field := ""
		topFilter, ok := filter[topLevelOp].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid logical operator '%s'", topLevelOp)
		}
		for key := range topFilter {
			field = key
			break
		}
		return []string{field}, nil
	}
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
			// If here then format is correct but logical operator is not supported
			return nil, fmt.Errorf("invalid logical operator '%s'", topLevelOp)
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
