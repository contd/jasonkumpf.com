# Builder image
FROM golang AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY server/server.go .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o server

# Production image
FROM golang:1.12.9-stretch
COPY --from=builder /app/server /app/
ENV SERVER_PORT ":8088"
ENV SERVER_HOST "0.0.0.0"
VOLUME [ "/photos" ]
EXPOSE 8088
ENTRYPOINT [ "/app/server" ]
