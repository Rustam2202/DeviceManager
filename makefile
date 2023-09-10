run:
	go run ./cmd/main.go
build:
	go build -o bin/device-manager ./cmd/main.go

swag:
	swag init -g ./internal/server/server.go

docker-build:
	docker build --tag device-manager .
docker-run:
	docker run --publish 8081:8080 device-manager
compose:
	docker-compose up 