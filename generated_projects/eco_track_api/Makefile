build:
	go build -o bin/eco_track_api cmd/main/main.go

run: build
	./bin/eco_track_api

test:
	go test ./...

migrate-up:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate -path ./sql/migrations -database $(DATABASE_URL) up

migrate-down:
	go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate -path ./sql/migrations -database $(DATABASE_URL) down

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

lint:
	golangci-lint run

clean:
	rm -rf bin/*

.PHONY: build run test migrate-up migrate-down docker-up docker-down lint clean