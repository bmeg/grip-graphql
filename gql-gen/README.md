# GQL Gen server:

A graphql query server plugin for querying GRIP

## Setup

```
go build --buildmode=plugin ./gql-gen
grip server -w graphql=gql-gen.so
```

Open http://localhost:8201/graphql/

## Generate

To generate new codegen after changing config or schema

```
go run github.com/99designs/gqlgen generate
```

## Local dev no auth

To start server with no auth checks for local development run:

```
grip server -w graphql=gql-gen.so  -l graphql:auth=false -l graphql:graph=CALIPER
```

## Filters

Filters only currently supported on the first node that is queried. Ex: specimen for the example query below

## Example FHIR query:

```
query($filter: JSON){
  specimen(filter: $filter first:100 offset: 5){
    id
    subject{
      ... on PatientType{
        identifier{
          system
          value
        }
      }
    }
   processing{
    method{
      coding{
        code
        display
        system
      }
    }
  }
  }
}


{
  "filter": {
        "=": {
          "processing.method.coding.display":
            "WhateverFieldYouWant"
        }
      }
}
```

## Query Filters

Queries follow the general structure of the FHIR data model except in places where references exist, the graph DB collects all of the vertex data on that specified reference edge and returns it in one dict.

Filters follow the general structure of including "and" and "or" logical operations with comparators as keys

Current supported comparator statements:

| Operator    | Explanation                    |
| ----------- | ------------------------------ |
| "eq", "=",  | equal                          |
| "neq", "!=" | not equal                      |
| "lt", "<"   | less than                      |
| "gt", ">"   | greater than                   |
| "gte", ">=" | greater than or equal          |
| "lte", "<=" | less than or equal             |
| "in"        | value is in the list of values |

Filter keys must be of form TYPE.jsonPath where TYPE is the vertex name without the 'Type' suffix and jsonpath is a '.' delimited path to
the field that will be filtered on. So for example in 'Specimen.processing.method.coding.display' Specimen is short for SpecimenType which is the vertex that is being filtered and 'processing.method.coding.display' is the path to the field that is filtered.

Nested filtering is supported and can be done by specifying nested queried node types instead of the root node type.

Example query using random SNOMED codings:

```
{
	"filter": {
		"and": [
			{
				"or": [
					{
						"=": {
							"Specimen.processing.method.coding.display": "Brief intervention"
						}
					},
					{
						"=": {
							"Specimen.processing.method.coding.display": "Cuboid syndrome"
						}
					}
				]
			},
			{
				">": {
					"Specimen.collection.bodySite.concept.coding.code": "261665006"
				}
			}
		]
	}
}
```
