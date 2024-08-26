# Stage 1: Build the application
FROM golang:1.22.4-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o app .

# Stage 2: Production image
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose the application port (change if needed)
EXPOSE 80

# Command to run the application
CMD ["./app"]