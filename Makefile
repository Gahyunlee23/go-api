# define variables
APP_NAME=main-admin-api
BUILD_DIR=build
MAIN_FILE=main.go

# go commands
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

# env
CONFIG_FILE=config/config.yaml

.PHONY: all build clean run test help deps lint swagger

# basic target
all: clean build

# install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# generate swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	swag init

# build
build: swagger
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

# execute
run:
	@echo "Running $(APP_NAME)..."
	@if [ ! -f $(CONFIG_FILE) ]; then \
		echo "Error: $(CONFIG_FILE) not found!"; \
		exit 1; \
	fi
	go run $(MAIN_FILE)

# execute with dev mode
dev:
	@echo "Running in development mode..."
	air

# test
test:
	@echo "Running tests..."
	go test -v ./...

# text coverage
cover:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# 린트 검사
lint:
	@echo "Running linter..."
	golangci-lint run

# clean
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out
	@go clean

# docker build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME) .

# docker execute
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(APP_NAME)

# help
help:
	@echo "Available commands:"
	@echo "  make deps          - Install project dependencies"
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application"
	@echo "  make dev           - Run the application in development mode with hot-reload"
	@echo "  make test          - Run tests"
	@echo "  make cover         - Run tests with coverage report"
	@echo "  make lint          - Run linter"
	@echo "  make clean         - Clean build files"
	@echo "  make swagger       - Generate Swagger documentation"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"