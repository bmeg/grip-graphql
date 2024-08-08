# GRIP FHIR RESTFUL API

## Setup

Create a grip executable or build one locally by cloning grip main branch and 
do a:

```
export PATH=$PATH:$HOME/go/bin
go build .
go install .
```

```
go build --buildmode=plugin ./gen3_writer
grip server -w api/writer=gen3_writer.so 
```

Note: the -w api/writer=gen3_writer.so option can be stacked to include multiple endpoints and long as you have
built the .so file to the corresponding endplint as shown above.

## Example queries: 

Delete an edge and then grep for it to see if it has been deleted or not
```
curl -X DELETE -H "Content-Type: applicationjson" http://localhost:8201/api/writer/test/del-edge/2XM1hqErQvIw9s0cUWDuRpKPbQR
grip query test "E()" | grep 2XM1hqErQvIw9s0cUWDuRpKPbQR
```

Bulk load some edges from an edge file:
```
curl -X POST -H "Content-Type: applicationjson" -d '{"edge": "../../aced-data/grip-aced-data/edge.ndjson"}' http://localhost:8201/api/writer/test/bulk-load 
```

Note: Each edge in edge file above will be of format:

```
{
    "label": "custodian",
    "from": "9bc10566-5d7e-4a53-bbc0-6fe9700584a5",
    "to": "Organization/ea4ea5e7-2780-46cf-8cc4-fbb40ad63928"
}
```
With required keys "label", "from" and "to"

Get the value of the vertex with id 302324d5-1d92-5425-80d5-ac6c63af84b6
```
curl -X GET http://localhost:8201/api/graphql/test/get-vertex/302324d5-1d92-5425-80d5-ac6c63af84b6
```

Get the list of graphs present
```
curl -X GET http://localhost:8201/api/graphql/list-graphs
```

## Version
go version go1.21.3

## Tests: not currently functional. 
