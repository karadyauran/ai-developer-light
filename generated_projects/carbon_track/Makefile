build:
	go build -o bin/carbon_track cmd/main/main.go

run:
	go run cmd/main/main.go

test:
	go test ./...

docker-build:
	docker build -t carbon_track .

docker-run:
	docker-compose up

docker-stop:
	docker-compose down

lint:
	golangci-lint run

migrate:
	go run cmd/migrate/migrate.go

.PHONY: build run test docker-build docker-run docker-stop lint migrate