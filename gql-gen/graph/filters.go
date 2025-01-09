package graph

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bmeg/grip/gripql"
	"github.com/bmeg/grip/gripql/inspect"
	"google.golang.org/protobuf/types/known/structpb"
)

func applyAuthFilters(q **gripql.Query, authList []any) {
	Has_Statement := &gripql.GraphStatement{Statement: &gripql.GraphStatement_Has{gripql.Within("auth_resource_path", authList...)}}
	steps := inspect.PipelineSteps((*q).Statements)
	//fmt.Println("STEPS: ", steps)
	FilteredGS := []*gripql.GraphStatement{}
	step_value := 0
	for i, v := range (*q).Statements {
		//fmt.Println("statement v: ", v)
		steps_index, _ := strconv.Atoi(steps[i])
		if i == 0 {
			FilteredGS = append(FilteredGS, v)
			continue
		} else if steps_index > step_value {
			FilteredGS = append(FilteredGS, v, Has_Statement)
			step_value++
		} else {
			FilteredGS = append(FilteredGS, v)
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

func (rt *renderTree) applyUnwinds(query **gripql.Query) {
	/* Assumes query is at f0 and only applies unwinds to that node currently*/
	for _, val := range rt.rFieldPaths["f0"].unwindPath {
		*query = (*query).Unwind(val)
	}
}

/*
Note since this function only has access to one row at a time it cannot merge
the rows and do a full rewind it can only add list objects to make it valid with the existing schema.
TODO: create proper merge function in grip
*/
func (rt *renderTree) applyRewinds(query **gripql.Query) error {
	/*
		Applies a JS function to every row in the grip query
		Args:
			x - input row object
			args - user defined function args used in the function
	*/
	jsfunc := `
function RewindObj(x, args) {
  args.paths.sort((a, b) => b.split('.').length - a.split('.').length);
  args.paths.forEach(path => {
    const keys = path.split('.');
    let current = x;

    for (let i = 0; i < keys.length - 1; i++) {
      if (current[keys[i]] === undefined) {
        return;
      }
      current = current[keys[i]];
    }

    const lastKey = keys[keys.length - 1];
    if (current[lastKey] !== undefined && !Array.isArray(current[lastKey])) {
      current[lastKey] = [current[lastKey]];
    }
  });

  return [x];
}
`

	fields := make(map[string]*structpb.Value)
	values := make([]*structpb.Value, len(rt.rFieldPaths["f0"].unwindPath))
	for i, p := range rt.rFieldPaths["f0"].unwindPath {
		values[i] = structpb.NewStringValue(p)
	}

	fields["paths"] = structpb.NewListValue(&structpb.ListValue{Values: values})
	pbStruct := &structpb.Struct{Fields: fields}
	*query = (*query).FlatMap(&gripql.Code{Function: "RewindObj", Source: jsfunc, Args: pbStruct})
	return nil
}

func (rt *renderTree) applyFilters(query **gripql.Query, filter map[string]any) error {
	//Todo: support "sort" operations

	rt.applyUnwinds(query)
	*query = (*query).As("f0")
	chainedFilter, err := applyJsonFilter(filter)
	if err != nil {
		return err
	}

	*query = (*query).Has(chainedFilter)
	err = rt.applyRewinds(query)
	if err != nil {
		return err
	}
	*query = (*query).As("f0")

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
