all:
	${MAKE} -C apps build
	@mkdir -p bin
	@cp apps/likeit-backend/bin/likeit-service bin/likeit-service
	@echo ""
	@echo "Build is done 🎉"
	@echo "Run likeit-service with ./bin/likeit-service".
	@echo ""

.PHONY: clean
clean:
	${MAKE} -C apps clean
	@rm -rf bin
