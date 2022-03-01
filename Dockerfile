FROM golang:1.17.5 as development
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod booking-app
# Copy app files
COPY . .
# Install Reflex for development
RUN go install github.com/cespare/reflex@latest
# Expose port
EXPOSE 4000
# Start app
CMD reflex -g '*.go' go run api.go --start-service
