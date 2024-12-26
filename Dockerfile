# Start with the official Golang image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download
RUN go mod vendor

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080 6060

# Start the application
CMD ["./main"]
