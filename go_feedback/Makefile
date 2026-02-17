# consumer/Makefile

# Variables
APP_NAME = basic-go
DOCKER_IMAGE = ms-consumer:latest
CONTAINER_NAME = ms-consumer-container

# Default target to build the app
all: build

# Build the Go binary
build: fmt
	@echo "Building the Go binary..."
	go build -o ./cmd/consumer/$(APP_NAME) ./cmd/consumer

# Format the Go code
fmt:
	@echo "Formatting the Go code..."
	gofmt -s -w .

# Generate SQLC code
sqlc-generate:
	@echo "Generating SQL code with sqlc..."
	sqlc generate

# Run the application
run: fmt
	@echo "Running the application..."
	go run ./cmd/consumer/main.go

# Dockerize the application
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) -f ./Dockerfile .

# Run the Docker container
docker-run: docker-build
	@echo "Running Docker container..."
	docker run --name $(CONTAINER_NAME) -p 8001:8001 -d $(DOCKER_IMAGE)

# Clean up binaries
clean:
	@echo "Cleaning up..."
	rm -f ./cmd/consumer/$(APP_NAME)

# Stop and remove the Docker container
docker-stop:
	@echo "Stopping Docker container..."
	docker stop $(CONTAINER_NAME) || true
	docker rm $(CONTAINER_NAME) || true

# Setup environment
setup-env:
	@./scripts/setup_env.sh

# Run tests
test:
	@./scripts/test.sh

# Help
help:
	@echo "Makefile Usage:"
	@echo "  make            - Build the Go binary"
	@echo "  make fmt        - Format the Go code"
	@echo "  make sqlc-generate - Generate Go code using sqlc"
	@echo "  make run        - Run the application locally"
	@echo "  make docker-build - Build the Docker image"
	@echo "  make docker-run - Build and run the Docker container"
	@echo "  make docker-stop - Stop and remove the Docker container"
	@echo "  make clean      - Clean up binaries"
	@echo "  make setup-env  - Setup development environment"
	@echo "  make test       - Run tests"
