.POSIX:

.PHONY: all
all: lint test

.PHONY: lint
lint:
	goimports --local github.com/sagmor/fun -w .
	golangci-lint run --fix
	cd tests && golangci-lint run --fix

.PHONY: test
test:
	go test ./tests/... -covermode=atomic  -coverpkg ./...
