FROM golang:1.19.0-alpine3.16

WORKDIR /go/src

RUN apk update

CMD [ "go", "mod", "tidy" ]