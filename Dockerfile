# ===== Stage 1: Building GO Binary =====
FROM golang:1.22-alpine AS builder

WORKDIR /app

# copying Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy source code
COPY . .

# Build the binary
RUN go build -o crypto-cli main.go

# ===== Stage 2: Slim runtime container =====
FROM alpine:latest

# Optional: Installing CA Certificates
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/crypto-cli .

ENTRYPOINT ["./crypto-cli"]
