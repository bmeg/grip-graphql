# Tests:

Tests can be run locally by specifying that you want to turn on the plugin in testing mode using the `TEST` Graph. For example:

```
grip server  -w writer=gen3_writer.so \
             -w reader=grip-js.so \
             -w graphql=gql-gen.so \
             -l reader:config=./grip-js/config/gen3.js \
             -l reader:graph=TEST \
             -l writer:test=true \
             -l graphql:test=true \
             -l reader:auth=true -log-level=debug
```

then cd to tests directory and run:

`go test` or `go test -v` for logs or `go test -run [specific_test_name]` to run only a specific test.

If the graph name is `TEST` and the config is setup for test=true, mock auth will be used, and these tests can be run locally outside of a gen3 instance.

The tests leave 2 graphs in the DB, so to clean out the state after the tests are run, run
```
grip drop JSONTEST
grip drop TEST
```

## Running locally

To run the gqlgen plugin locally with no auth on graph CALIPER:

```
grip server -w graphql=gql-gen.so -l graphql:auth=false -l graphql:graph=CALIPER
```

## Version

go version go1.23.6
