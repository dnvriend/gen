.PHONY: helpgo
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

generate: ## generate source
	go generate ./...

fmt: generate ## format code
	go fmt ./...

test: fmt ## test
	go test ./...

build: test ## build gen
	go build

bench: test ## run benchmarks
	cd test && go test -bench=. && cd ..

build-and-copy: ## build an copy gen to /usr/local/bin
	go build
	mv gen /usr/local/bin

