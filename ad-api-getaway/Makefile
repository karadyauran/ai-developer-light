start: ## Run the project
	go run cmd/main/main.go

test: ## Run all tests
	go test ./...

install: ## Install dependencies
	go mod tidy

swag: ##  Generate swagger specification
	swag init -g cmd/main/main.go

proto-generate-oauth-service:
	protoc --proto_path=../proto --go_out=./internal --go-grpc_out=./internal ../proto/oauth.proto

proto-generate-kafka-service:
	protoc --proto_path=../proto --go_out=./internal --go-grpc_out=./internal ../proto/kafka.proto

help: ## Display help message
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' Makefile