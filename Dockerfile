FROM golang:stretch

WORKDIR /tmp/build
ADD . /tmp/build

RUN CGO_ENABLED=0 \
    go get -v ./... && \
    go test -v ./... && \
    go build -o build/email ./cmd/