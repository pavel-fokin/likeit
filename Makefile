.PHONY: clean
clean:
	@rm -rf web/dist/*

.PHONY: web
web:
	@npm run build

.PHONY: run
run: web
	@go run main.go
