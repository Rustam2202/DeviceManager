FROM golang:1.19

WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

COPY . .

RUN go build -o ./bin/device-manager ./cmd/main.go

EXPOSE 8080

ENTRYPOINT [ "./bin/device-manager" ]
