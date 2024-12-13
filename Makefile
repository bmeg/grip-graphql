
all: grip-graphql-endpoint.so gen3_writer.so grip-graphql-proxy gql-gen

gen3_writer.so :  $(shell find gen3_writer -name "*.go")
	go build --buildmode=plugin ./gen3_writer

grip-graphql-endpoint.so :  $(shell find grip-graphql-endpoint -name "*.go")
	go build --buildmode=plugin ./grip-graphql-endpoint

grip-graphql-proxy : $(shell find cmd/grip-graphql-proxy -name "*.go")
	go build ./cmd/grip-graphql-proxy

graphql_gen3 : $(shell find graphql_gen3 -name "*.go")
	go build --buildmode=plugin ./graphql_gen3

gql-gen : $(shell find gql-gen -name "*.go")
	go build --buildmode=plugin ./gql-gen

clean:
	rm grip-graphql-proxy grip-graphql-endpoint.so gen3_writer.so gql-gen.so