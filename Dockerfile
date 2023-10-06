FROM golang:1.21-alpine AS build
WORKDIR /magnetico

RUN apk add --no-cache build-base libsodium libzmq czmq

ADD ./pkg             /magnetico/pkg
ADD ./go.mod          /magnetico/go.mod
ADD ./src             /magnetico

RUN go mod download && go build -o magneticod

FROM alpine:latest
LABEL maintainer="git@notfab.net"
WORKDIR /

RUN apk add --no-cache libsodium libzmq czmq
COPY --from=build /magnetico/magneticod /magneticod

ENTRYPOINT ["/magneticod"]
