
# GRIP GraphQL Endpoint
Configurable GraphQL endpoint for the GRaph Integration Platform


## Build

Run `make` to build both the plugin (grip-graphql-endpoint.so) and proxy server (grip-graphql-proxy).


## Plugin
Run as a shared object within the GRIP server
```
grip server -w graphql=grip-graphql-endpoint.so -l graphql:config=./config/config.js -l graphql:graph=test-db
```

## Proxy
Run the server as a proxy endpoint connected to an external GRIP service
```
./grip-graphql-proxy <grip server> <server port> <config> <database>
```

## Example configuration file

```javascript

endpoint.add({
    name: "projects",
    schema: [
        "String"
    ],
    handler: (G, args) => {
        return G.V().hasLabel("Project").render("_gid").toList()
    }
})

endpoint.add({
    name: "cases",
    schema: [
        "String"
    ],
    args: {
        offset: "Int",
        limit: "Int",
        project_id : "String"
    }, 
    defaults: {
        offset: 0,
        limit: 100
    },
    handler: (G, args) => {
        if (args.project_id === undefined) {
            return G.V().hasLabel("Case").skip(args.offset).limit(args.limit).render("_gid").toList()
        } else {
            return G.V().hasLabel("Case").has(gripql.eq("project_id", args.project_id)).skip(args.offset).limit(args.limit).render("_gid").toList()
        }
    }
})

endpoint.add({
    name: "caseCounts",
    schema: {
        cases: "Int",
        samples: "Int",
        aliquots: "Int"
    },
    args: {
        project_id: "String"
    },
    handler: (G, args) => {
        return {
            "cases": G.V().hasLabel("Case").has(gripql.eq("project_id", args.project_id)).count().toList()[0],
            "samples": G.V().hasLabel("Case").has(gripql.eq("project_id", args.project_id)).out("samples").count().toList()[0],
            "aliquots": G.V().hasLabel("Case").has(gripql.eq("project_id", args.project_id)).out("samples").out("aliquots").count().toList()[0],
        }
    }
})
```