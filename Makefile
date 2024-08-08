
all: grip-graphql-endpoint.so grip-graphql-proxy

grip-graphql-endpoint.so :  $(shell find grip-graphql-endpoint -name "*.go")
	go build --buildmode=plugin ./grip-graphql-endpoint

grip-graphql-proxy : $(shell find cmd/grip-graphql-proxy -name "*.go")
	go build ./cmd/grip-graphql-proxy

graphql_gen3 : $(shell find graphql_gen3 -name "*.go")
	go build --buildmode=plugin ./graphql_gen3

clean: 
	rm grip-graphql-proxy grip-graphql-endpoint.so