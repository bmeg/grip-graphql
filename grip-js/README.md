# GRIP GraphQL Endpoint

Configurable GraphQL endpoint for the GRaph Integration Platform

## Build

Run `make` to build both the plugin (grip-graphql-endpoint.so) and proxy server (grip-graphql-proxy).

## Plugin

Run as a shared object within the GRIP server

```
grip server -w reader=grip-graphql-endpoint.so -l reader:config=./config/gen3.js -l graphql:graph=CALIPER
```

## Example Query

Given a config file with the `PatientIdsWithDocumentEdge` graphql schema keyword specified:

```
curl -X POST localhost:8201/reader/api \
    -D '{"query":"query PatientIdsWithDocumentEdge { PatientIdsWithDocumentEdge {    id  }}", "variables":{"limit":1000}}' \
    -H "content-type: application/json" \
    -H "Authorization: $ACCESS_TOKEN"
```

The result will be in this case a list of vertices where there exists an out edge from DocumentReference vertex with label subject to it.

## Proxy

Run the server as a proxy endpoint connected to an external GRIP service

```
./grip-graphql-proxy <grip server> <server port> <config> <database>
```

## Example configuration file

See ./config directory for examples
