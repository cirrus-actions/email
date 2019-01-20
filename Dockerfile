FROM golang:alpine as builder

WORKDIR /build/action
ADD . /build/action

RUN apk add --no-cache gcc musl-dev git

RUN CGO_ENABLED=0 GOOS=linux \
    go get -v ./... && \
    go test -v ./... && \
    go build -o build/email ./cmd/

FROM alpine:latest

LABEL version="1.0.0"
LABEL repository="https://github.com/cirrus-actions/actions-trigger/"
LABEL homepage="https://github.com/marketplace/cirrus-ci"
LABEL maintainer="Cirrus Labs"
LABEL "com.github.actions.name"="Email"
LABEL "com.github.actions.description"="Emails check suite results upon completion"
LABEL "com.github.actions.icon"="mail"
LABEL "com.github.actions.color"="green"

COPY --from=builder /build/action/build/email /actions/email
ENTRYPOINT exec /actions/email