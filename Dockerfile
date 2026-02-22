# STAGE 1: Build
FROM golang:1.21-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-bin ./cmd/api/main.go
RUN go build -o worker-bin ./cmd/worker/main.go

# STAGE 2: Run
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

COPY --from=builder /app/api-bin .
COPY --from=builder /app/worker-bin .

# Secara default menjalankan API, tapi bisa di-override di Docker Compose
CMD ["./api-bin"]