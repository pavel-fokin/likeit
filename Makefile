build:
	${MAKE} -C apps build
	@mkdir -p bin
	@cp apps/likeit-backend/bin/likeit-service bin/likeit-service
	@echo ""
	@echo "Build is done 🎉"
	@echo "Run likeit-service with ./bin/likeit-service"
	@echo ""

.PHONY: clean
clean:
	${MAKE} -C apps clean
	@rm -rf bin

.PHONY: run-be
run-be:
	-cd apps/likeit-backend && go run cmd/likeit-service/main.go

.PHONY: run-fe
run-fe:
	-cd apps/likeit-frontend && bun run dev