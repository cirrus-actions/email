FROM golang:stretch as builder

WORKDIR /tmp/action
ADD . /tmp/action

RUN CGO_ENABLED=0 \
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

COPY --from=builder /tmp/action/build/email /bin/email

ENTRYPOINT /bin/email