# GQL Gen server:

A graphql query server plugin for querying GRIP

## Setup

```
go build --buildmode=plugin ./gql-gen
grip server -w graphql=gql-gen.so
```

Open http://localhost:8201/graphql/

## Example FHIR query:

```
query($filter: JSON){
  specimen(filter: $filter first:10){
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
          "id":
            "example-uuid"
        }
      }
}
```
