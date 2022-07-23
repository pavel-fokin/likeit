.PHONY: clean
clean:
	@rm -rf dist/*.js dist/*.js.map dist/index.html

.PHONY: web
web:
	@npm run build

.PHONY: run
run: web
	@go run main.go