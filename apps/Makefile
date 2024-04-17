.PHONY: clean
clean:
	$(MAKE) -C likeit-frontend clean
	$(MAKE) -C likeit-backend clean

.PHONY: frontend
frontend:
	$(MAKE) -C likeit-frontend build

.PHONY: run
build: frontend
	@cp -r likeit-frontend/dist likeit-backend/web
	$(MAKE) -C likeit-backend build

.PHONY: tests
tests:
	${MAKE} -C likeit-backend tests

.PHONY: docker
docker:
	@docker build -t likeit -f Dockerfile .

.PHONY: run-docker
docker-run:
	@docker run -p 8080:8080 likeit
