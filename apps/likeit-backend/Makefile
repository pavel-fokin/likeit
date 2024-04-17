.PHONY: build
build:
	@go build -o bin/likeit-service cmd/likeit-service/main.go

.PHONY: tests
tests:
	@go test -cover ./... -count=1

.PHONY: run
run: build
	@go run cmd/likeit-service/main.go