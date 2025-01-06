package graph

import (
	"fmt"
	"strings"

	"github.com/bmeg/grip/gripql"
	"google.golang.org/protobuf/types/known/structpb"
)

func applyUnwinds(query **gripql.Query, rt *renderTree) {
	/* Assumes query is at f0 and only applies unwinds to that node currently*/
	for _, val := range rt.rFieldPaths["f0"].unwindPath {
		*query = (*query).Unwind(val)
		fmt.Println(*query)
	}
}

/*
Note since this function only has access to one row at a time it cannot merge
the rows and do a full rewind it can only add list objects to make it valid with the existing schema.
TODO: create proper merge function in grip
*/
func applyRewinds(query **gripql.Query, rt *renderTree) error {
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

func applyFilters(query **gripql.Query, args map[string]any) error {
	//Todo: support "sort" operations
	//fmt.Printf("FIRST: %v, TYPE: %T\n", args, args["first"])
	if filter, ok := args["filter"]; ok {
		if filter != nil && len(filter.(map[string]any)) > 0 {
			chainedFilter, err := applyJsonFilter(filter.(map[string]any))
			if err != nil {
				fmt.Println("ERR != NIL: ", err)
				return err
			}
			//fmt.Printf("CHAINED FILTER: %s\n", chainedFilter.String())
			*query = (*query).Has(chainedFilter)
		}
	}
	if first, ok := args["first"]; ok {
		firstPtr, _ := first.(*int)
		if firstPtr == nil {
			*query = (*query).Limit(uint32(10))
		} else {
			*query = (*query).Limit(uint32(*firstPtr))
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
