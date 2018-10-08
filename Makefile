all: build
run:
	@go run cmd/server/main.go
build:
	@go build -o app cmd/server/main.go
test:
	@go test -cover -race ./...
cov:
	@go test -race -coverpkg=./... -coverprofile=coverage.txt ./...
	@go tool cover -html=coverage.txt -o coverage.html
