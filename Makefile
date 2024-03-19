
all: grip-graphql-endpoint.so grip-graphql-proxy

grip-graphql-endpoint.so :  $(shell find grip-graphql-endpoint -name "*.go")
	go build --buildmode=plugin ./grip-graphql-endpoint

grip-graphql-proxy: $(shell find cmd/grip-graphql-proxy -name "*.go")
	go build ./cmd/grip-graphql-proxy

