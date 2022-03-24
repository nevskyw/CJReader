FROM golang:latest

WORKDIR /CJReader

COPY ./ /CJReader

RUN go mod download

ENTRYPOINT go run commands/main.go