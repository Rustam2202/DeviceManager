run:
	go run ./cmd/main.go
build:
	go build -o bin/device-manager ./cmd/main.go

swag:
	swag fmt
	swag init -g ./internal/server/server.go

docker-build:
	docker build --tag device-manager .
docker-run:
	docker run -p 8081:8080 device-manager
compose:
	docker-compose up -d
	
docker-debug-build:
	docker build --file Dockerfile.debug --tag device-manager-debugger .
docker-debug-run:
	docker run -d -p 8082:8080 -p 4000:4000 --name  device-manager-debug device-manager-debugger

lint:
	golangci-lint run

test:
	go test ./... -cover -coverprofile=coverage.out
test-cover-report:
	make test
	go tool cover -html=coverage.out

# zookeeper-run:
# 	bin/zookeeper-server-start.sh config/zookeeper.properties
# 	bin/windows/zookeeper-server-start.bat config/zookeeper.properties
# kafka-run:
# 	bin/kafka-server-start.sh config/server.properties
# 	bin/windows/kafka-server-start.bat config/server.properties
