run:
	go run ./cmd/main.go
build:
	go build -o bin/party-calc ./cmd/main.go

swag:
	swag init -g ./internal/server/server.go