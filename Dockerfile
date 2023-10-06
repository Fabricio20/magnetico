FROM golang:1.21-alpine AS build
WORKDIR /magnetico

RUN apk add --no-cache build-base pkgconf git curl libsodium libzmq czmq

ADD ./pkg             /magnetico/pkg
ADD ./go.mod          /magnetico/go.mod
ADD ./src             /magnetico/src

RUN go mod tidy && cd src && go build -o magneticod

FROM alpine:latest
LABEL maintainer="git@notfab.net"
WORKDIR /

RUN apk add --no-cache libsodium libzmq czmq
COPY --from=build /magnetico/src/magneticod /magneticod

ENTRYPOINT ["/magneticod"]
