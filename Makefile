.PHONY: dev build test clean

dev:
	docker-compose -f deploy/docker/docker-compose.yaml up -d
	cd api && go run cmd/api/main.go

build:
	cd api && go build -o bin/api cmd/api/main.go
	cd gateway && go build -o bin/gateway cmd/gateway/main.go

test:
	cd api && go test ./...
	cd gateway && go test ./...

clean:
	docker-compose -f deploy/docker/docker-compose.yaml down -v 