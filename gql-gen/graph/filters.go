package graph

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/bmeg/grip-graphql/gql-gen/model"
	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/gripql/inspect"
	"github.com/bmeg/grip/log"
)

func applyAuthFilters(q **gripql.Query, authList []any) {
	/* Applies gen3 RBAC auth filters to the query */
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
	/* Applies rewinds negating unwinds so that output data list structures are preserved */
	for f, paths := range rt.rUnwinds {
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

func (rt *renderTree) applySort(q **gripql.Query, sortings []*model.SortInput) error {
	/*
		sort must be in form -- sort: [
		   {
		     field: "field",
			 descending: false
		   }]

		where field is a valid json path
	*/
	sortFields := gripql.Sorting{Fields: []*gripql.SortField{}}
	for _, sort := range sortings {
		path, err := rt.formatField(sort.Field)
		if err != nil {
			return err
		}
		sortFields.Fields = append(sortFields.Fields, &gripql.SortField{Field: path, Descending: *sort.Descending})
	}

	*q = (*q).Sort(sortFields.Fields)
	log.Infof("Sort Filter: %v\n", sortFields.Fields)

	return nil
}

func (rt *renderTree) applyFilters(query **gripql.Query, filter map[string]any) error {
	/* Applies json filters to query as one Has statment at the end of the traversal */
	chainedFilter, err := rt.makeJsonFilter(filter)
	if err != nil {
		return err
	}

	*query = (*query).Has(chainedFilter)
	log.Infof("Has Filter: %v\n", chainedFilter)

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

func (rt *renderTree) makeJsonFilter(filter map[string]any) (*gripql.HasExpression, error) {
	/* Constructs the Grip Has expression used to do json filters */
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
			subExpr, err := rt.makeJsonFilter(itemObj)
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
		if !strings.Contains(field, ".") {
			return nil, fmt.Errorf("filter key '%s' must be of format TYPE.property", field)
		}
		formattedField, err := rt.formatField(field)
		if err != nil {
			return nil, err
		}
		hasExpr, err := mapGraphQLOperatorToGrip(formattedField, topFilter[field], topLevelOp)
		if err != nil {
			return nil, err
		}

		return hasExpr, nil
	}
}

func (rt *renderTree) formatField(field string) (string, error) {
	/* Swaps input field like Observation into traversal state '$fn'*/
	filestrs := strings.Split(field, ".")
	f, ok := rt.fLookup[filestrs[0]]
	if !ok {
		return "", fmt.Errorf("node type %s not found in traversal types %s", filestrs[0], rt.fLookup)
	}
	filestrs[0] = "$" + f
	return strings.Join(filestrs, "."), nil
}

func mapGraphQLOperatorToGrip(field string, value any, op string) (*gripql.HasExpression, error) {
	/* Translates query operator into Grip Has Expression filter operation */
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
		return nil, fmt.Errorf("Operator %s does not match any of the known operators", op)
	}
}
