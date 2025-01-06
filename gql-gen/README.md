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
  specimen(filter: $filter first:100){
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
