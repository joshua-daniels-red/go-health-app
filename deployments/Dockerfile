# ---- Build stage ----
FROM golang:1.24 AS builder

WORKDIR /app

# Cache dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# ---- Runtime stage ----
FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /

# Copy binary from builder
COPY --from=builder /app/server .

# Expose HTTP port
EXPOSE 8080

# Run as non-root user
USER nonroot:nonroot

ENTRYPOINT ["/server"]
