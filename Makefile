.PHONY: build run start test clean db-up db-down migrate migrate-down

# Default target executed when no arguments are given to make.
default: help

# Help target
help:
	@echo "Application Portfolio Management (APM) Development Commands"
	@echo ""
	@echo "Usage:"
	@echo "  make build        Build the application"
	@echo "  make run          Run the application locally"
	@echo "  make test         Run tests"
	@echo "  make db-up        Start the database"
	@echo "  make db-down      Stop the database"
	@echo "  make migrate      Run database migrations up"
	@echo "  make migrate-down Revert database migrations"
	@echo "  make clean        Clean build artifacts"
	@echo ""

# Build the application
build:
	@echo "Building the application..."
	go build -o bin/apm main.go

# Run the application
run:
	@echo "Running the application..."
	go run main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Start the database
db-up:
	@echo "Starting the database..."
	docker-compose up -d postgres

# Stop the database
db-down:
	@echo "Stopping the database..."
	docker-compose down

# Run database migrations up
migrate:
	@echo "Running database migrations..."
	@echo "For now, this is handled by the docker-compose startup"
	@echo "In the future, you might want to use a migration tool like golang-migrate"

# Revert database migrations
migrate-down:
	@echo "Reverting database migrations..."
	@echo "For now, migrations are handled by the docker-compose startup"
	@echo "In the future, you might want to use a migration tool like golang-migrate"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean 