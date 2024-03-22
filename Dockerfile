# Stage 1: Build the application
FROM golang:1.18.2 AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o my-app *.go

# Stage 2: Create a minimal image
#FROM alpine:latest
FROM scratch


WORKDIR /app

# Copy the binary from the previous stage
COPY --from=builder /app/my-app .
COPY --from=builder /app/.env .
COPY --from=builder /app/migrations/* .

## Copy the .env file
#COPY .env .
#
#COPY migrations/ .

# Set environment variables if needed
# ENV PORT=8080

# Expose the desired port
# EXPOSE 8080

# Run the application
CMD ["./my-app"]