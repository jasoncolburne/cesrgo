.PHONY: test build clean lint fmt vet help

# Default target
all: test build

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Build the library
build:
	go build ./...

# Clean build artifacts
clean:
	go clean
	rm -f coverage.out coverage.html

# Run linter
lint:
	docker run -t --rm -v "$(shell pwd):/app" -w /app golangci/golangci-lint:v1.64.8 golangci-lint run

# Format code
fmt:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run the example
example:
	go run examples/main.go

# Show help
help:
	@echo "Available targets:"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  build         - Build the library"
	@echo "  clean         - Clean build artifacts"
	@echo "  lint          - Run linter"
	@echo "  fmt           - Format code"
	@echo "  vet           - Run go vet"
	@echo "  deps          - Install dependencies"
	@echo "  example       - Run the example"
	@echo "  help          - Show this help" 