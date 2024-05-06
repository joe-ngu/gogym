.PHONY: build run test

build:
	@echo "Building gogym..."
	@go build -o bin/gogym

run: build
	@echo "Running gogym..."
	@./bin/gogym

test:
	@echo "Running tests..."
	@go test -v ./...

