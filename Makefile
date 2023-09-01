.PHONY: clean
clean:
	@rm -f web/dist/*.js web/dist/*.css web/dist/*.map web/dist/*.html web/dist/*.ico

.PHONY: build-docker
build-docker:
	@docker build -t likeit -f Dockerfile .

.PHONY: web
web:
	@npm run build

.PHONY: tests
tests:
	@go test -cover ./... -count=1

.PHONY: run
run: web
	@go run cmd/likeit-server/main.go

.PHONY: run-docker
run-docker:
	@docker run likeit