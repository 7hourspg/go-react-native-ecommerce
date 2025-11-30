#!/bin/bash

# Script to generate OpenAPI 3.0 documentation
# Note: swaggo/swag generates Swagger 2.0, so we convert to OpenAPI 3.0

set -e

echo "Generating API documentation..."
swag init

echo "Converting to OpenAPI 3.0..."
npx --yes swagger2openapi docs/swagger.json -o docs/openapi.json -y

# Use OpenAPI 3.0 as the default for Swagger UI
cp docs/openapi.json docs/swagger.json

echo "✅ OpenAPI 3.0 documentation ready"

