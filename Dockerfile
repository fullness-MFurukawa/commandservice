FROM golang:1.21.3-alpine3.17
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/command
WORKDIR /go/src/command
ADD . /go/src/command