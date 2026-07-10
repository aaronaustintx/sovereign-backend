# ---------- Build Stage ----------
FROM golang:1.22-alpine AS builder

# Install build tools
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o sovereign ./cmd/api

# ---------- Run Stage ----------
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/sovereign .

# Copy config if needed
COPY config.yaml .

# Expose API port
EXPOSE 8080

# Run the server
CMD ["./sovereign"]
