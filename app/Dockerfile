FROM golang:1.19.1-alpine3.16

ENV ROOT /work
ENV GO_ENV dev

WORKDIR ${ROOT}

RUN apk update && apk add --no-cache alpine-sdk build-base

COPY ./main.go ${ROOT}

CMD ["go","run","main.go"]
