FROM golang:1.19.1-alpine3.16

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY . ${ROOT}

RUN go mod tidy
