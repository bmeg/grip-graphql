# GRIP GraphQL Endpoint

Configurable GraphQL endpoint for the GRaph Integration Platform

## Build

Run `make` to build both the plugin (grip-js.so) and proxy server (grip-graphql-proxy).

## Plugin

The plugin is run as a shared object within the GRIP server. There are a few required configs that must be set for this plugin to operate correctly:

Tell grip to load the reader grip-js shared object with
```
-w reader=grip-js.so
```

This option tells grip to use the gen3.js config file for this reader plugin
```
-l reader:config=./grip-js/config/gen3.js
```

You also need to tell the custom endpoints which graph you want to query from. This can be done within
```
-l reader:graph=TEST
```

If you are running locally you will need to turn off auth.
```
-l reader:auth=false
```
Putting it all together:
```
grip server -w reader=grip-js.so  -l reader:config=./grip-js/config/gen3.js -l reader:graph=TEST -l reader:auth=false```
```

## Example Query

If running locally, open http://localhost:8201/reader/ on your browser
set the sandbox url to http://localhost:8201/reader/api and use the frontend to query. See config/gen3.js for examples


## Testing on cluster deployment

Given a config file config/gen3.js with the `IdsWithSpecimenEdge` graphql schema keyword specified:
from the k8s cluster you could do a curl like below if auth=true
```
curl -X POST localhost:8201/reader/api \
    -D '{"query":"query PatientIdsWithSpecimenEdge { PatientIdsWithSpecimenEdge {    id  }}", "variables":{"limit":1000}}' \
    -H "content-type: application/json" \
    -H "Authorization: $ACCESS_TOKEN"
```

The result will be in this case a list of vertices where there exists an out edge from Specimen vertex.

## Proxy

Run the server as a proxy endpoint connected to an external GRIP service

```
./grip-graphql-proxy <grip server> <server port> <config> <database>
```

## Example configuration file

See ./config directory for examples
