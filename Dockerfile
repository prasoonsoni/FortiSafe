# Start from the latest golang base image
FROM golang:alpine as builder
LABEL maintainer="Prasoon Soni <prasoonsoni.work@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go app inside the container
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use the lightweight Alpine Linux as the base image
FROM alpine:latest

# Install CA certificates to support secure connections
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory to /root/ inside the container
WORKDIR /root/

# Copy the .env file from the builder stage to the current working directory in this stage
COPY --from=builder /app/.env .env

# Copy the Pre-built binary file (main) from the builder stage to the current working directory in this stage
COPY --from=builder /app/main .

EXPOSE 3000
CMD ["./main"]