.PHONY: clean
clean:
	@rm -f web/dist/*.js web/dist/*.css web/dist/*.map web/dist/*.html web/dist/*.ico

.PHONY: web
web:
	@npm run build

.PHONY: tests
tests:
	@go test -cover ./... -count=1

.PHONY: run
run: web
	@go run cmd/likeit-server/main.go

# Docker.
.PHONY: build-docker
build-docker:
	@docker build -t likeit -f Dockerfile .

.PHONY: run-docker
run-docker:
	@docker run -p 8080:8080 likeit
q