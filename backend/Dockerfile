# Build stage
FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

# Add build arguments for platform targeting
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with the target architecture
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o yamlres ./cmd/backend

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Create data directory for resume storage
RUN mkdir -p /app/data

# Copy the binary from builder
COPY --from=builder /app/yamlres .

# Copy default resume configuration
COPY config/sample-resume.yaml /app/config/sample-resume.yaml

# Set user to non-root
RUN adduser -D appuser
USER appuser

# Expose the API port
EXPOSE 8080

# Command to run the application
CMD ["./yamlres"]