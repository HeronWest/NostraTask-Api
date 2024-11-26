.PHONY: default run build test docs clean

APP_NAME=nostratask
MAIN_FILE=cmd/api/main.go
DOCS_DIR=docs
FEATURES_DIR=internal

# Define the default target
default: docs run

# Run the application
run: docs
	@go run $(MAIN_FILE)

# Build the application
build: docs
	@go build -o $(APP_NAME) $(MAIN_FILE)

# Run tests
test:
	@go test ./...

# Generate Swagger docs
docs:
	@swag init -g $(MAIN_FILE) -o $(DOCS_DIR) # Gerar docs no diretório correto

# Clean build and docs files
clean:
	@rm -f $(APP_NAME)
	@rm -rf $(DOCS_DIR)
