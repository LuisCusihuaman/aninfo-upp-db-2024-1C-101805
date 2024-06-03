# Stage 1: Build the Go application
FROM golang:1.21.7-alpine3.19 AS builder

# Set the Current Working Directory inside the container
WORKDIR /go/src

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Fetch the missing dependencies
RUN go get -d ./...

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Create a small image and copy the built binary from the previous stage
FROM scratch

# Set the Current Working Directory inside the container
WORKDIR /root/
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/main .

# Command to run the executable
CMD ["./main"]
