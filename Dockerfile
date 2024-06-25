# Use the official Golang image as a builder
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o slot-machine-api

# Use a minimal base image to package the binary
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/slot-machine-api .

# Copy .env file if needed
COPY .env .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./slot-machine-api"]
