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

.PHONY: gendata
gendata: clean build
	@cd tests && for ((i=0; i<=10000; i++)); do cat ./testdata/1_in.txt >> ../bin/in1_7Mb.txt; done
	@cd tests && for ((i=0; i<=1000; i++)); do cat ../bin/in1_7Mb.txt >> ../bin/in1_7G.txt; done
	@cd tests && for ((i=0; i<=10; i++)); do cat ../bin/in1_7G.txt >> ../bin/in18G.txt; done
