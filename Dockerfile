# Use the official GoLang base image
FROM golang:1.20.4

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the GoLang Web API
RUN go build -o main .

# Expose the necessary port(s)
EXPOSE 8080

# Set the startup command for the container
CMD ["./main"]