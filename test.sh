#!/bin/bash

set -e

echo "Running yt-stats tests..."
echo

# Run unit tests
echo "=== Running Unit Tests ==="
go test -v -short ./...
echo

# Check if API key is set for integration tests
if [ -n "$YTSTATS_API_KEY" ]; then
    echo "=== Running Integration Tests ==="
    go test -v -tags=integration ./...
    echo
else
    echo "=== Skipping Integration Tests (YTSTATS_API_KEY not set) ==="
    echo
fi

# Run tests with coverage
echo "=== Generating Coverage Report ==="
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
echo

echo "✅ All tests passed!"
