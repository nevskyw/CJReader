FROM golang:1.18 AS builder

COPY . /github.com/nevskyw/CJReader
WORKDIR /github.com/nevskyw/CJReader

RUN go mod download
RUN go build -o /bin/CJReader

CMD ["./CJReader"]