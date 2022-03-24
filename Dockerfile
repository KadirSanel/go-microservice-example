# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/core

RUN go build -o /app/bin/example-microservice

EXPOSE 8080 8080

CMD [ "/app/bin/example-microservice" ]