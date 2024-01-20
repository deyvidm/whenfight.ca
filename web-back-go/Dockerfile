# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go server
RUN go build -o server .

# Expose port 8080 for incoming HTTP requests
EXPOSE 8080

# Start the Go server when the container starts
CMD ["./server"]
