FROM golang:1.22.5-alpine AS build-env
RUN apk add make git bash build-base libc-dev binutils-gold curl
ENV GOPATH=/go
ENV PATH="/go/bin:${PATH}"

ADD ./ /go/src/github.com/bmeg/grip-graphql
WORKDIR /go/src/github.com/bmeg/grip-graphql

RUN go install github.com/bmeg/grip@7a9207d1ea5f6a3c9e762729faa9fd2b117157bf
RUN make all

#FROM alpine
#WORKDIR /data
#VOLUME /data
#ENV PATH="/app:${PATH}"
#COPY --from=build-env /go/src/github.com/bmeg/grip-graphql/gql-gen.so /data/
#COPY --from=build-env /go/src/github.com/bmeg/grip-graphql/gen3_writer.so /data/
#COPY --from=build-env /go/src/github.com/bmeg/grip-graphql/grip-graphql-endpoint.so /data/
#COPY --from=build-env /go/src/github.com/bmeg/grip-graphql/config/gen3.js /data/config/
#COPY --from=build-env /go/src/github.com/bmeg/grip-graphql/mongo.yml /data/
#COPY --from=build-env /go/bin/grip /app/
