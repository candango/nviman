.DEFAULT_GOAL := help

.PHONY: clean help test tidy

## Display this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@awk '/^[a-zA-Z0-9_-]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 1, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  %-15s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

## Remove build artifacts and test cache
clean:
	go clean -testcache
	rm -rf dist/
	rm -rf bin/

## Run unit tests with race detection and no cache
test:
	CGO_ENABLED=0 go clean -testcache && go test -race -v  ./...

## Add missing and remove unused modules
tidy:
	go mod tidy
