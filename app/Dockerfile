FROM golang:1.17.2-alpine3.14

RUN apk update && apk add git

RUN mkdir -p /go/src/app/src

WORKDIR /go/src/app/src

COPY ./src/ /go/src/app/src

CMD ["go", "run", "main.go"]
