.POSIX:

.PHONY: all
all: lint test

.PHONY: lint
lint:
	golangci-lint run --fix
	cd tests && golangci-lint run --fix

.PHONY: test
test:
	go test ./tests/...
