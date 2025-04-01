FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy only go.mod and go.sum first to cache dependencies
COPY backend/go.mod backend/go.sum ./
COPY backend/ ./
COPY .env ./

RUN go mod download
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Set the working directory for the final image
WORKDIR /root/

# Copy the pre-built binary and .env file from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose the necessary ports
EXPOSE 9080
EXPOSE 3306

# Command to run the executable
CMD ["./main"]