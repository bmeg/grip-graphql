# GRIP FHIR RESTFUL API

## Local Setup

```
go install github.com/bmeg/grip
go build --buildmode=plugin ./gen3_writer
grip server -w writer=gen3_writer.so
```

Note: the -w writer=gen3_writer.so option can be stacked to include multiple endpoints and long as you have
built the .so file to the corresponding endplint as shown above. Ex:

`grip server -c mongo.yml -w writer=gen3_writer.so -w reader=grip-graphql-endpoint.so -l reader:config=./config/gen3.js -l reader:graph=CALIPER`

## Example queries:

Note: ENV var ACCESS_TOKEN is a valid Gen3 jwt token. An access token is needed for all queries except GET \_status and GET list-graphs

### Delete an edge and then grep for it to see if it has been deleted or not. Format:

_http://localhost:8201/writer/[graph-name]/del-edge/[edge-id]/[gen3-project-id]_

```
curl -X DELETE  http://localhost:8201/writer/CALIPER/del-edge/fb60e763-e799-4d59-82a3-66977cc6696c/ohsu-test
-H "Content-Type: applicationjson" \
-H "Authorization: bearer $ACCESS_TOKEN"

grip query CALIPER "E()" | grep fb60e763-e799-4d59-82a3-66977cc6696c
```

### Bulk load some edges from a file:

_http://localhost:8201/writer/[graph-name]/bulk-load/[gen3-project-id]_

```
curl -X POST "http://localhost:8201/writer/CALIPER/bulk-load/ohsu-test" \
    -H "Authorization: bearer $ACCESS_TOKEN" \
    -F "types=edge" \
    -F "file=@edge.ndjson"
```

Newline delimited edges should be of form:

```
{
    "gid": "bee5bd86-4f06-5eb2-b71a-f62110cf5aa9",
    "label": "specimen_observation",
    "from": "9bc10566-5d7e-4a53-bbc0-6fe9700584a5",
    "to": "ea4ea5e7-2780-46cf-8cc4-fbb40ad63928"
}
```

With required keys "label", "from", "to", and "gid" and optional key "data" with value of type dict.

### Get the data from a vertex given a known vertex id

_http://localhost:8201/writer/[graph-name]/get-vertex/[vertex-id]/[gen3-project-id]_

```
curl -X GET http://localhost:8201/writer/CALIPER/get-vertex/875ddaf8-42da-5d72-b5c5-39c2b16151cd/ohsu-test \
-H "Authorization: bearer $ACCESS_TOKEN"

```

### Get the list of graphs present

```
curl http://localhost:8201/writer/list-graphs
```

### Revproxy Setup --

The above curl commands assume that you are acessing this grip plugin from within the cluster. equivalent queries can be used from outside the cluster by changing the nginx paths to the form:

`https://[your_instance_endpoint]/grip/writer/writer/list-graphs` for the writer or
`https://[your_instance_endpoint]/grip/reader` for the reader api

These paths assume you have checked out to the grip branch of helm and reployed
