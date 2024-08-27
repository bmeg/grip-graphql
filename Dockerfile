FROM golang:1.22.6-alpine AS build-env
RUN apk add make git bash build-base libc-dev binutils-gold curl
ENV GOPATH=/go
ENV PATH="/go/bin:${PATH}"

WORKDIR /go/src/github.com/bmeg/grip-graphql
ADD ./ /go/src/github.com/bmeg/grip-graphql

RUN go install github.com/bmeg/grip@v0.0.0-20240812223417-99c314b06713
RUN go build  --buildmode=plugin ./graphql_gen3
RUN go build  --buildmode=plugin ./gen3_writer
RUN go build  --buildmode=plugin ./grip-graphql-endpoint

RUN cp /go/src/github.com/bmeg/grip-graphql/mongo.yml /
RUN cp /go/src/github.com/bmeg/grip-graphql/gen3_writer.so /
RUN cp /go/src/github.com/bmeg/grip-graphql/graphql_gen3.so /
RUN cp /go/src/github.com/bmeg/grip-graphql/grip-graphql-endpoint.so /



FROM alpine
WORKDIR /data
VOLUME /data
ENV PATH="/app:${PATH}"
COPY --from=build-env /graphql_gen3.so /data/
COPY --from=build-env /gen3_writer.so /data/
COPY --from=build-env /grip-graphql-endpoint.so /data/
COPY --from=build-env /mongo.yml /data/
COPY --from=build-env /go/bin/grip /app/
