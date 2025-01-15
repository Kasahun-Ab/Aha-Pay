# Go binary name
BIN_NAME=go_ecommerce

# Go command
GO=go

# Build the project
build:
	$(GO) build -o $(BIN_NAME) ./cmd/api

# Run the project
run:
	$(GO) run ./cmd/api

# Test the project
test:
	$(GO) test ./...

# Format Go code
fmt:
	$(GO) fmt ./...

# Lint the code (optional)
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -f $(BIN_NAME)

# Generate a Docker image (optional)
# docker:
# 	docker build -t $(BIN_NAME) .

# Git branch management
# git branch -M main
