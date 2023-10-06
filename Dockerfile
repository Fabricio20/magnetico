FROM golang:1.21-alpine AS build
WORKDIR /magnetico

RUN apk add --no-cache build-base curl git

ADD ./Makefile        /magnetico/
ADD ./pkg             /magnetico/pkg
ADD ./go.mod          /magnetico/go.mod
ADD ./cmd/magneticod  /magnetico/cmd/magneticod

RUN     make magneticod

FROM alpine:latest
LABEL maintainer="git@notfab.net"
WORKDIR /
VOLUME /root/.local/share/magneticod
VOLUME /root/.config/magneticod

COPY --from=build /go/bin/magneticod /magneticod

ENTRYPOINT ["/magneticod"]
