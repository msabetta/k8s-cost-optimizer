# Multi-stage Dockerfile for k8s-cost-optimizer: builds the Go binary and creates a slim runtime image
# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS builder

# install tools
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copia go.mod e go.sum per caching
COPY go.mod go.sum ./
RUN go mod download

# Copia tutto il codice
COPY . .

# Compila binario
RUN go build -o k8s-cost-optimizer cmd/api/main.go

# 🔹 Runtime image leggera
FROM alpine:3.18
WORKDIR /app

# Certificati e binario
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/k8s-cost-optimizer .

# Copia configs
COPY configs ./configs

# Porta default
EXPOSE 8080

ENTRYPOINT ["./k8s-cost-optimizer"]
