# STEP 1: build executable binary

# Use the official Golang image as the build stage
FROM golang:alpine3.18 AS build

# Set a label for the maintainer
LABEL maintainer="Titanio Yudista <titanioyudista98@gmail.com>"

# Set the working directory in the build stage
WORKDIR /app

# Copy the source code and any additional files
COPY . .

# Build the Go application
RUN go build -o goapps

# Create a vendor directory and copy dependencies
# RUN go mod vendor

# STEP 2: create a smaller image for the final application

# Use a minimal Alpine Linux image as the final stage
FROM alpine:latest

# # Install necessary packages for the final image
# RUN apk add --no-cache bash

# Copy the compiled application from the build stage to the final image
COPY --from=build /app/goapps .

# Expose port if your application listens on a specific port
EXPOSE 3000

# Define the command to run your application
CMD ["./goapps"]