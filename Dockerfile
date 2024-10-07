# Use the official Golang image as a base image
FROM golang:1.20-alpine

# Set environment variables
ENV GO111MODULE=on

# Install SQLite and migrate CLI
RUN apk update && apk add --no-cache gcc musl-dev sqlite bash \
    && go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go Gin application
RUN go build -o main .

# Copy the migration files
COPY ./migrations /app/migrations

# Expose the application port
EXPOSE 8080

# Command to run migrations first, then start the Go app
CMD migrate -path /app/migrations -database sqlite3://test.db up && ./main
