FROM golang:1.18

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]

# $ docker build -t my-golang-app .
# $ docker run -it --rm --name my-running-app my-golang-app