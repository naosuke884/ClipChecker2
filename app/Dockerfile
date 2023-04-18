FROM golang:1.19.1-alpine3.16

ENV ROOT=/work
WORKDIR ${ROOT}

RUN apk update && apk add git
RUN apk add gcc

COPY ./ ${ROOT}

RUN go mod download

# CMD ["go", "run" ,"main.go"]
CMD ["ash"]
