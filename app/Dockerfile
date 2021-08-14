FROM golang:1.13.8-alpine

RUN apk update && apk add git

RUN mkdir -p /go/src/app/src

WORKDIR /go/src/app/src

COPY ./src/ /go/src/app/src
