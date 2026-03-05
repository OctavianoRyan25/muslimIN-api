# syntax=docker/dockerfile:1.7

################################
# STAGE 1 — Builder
################################
FROM golang:1.24.5-alpine AS builder

# Install dependency minimal
RUN apk add --no-cache git

WORKDIR /app

# Cache go modules lebih efisien
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy source
COPY . .

# Disable CGO supaya binary static & kecil
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Build lebih hemat RAM
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -p=1 -ldflags="-s -w" -o api-bin ./cmd/api

RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -p=1 -ldflags="-s -w" -o worker-bin ./cmd/worker


################################
# STAGE 2 — Runtime (Super kecil)
################################
FROM alpine:3.20

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/api-bin .
COPY --from=builder /app/worker-bin .

# default API
CMD ["./api-bin"]