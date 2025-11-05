.PHONY: help build test test-unit test-integration clean run docker-build docker-up docker-down

# Default target
help:
	@echo "Available commands:"
	@echo "  make build              - Build the application binary"
	@echo "  make test               - Run all tests"
	@echo "  make test-unit          - Run only unit tests"
	@echo "  make test-integration   - Run integration tests (requires YTSTATS_API_KEY)"
	@echo "  make clean              - Remove build artifacts"
	@echo "  make run                - Run the application"
	@echo "  make docker-build       - Build Docker image"
	@echo "  make docker-up          - Start services with docker-compose"
	@echo "  make docker-down        - Stop docker-compose services"

# Build the application
build:
	@echo "Building yt-stats..."
	@go build -v -o yt-stats ./cmd/yt-stats

# Run all tests
test:
	@echo "Running all tests..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out

# Run unit tests only
test-unit:
	@echo "Running unit tests..."
	@go test -v -short ./...

# Run integration tests
test-integration:
	@echo "Running integration tests..."
	@if [ -z "$$YTSTATS_API_KEY" ]; then \
		echo "Warning: YTSTATS_API_KEY not set. Integration tests will be skipped."; \
	fi
	@go test -v -tags=integration ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f yt-stats
	@rm -f coverage.out
	@go clean

# Run the application
run: build
	@echo "Starting yt-stats..."
	@./yt-stats

# Docker commands
docker-build:
	@echo "Building Docker image..."
	@docker build -t yt-stats:latest .

docker-up:
	@echo "Starting services..."
	@docker-compose up -d

docker-down:
	@echo "Stopping services..."
	@docker-compose down
