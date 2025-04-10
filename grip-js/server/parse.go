package server

import (
	"fmt"
	"strings"

	"github.com/bmeg/grip/log"
	"github.com/dop251/goja"
	"github.com/graphql-go/graphql"
)

func (e *Endpoint) parseJSHandler(x map[string]any) (QueryField, error) {
	name := ""
	if nameA, ok := x["name"]; ok {
		if nameStr, ok := nameA.(string); ok {
			name = nameStr
		}
	}
	if name == "" {
		return QueryField{}, fmt.Errorf("name not defined")
	}

	var jHandler func(goja.FunctionCall) goja.Value
	if handlerA, ok := x["handler"]; ok {
		if handler, ok := handlerA.(func(goja.FunctionCall) goja.Value); ok {
			jHandler = handler
		} else {
			return QueryField{}, fmt.Errorf("unknown handler type: %#T", handlerA)
		}
	}

	defaults := map[string]any{}
	if defaultsA, ok := x["defaults"]; ok {
		if defaultsM, ok := defaultsA.(map[string]any); ok {
			for k, v := range defaultsM {
				defaults[k] = v
			}
		}
	}

	arguments := map[string]any{}
	if argumentsA, ok := x["args"]; ok {
		if argumentsM, ok := argumentsA.(map[string]any); ok {
			for k, v := range argumentsM {
				arguments[k] = v
			}
		}
	}
	log.Info("ARGS: ", arguments)

	log.Infof("Loading handler %s", name)
	if schemaA, ok := x["schema"]; ok {
		objField, err := parseField(name, schemaA, false) // false for schema (output)
		if err == nil {
			objField.Resolve = func(params graphql.ResolveParams) (any, error) {
				log.Debug("Calling resolver")
				uArgs := map[string]any{}
				for k, v := range defaults {
					uArgs[k] = v
				}
				for k, v := range params.Args {
					uArgs[k] = v
				}

				ctx := params.Context
				vArgs := e.vm.ToValue(uArgs)
				e.vm.Set("ResourceList", ctx.Value("ResourceList"))
				e.vm.Set("Header", ctx.Value("Header"))

				args := goja.FunctionCall{
					Arguments: []goja.Value{e.cw.toValue(), vArgs},
				}
				log.Infof("Calling user function")
				val := jHandler(args)
				log.Infof("Raw handler return: %#v", val.Export())
				out := jsExport(val)
				log.Infof("User function returned: %#v", out)
				return out, nil
			}

			if len(arguments) > 0 {
				args := graphql.FieldConfigArgument{}
				for k, v := range arguments {
					switch v := v.(type) {
					case string:
						if v == "String" {
							args[k] = &graphql.ArgumentConfig{Type: graphql.String}
						} else if v == "Int" {
							args[k] = &graphql.ArgumentConfig{Type: graphql.Int}
						} else if v == "Boolean" {
							args[k] = &graphql.ArgumentConfig{Type: graphql.Boolean}
						} else if v == "Float" {
							args[k] = &graphql.ArgumentConfig{Type: graphql.Float}
						}
					case map[string]any:
						inputObj, err := parseInputObject(k, v)
						if err != nil {
							return QueryField{}, fmt.Errorf("failed to parse input object %s: %s", k, err)
						}
						args[k] = &graphql.ArgumentConfig{Type: inputObj}
					case []any:
						if len(v) != 1 {
							return QueryField{}, fmt.Errorf("incorrect elements in arg array for %s", k)
						}
						if nested, ok := v[0].(map[string]any); ok {
							inputObj, err := parseInputObject(k, nested)
							if err != nil {
								return QueryField{}, fmt.Errorf("failed to parse list input object %s: %s", k, err)
							}
							args[k] = &graphql.ArgumentConfig{Type: graphql.NewList(inputObj)}
						}
					}
				}
				objField.Args = args
			}
			return QueryField{
				name:    name,
				field:   objField,
				handler: jHandler,
			}, nil
		} else {
			return QueryField{}, fmt.Errorf("parse error: %s", err)
		}
	} else {
		return QueryField{}, fmt.Errorf("schema not found for %s", name)
	}
}

func parseField(name string, x any, isInput bool) (*graphql.Field, error) {
	switch v := x.(type) {
	case string:
		if v == "Int" {
			return &graphql.Field{Name: name, Type: graphql.Int}, nil
		}
		if v == "String" {
			return &graphql.Field{Name: name, Type: graphql.String}, nil
		}
		if v == "Boolean" {
			return &graphql.Field{Name: name, Type: graphql.Boolean}, nil
		}
		if v == "Float" {
			return &graphql.Field{Name: name, Type: graphql.Float}, nil
		}
	case []any:
		if len(v) != 1 {
			return nil, fmt.Errorf("incorrect elements in schema array (only 1)")
		}
		if lf, err := parseField(name, v[0], isInput); err == nil {
			l := graphql.NewList(lf.Type)
			return &graphql.Field{Name: name, Type: l}, nil
		} else {
			log.Errorf("Error parsing list: %s", err)
		}
	case map[string]any:
		if isInput {
			obj, err := parseInputObject(name, v)
			if err != nil {
				return nil, err
			}
			return &graphql.Field{Name: name, Type: obj}, nil
		} else {
			obj, err := parseObject(name, v, isInput)
			if err != nil {
				return nil, err
			}
			return &graphql.Field{Name: name, Type: obj}, nil
		}
	}
	return nil, fmt.Errorf("type not found: %#v", x)
}

// parseInputObject creates a unique input object type for nested args
func parseInputObject(parentName string, x map[string]any) (*graphql.InputObject, error) {
	fields := graphql.InputObjectConfigFieldMap{}
	for k, v := range x {
		switch v := v.(type) {
		case string:
			if v == "String" {
				fields[k] = &graphql.InputObjectFieldConfig{Type: graphql.String}
			} else if v == "Int" {
				fields[k] = &graphql.InputObjectFieldConfig{Type: graphql.Int}
			} else if v == "Boolean" {
				fields[k] = &graphql.InputObjectFieldConfig{Type: graphql.Boolean}
			} else if v == "Float" {
				fields[k] = &graphql.InputObjectFieldConfig{Type: graphql.Float}
			}
		case []any:
			if len(v) != 1 {
				return nil, fmt.Errorf("incorrect elements in input array (only 1)")
			}
			if nested, ok := v[0].(map[string]any); ok {
				// Use parentName + k to ensure uniqueness (e.g., "conceptCoding", "methodCoding")
				nestedObj, err := parseInputObject(parentName+k, nested)
				if err != nil {
					return nil, err
				}
				fields[k] = &graphql.InputObjectFieldConfig{Type: graphql.NewList(nestedObj)}
			} else {
				return nil, fmt.Errorf("array element must be an object")
			}
		case map[string]any:
			// Use parentName + k for nested objects (e.g., "collectionBodySite")
			nestedObj, err := parseInputObject(parentName+k, v)
			if err != nil {
				return nil, err
			}
			fields[k] = &graphql.InputObjectFieldConfig{Type: nestedObj}
		default:
			return nil, fmt.Errorf("unsupported input type for %s: %#v", k, v)
		}
	}
	// Capitalize the first letter of the name for GraphQL convention
	typeName := parentName + "Input"
	if len(typeName) > 0 {
		typeName = strings.ToUpper(string(typeName[0])) + typeName[1:]
	}
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name:   typeName,
		Fields: fields,
	}), nil
}

func parseObject(name string, x map[string]any, isInput bool) (*graphql.Object, error) {
	fields := graphql.Fields{}
	for k, v := range x {
		f, err := parseField(k, v, isInput)
		if err == nil {
			fields[k] = f
		} else {
			return nil, err
		}
	}
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   name,
		Fields: fields,
	}), nil
}

func jsExport(val goja.Value) any {
	o := val.Export()
	log.Infof("jsExport input: %#v", o) // Debug
	if oList, ok := o.([]any); ok {
		out := []any{}
		for _, i := range oList {
			if ov, ok := i.(goja.Value); ok {
				out = append(out, jsExport(ov))
			} else {
				out = append(out, i)
			}
		}
		return out
	}
	if oMap, ok := o.(map[string]any); ok {
		out := map[string]any{}
		for k, v := range oMap {
			if ov, ok := v.(goja.Value); ok {
				out[k] = jsExport(ov)
			} else {
				out[k] = v
			}
		}
	}
	return o
}
