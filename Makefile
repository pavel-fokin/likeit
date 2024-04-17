all:
	${MAKE} -C apps build
	@mkdir -p bin
	@cp apps/likeit-backend/bin/likeit-service bin/likeit-service

.PHONY: clean
clean:
	${MAKE} -C apps clean
	@rm -rf bin
