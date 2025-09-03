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
FROM python:3.12-slim

WORKDIR /app

# Copy Go binary
COPY --from=builder /app/server .

# Copy Python app
COPY python-app/ ./python_app/

# Copy data files
COPY --from=builder /app/data ./data

# Install Python dependencies
RUN pip install --no-cache-dir requests

# Expose HTTP port
EXPOSE 8080

ENTRYPOINT ["./server"]
