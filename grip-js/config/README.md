## Configs

Each config needs a Name, schema and a handler. It is also useful to define defaults and args so that you can customize the amount of data being returned.

### Schema

This is where you define how the graphql query should be formatted.

In the examples below I am defining the format of a FHIR specimen, but the query doesn't need to be this complicated.

At every leaf in the schema define a GraphQl primitive data type. Int, String, Boolean, or Float.

## Args

Currently there is only support for offset and limit. There might be more args added in the future

## handler

Handler is where you define the query you want the corresponding schema you just specified to execute. The schema should match the corresponding grip query expected output.

At the end of the query you have to put a "toList()" call otherwise the query will not be executed.

## Mutation endpoints.

There is also minimal support for mutation endpoints.
inside Args, specify what data you want to populate in grip. These args match the schema of the read command to illustrate a Read/Write operation on a FHIR specimen.

Handler is the place where you can inject custom JS for formatting. In this case in config/gen3.js it is simply pulling together all of the definitions in args.
