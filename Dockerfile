FROM golang:1.14

WORKDIR /usr/src
COPY . .

RUN go build .

ENTRYPOINT ./relay

# Relay service listens on http://localhost:8080
EXPOSE 8080