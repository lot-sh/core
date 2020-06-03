GOPATH=$(shell go env GOPATH)

.PHONY: test
test:
	@echo "==> Running tests"
	go test -v ./...
