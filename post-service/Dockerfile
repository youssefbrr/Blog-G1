# Use an official Go image as the base
FROM golang:1.23.3

# Enable Go modules
ENV GO111MODULE=on

# Set the working directory in the container
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Copy and customize .env from the parent directory
COPY .env .env

# Build the Go application
RUN go build -o post .

# Expose the port your application will run on
EXPOSE 8080

# Start the application
CMD ["./post"]
