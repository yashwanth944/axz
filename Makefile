.PHONY: build run dev clean test

build:
	go build -o bin/api-server ./cmd/api-server

run: build
	./bin/api-server

dev:
	docker-compose up -d postgres
	DATABASE_HOST=localhost DATABASE_PORT=5433 go run ./cmd/api-server/main.go

clean:
	rm -rf bin/
	docker-compose down -v

test:
	go test ./... 