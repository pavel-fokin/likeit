.PHONY: clean
clean:
	@rm -rf web/dist/*

.PHONY: web
web:
	@npm run build

.PHONY: tests
tests:
	@go test -coverprofile=coverage.out ./... -count=1

.PHONY: run
run: web
	@go run main.go
