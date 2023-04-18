FROM golang:1.19.1-alpine3.16

ENV ROOT=/app
WORKDIR ${ROOT}

RUN apk update && apk add git

COPY go.mod go.sum ${ROOT}

CMD ["go", "run" ,"main.go"]
