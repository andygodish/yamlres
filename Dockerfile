# Stage 1: Build the Vue.js frontend
FROM node:22-alpine AS ui-builder
WORKDIR /app
COPY ui/package*.json ./
RUN npm install
COPY ui/ ./
RUN npm run build

# Stage 2: Build the Go backend
FROM golang:1.24-alpine AS backend-builder
WORKDIR /app
COPY backend/ ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o yamlres ./cmd/backend

# Stage 3: Create the final image with Nginx
FROM nginx:alpine
WORKDIR /app

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates

# Create the directory structure expected by the Go app
RUN mkdir -p /app/config

# Copy the built backend binary
COPY --from=backend-builder /app/yamlres /app/
# Copy configuration files
COPY backend/config/sample-resume.yaml /app/config/sample-resume.yaml
# Copy the Vue.js built files to Nginx's html directory
COPY --from=ui-builder /app/dist /usr/share/nginx/html

# Copy a custom Nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Set environment variables (Backend port)
ENV PORT=8081

# Expose ports for Nginx and the backend
EXPOSE 8080

# Start both Nginx and the backend
CMD ["/bin/sh", "-c", "/app/yamlres & nginx -g 'daemon off;'"]