# Step 1: Build the Go application
FROM golang:1.23.2 AS builder

WORKDIR /app

# Copy the go.mod and go.sum files to install dependencies
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the entire project (including cmd, internal, and config directories) into the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o app ./cmd/server

# Step 2: Create a lightweight final image
FROM alpine:latest  

WORKDIR /root/

# Copy the built app from the builder stage
COPY --from=builder /app/app .

EXPOSE 8081

# Command to run the application
CMD ["./app"]
