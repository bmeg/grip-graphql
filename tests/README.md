# Tests:

Tests can be run locally by specifying that you want to turn on the plugin in testing mode using the `TEST` Graph. For example:

```
grip server  -w writer=gen3_writer.so \
             -w reader=grip-graphql-endpoint.so \
             -w graphql=gql-gen.so \
             -l reader:config=./config/gen3.js \
             -l reader:graph=TEST \
             -l writer:test=true \
             -l graphql:test=true \
             -l reader:test=true
```

then cd to tests directory and run:

`go test` or `go test -v` for logs or `go test -run [specific_test_name]` to run only a specific test.

If the graph name is `TEST` and the config is setup for test=true, mock auth will be used, and these tests can be run locally outside of a gen3 instance.

Note: some tests that rely on the underlying data in the DB only work when all tests are run, since not all tests populate the grip with data

## Running locally

To run the gqlgen plugin locally with no auth on graph CALIPER:

```
grip server -w graphql=gql-gen.so -l graphql:auth=false -l graphql:graph=CALIPER
```

## Version

go version go1.23.6
