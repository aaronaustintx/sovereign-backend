# ---------- Build Stage ----------
FROM golang:1.27-alpine AS builder

# Install build tools
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o sovereign ./cmd/api

# ---------- Run Stage ----------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/sovereign .

EXPOSE 8080

CMD ["./sovereign"]
