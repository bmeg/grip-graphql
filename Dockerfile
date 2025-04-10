FROM golang:1.24.0-alpine AS build-env
RUN apk add make git bash build-base libc-dev binutils-gold curl
ENV GOPATH=/go
ENV PATH="/go/bin:${PATH}"

ADD ./ /go/src/github.com/bmeg/grip-graphql
WORKDIR /go/src/github.com/bmeg/grip-graphql

RUN go install github.com/bmeg/grip@d0f5c735c219fae0c9b39c3733374015084036db
RUN make all

FROM alpine
WORKDIR /data
VOLUME /data
ENV PATH="/app:${PATH}"
COPY --from=build-env /go/src/github.com/bmeg/grip-graphql/gql-gen.so /data/
COPY --from=build-env /go/src/github.com/bmeg/grip-graphql/gen3_writer.so /data/
COPY --from=build-env /go/src/github.com/bmeg/grip-graphql/grip-js.so /data/
COPY --from=build-env /go/src/github.com/bmeg/grip-graphql/grip-js/config/gen3.js /data/config/
COPY --from=build-env /go/bin/grip /app/
