FROM golang:1.23-alpine

# Install git and docker-compose
RUN apk add --no-cache git bash curl docker-cli docker-compose

WORKDIR /app

# Copy your Go webhook code
COPY . .

# Build Go binary
RUN go build -o webhook-listener main.go

# Expose webhook port
EXPOSE 8080

# Run the webhook listener
CMD ["./webhook-listener"]
