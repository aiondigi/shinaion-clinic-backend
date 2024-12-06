# Use the official Go 1.23 image
FROM golang:1.23-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules and Sum Files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch (to make the image smaller)
FROM alpine:latest

# Install necessary dependencies (like ca-certificates) for making HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/main /app/

# Expose port 8080 for the app
EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
