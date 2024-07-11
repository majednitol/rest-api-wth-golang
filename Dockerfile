FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY . .

# Expose port 3000 to the outside world
EXPOSE 3000

# Run the Go application
CMD ["go", "run", "main.go"]
