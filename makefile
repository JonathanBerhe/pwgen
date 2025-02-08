# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=pwgen
COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html

# Make sure we use proper OS binary extension
ifeq ($(OS),Windows_NT)
	BINARY_NAME := $(BINARY_NAME).exe
endif

# Main build target directory
BUILD_DIR=build

# Get the current commit hash
COMMIT_HASH=$(shell git rev-parse --short HEAD)

# Build flags
LDFLAGS=-ldflags "-X main.version=$(COMMIT_HASH)"

.PHONY: all build test coverage clean lint deps help

# Default target
all: clean build test

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/pwgen

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v -race ./...

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=$(COVERAGE_FILE) ./...
	$(GOCMD) tool cover -func=$(COVERAGE_FILE)
	$(GOCMD) tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	rm -f $(COVERAGE_FILE)
	rm -f $(COVERAGE_HTML)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GOCMD) mod download
	$(GOCMD) mod tidy

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Build for multiple platforms
build-all: clean
	@echo "Building for multiple platforms..."
	# Linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 ./cmd/pwgen
	# Windows
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cmd/pwgen
	# MacOS
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 ./cmd/pwgen

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Show help
help:
	@echo "Available commands:"
	@echo "  make build      - Build the application"
	@echo "  make test       - Run tests"
	@echo "  make coverage   - Run tests with coverage"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make deps       - Install dependencies"
	@echo "  make lint       - Run linter"
	@echo "  make build-all  - Build for multiple platforms"
	@echo "  make run        - Run the application"
	@echo "  make all        - Clean, build, and test"
	@echo "  make help       - Show this help message"
