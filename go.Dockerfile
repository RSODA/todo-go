FROM golang:1.24-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .
COPY .env .env

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o bin/todo cmd/todo/main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bin/todo .
COPY --from=builder /app/.env .env

CMD ["./todo"]