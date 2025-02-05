# Stage 1: Build the Go application
FROM golang:alpine as builder
LABEL authors="shivam"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o url-shortener ./api-services/main.go

# Stage 2: Final lightweight image
FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/url-shortener .

COPY --from=builder /app/frontend ./frontend

EXPOSE 8080

# Start the app
CMD ["./url-shortener"]
