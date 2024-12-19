# Use the official golang image based on Alpine as a base image
FROM golang:alpine


# Install necessary tools
RUN apk add --no-cache git

# Install air for hot reloading
RUN go install github.com/air-verse/air@latest

# Install MongoDB client tools
RUN apk add --no-cache mongodb-tools

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY . .

# Expose port 3000 for the Go application
EXPOSE 3000

# Expose port 27017 for MongoDB (default MongoDB port)
EXPOSE 27017

# CMD instruction to start MongoDB in the background and then run the Go application with air
CMD ["sh", "-c", "mongod --bind_ip_all --smallfiles --port 27017 --dbpath /data/db & air"]




