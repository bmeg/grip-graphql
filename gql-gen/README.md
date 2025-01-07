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
grip server -w graphql=gql-gen.so  -l graphql:auth=false
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

Example query using random SNOMED codings:

```
{
	"filter": {
		"and": [
			{
				"or": [
					{
						"=": {
							"processing.method.coding.display": "Brief intervention"
						}
					},
					{
						"=": {
							"processing.method.coding.display": "Cuboid syndrome"
						}
					}
				]
			},
			{
				">": {
					"collection.bodySite.concept.coding.code": "261665006"
				}
			}
		]
	}
}
```
