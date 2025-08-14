FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o qr-generator

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/qr-generator .
COPY --from=builder /app/templates ./templates

EXPOSE 8080

CMD ["./qr-generator"]