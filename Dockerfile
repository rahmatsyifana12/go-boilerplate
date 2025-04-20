# Use the official Golang Alpine image to build the Go app
FROM golang:1.24-alpine AS build

# Install Git and other dependencies
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache for dependencies
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY ./src ./src

# Build the Go app
RUN go build -o main ./src

# Start a new minimal runtime container
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /root

# Install tzdata for timezone support
RUN apk add --no-cache tzdata

# Set timezone to Asia/Jakarta
ENV TZ=Asia/Jakarta
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && \
    echo "Asia/Jakarta" > /etc/timezone

# Copy the pre-built Go binary
COPY --from=build /app/main .

# Copy the .env file
COPY .env .env

# Set environment variables from the .env file
ENV $(cat .env | xargs)

# Expose port from .env using build-time ARG
ARG PORT
EXPOSE ${PORT}

# Command to run the executable
CMD ["./main"]
