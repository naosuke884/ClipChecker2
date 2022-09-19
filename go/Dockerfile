FROM golang:1.19.1-alpine3.16

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY ./src/main.go ${ROOT}

COPY ./src/go.mod ${ROOT}

RUN go mod tidy
