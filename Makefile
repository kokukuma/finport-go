all: build
run:
	go run cmd/server/main.go
build:
	go build -o app cmd/server/main.go
