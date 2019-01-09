FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /tmp/build
ADD . /tmp/build

RUN go get -v ./... && \
    go test -v ./... && \
    go build ./cmd/