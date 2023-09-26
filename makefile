run:
	go run ./cmd/main.go -confpath=./cmd/
build:
	go build -o bin/device-manager ./cmd/main.go

swag:
	swag fmt
	swag init -g ./internal/server/server.go
	npx @redocly/cli build-docs ./docs/swagger.json -o ./docs/index.html

docker-build:
	docker build --tag device-manager .
docker-run:
	docker run -p 8081:8080 device-manager

compose:
	docker-compose up -d --build
compose-down:
	docker-compose down
	
docker-debug-build:
	docker build --file Dockerfile.debug --tag device-manager-debugger .
docker-debug-run:
	docker-compose -f ./debug/docker-compose.yml up -d --build

lint:
	golangci-lint run

test:
	go test ./... -cover -coverprofile=coverage.out
test-cover-report:
	make test
	go tool cover -html=coverage.out
