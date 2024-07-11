# Use the official golang image based on Alpine as a base image
FROM golang:alpine

# Install necessary tools
RUN apk add --no-cache git

# Install air for hot reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY . .

# Expose port 3000 to the outside 
EXPOSE 3000

# Run the Go application using air for hot reloading
CMD ["air"]
