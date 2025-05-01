# Variables
PROJECT_NAME := go-shopping-store
BINARY_NAME := $(PROJECT_NAME)
GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")

# Default target
all: build

# Build the project
build:
	go build -o bin/$(BINARY_NAME)

# Run the project
run: build
	./bin/$(BINARY_NAME)

# Test the project
test:
	go test ./...

# Format the code
fmt:
	go fmt $(GO_FILES)

# Clean the project
clean:
	go clean
	rm -f bin/$(BINARY_NAME)

# Install dependencies
deps:
	go mod tidy
	go mod download

# Coverage
coverage:
	go clean -testcache
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Lint the code
lint:
	golangci-lint run

# Run all checks
check: fmt lint test

# Build for multiple platforms
build-multiarch:
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)_darwin_amd64
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)_windows_amd64.exe

# Install the binary in $GOPATH/bin
install: build
	go install

# Update dependencies
update-deps:
	go get -u ./...
	go mod tidy
	go mod download