# Stage 1: Build the Go application
FROM golang:alpine as builder
LABEL authors="shivam"

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod, go.sum, and source code to the container
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go application
RUN go build -o url-shortener ./api-services/main.go

# Stage 2: Create a lightweight runtime image
FROM alpine:latest
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/url-shortener .

# Expose the port the app runs on
EXPOSE 8080

# Start the app
CMD ["./url-shortener"]
