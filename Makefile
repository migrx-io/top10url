NAME=top10url

.PHONY: default
default: build

.PHONY: build
build:
	@CGO_ENABLED=0 go build -o bin/$(NAME) main.go

.PHONY: fmt
fmt:
	@gofmt -l -w .


.PHONY: e2e
e2e: build
	@cd tests && ./e2e.sh

.PHONY: tests
tests:
	@go test -v ./...

.PHONY: clean
clean:
	@rm -rf bin
