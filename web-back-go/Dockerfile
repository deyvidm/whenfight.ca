# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Install Python 3.12 and pip
RUN apt-get update && apt-get install -y python3.12 python3-pip

# Install BeautifulSoup4
RUN pip3 install beautifulsoup4

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
