.PHONY: docs help build

help:
	@echo "Available targets:"
	@echo "  docs     - Generate OpenAPI 3.0 documentation (default)"
	@echo "  build    - Build the application"

docs:
	@./scripts/generate-openapi.sh

build:
	@echo "Building application..."
	@go build -o tmp/main .

