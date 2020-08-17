FROM golang:1.14

WORKDIR /usr/src
COPY . .

RUN "go build relay.go"

CMD ["relay"]