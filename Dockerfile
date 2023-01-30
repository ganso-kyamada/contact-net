FROM golang:1.19.4-alpine3.16

WORKDIR /app

RUN apk update && apk add bash

CMD [ "go", "mod", "tidy" ]