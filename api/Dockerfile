
# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your application will run on
EXPOSE 8080

# Run the Go application
CMD ["./main"]
