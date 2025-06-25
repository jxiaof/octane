# Makefile for Octane Performance Analyzer

# Define the Go binary name
BINARY_NAME=octane

# Define the Go source files
SRC=$(shell find ./cmd -name '*.go') $(shell find ./pkg -name '*.go') main.go

# Define the output directory for binaries
OUTPUT_DIR=bin

# Default target
all: build

# Build the Go binary
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME) $(SRC)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Clean the build
clean:
	@echo "Cleaning up..."
	rm -rf $(OUTPUT_DIR)

# Install the binary
install: build
	@echo "Installing $(BINARY_NAME)..."
	cp $(OUTPUT_DIR)/$(BINARY_NAME) /usr/local/bin/

# Docker build
docker:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME) .

# Help command
help:
	@echo "Makefile commands:"
	@echo "  all        - Build the project"
	@echo "  build      - Build the Go binary"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean the build"
	@echo "  install     - Install the binary"
	@echo "  docker     - Build Docker image"
	@echo "  help       - Show this help message"