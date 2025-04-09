### Grip Graphql Plugins

grip-graphql is a collection of go plugins designed and implemented for connecting a [Grip](https://github.com/bmeg/grip) server to other microservices in a modified [Gen3](https://gen3.org/) software stack.

gen3_writer directory contains a [Gin](https://github.com/gin-gonic/gin) go server plugin that is used for Writing / Deleting data from Graphs on a grip server

gripgraphql directory contains a graphql based read query plugin that uses a [goja](https://github.com/dop251/goja) engine to read from a static schema defined as a config file to create custom graphql queries that can be used to abstract Grip's complex query language into a more digestible query format for frontend usage.

gql-gen directory contains a static graphql schema based on the [FHIR](https://build.fhir.org/downloads.html) data model. It translates fhir queries into the grip query language, then unmarshalls json results back into a predictable FHIR like format.

### Avoiding package version mismatches

Error messages when loading plugins into grip like the one below are quite common:

```
message              Error loading pluging test: plugin.Open("gql-gen"): plugin was built with a different version of package golang.org/x/sys/unix
time                 2024-12-11T16:27:49-08:00
```

These are caused when the grip executable's build package versions are not equivalent with the plugin's build package versions. There are a veriety of reasons why this might be true:

1. Make sure that the go version of Grip go.mod and the go version of this go.mod are the same.

2. Make sure that the versions of all of the packages across both go.mod files are the same

3. install grip with go install github.com/bmeg/grip@[hash] when deploying using a docker container and debug with a replace statement to avoid having to make a commit everytime a change is made to the grip go.mod file.

Ex:

```
replace github.com/bmeg/grip v0.0.0-20241211235035-b772edec00b9 => ../grip
```

Note: Go versions of packages can change when a new package is added in this repo, or when a new package is added in grip.
