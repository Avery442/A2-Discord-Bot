# Use Go 1.25-alpine to satisfy go.mod
FROM golang:1.25-alpine

WORKDIR /app

# Install git and ca-certificates if needed
RUN apk add --no-cache git ca-certificates

# Copy modules and download dependencies (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build statically linked binary
RUN go build -o a2bot main.go
RUN chmod +x a2bot

# Copy assets
COPY assets ./assets

# Set working directory
WORKDIR /app

# No ports exposed, bot connects outbound
# Environment variables will be injected by Pterodactyl
ENV TOKEN=""
ENV A2_API_KEY=""
ENV HOTMIKE=""

# Start the bot
CMD ["./a2bot"]
