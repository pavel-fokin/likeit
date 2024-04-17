.PHONY: build-frontend
build-frontend:
	$(MAKE) -C likeit-frontend build

.PHONY: tests
tests:
	${MAKE} -C likeit-backend tests

.PHONY: run
run: frontend
	@cp -r likeit-frontend/dist likeit-backend/web
	$(MAKE) -C likeit-backend run

.PHONY: build-docker
docker-build:
	@docker build -t likeit -f Dockerfile .

.PHONY: run-docker
docker-run:
	@docker run -p 8080:8080 likeit
