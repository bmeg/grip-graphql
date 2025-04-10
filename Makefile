
all: grip-js.so gen3_writer.so gql-gen.so grip-graphql-proxy

gen3_writer.so :  $(shell find gen3_writer -name "*.go")
	go build --buildmode=plugin ./gen3_writer

grip-js.so :  $(shell find grip-js -name "*.go")
	go build --buildmode=plugin ./grip-js

gql-gen.so: $(shell find gql-gen -name "*.go")
	go build --buildmode=plugin ./gql-gen

grip-graphql-proxy : $(shell find cmd/grip-graphql-proxy -name "*.go")
	go build ./cmd/grip-graphql-proxy

clean:
	rm grip-graphql-proxy grip-js.so gen3_writer.so gql-gen.so