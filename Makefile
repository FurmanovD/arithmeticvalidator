APP_NAME := validator-app
TEST_DIR := ./validator
BIN_DIR := ./bin

.PHONY: all build test clean run lint

all: build

build:
	@echo "Building $(APP_NAME)..."
	go build -o $(BIN_DIR)/$(APP_NAME) main.go

test:
	@echo "Running tests..."
	go test $(TEST_DIR) -v -race -count=1

benchmark:
	@echo "Running benchmarks..."
	go test $(TEST_DIR) -bench=. -benchmem

lint:
	@echo "Linting..."
	golangci-lint run ./...

run: build
	@echo "Running $(APP_NAME)..."
	./$(BIN_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BIN_DIR)

