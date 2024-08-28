test: ## test
	go test -v -cover ./...

lint: ## lint
	golangci-lint run ./...

benchmark: ## benchmark
	go test -bench . -benchmem

help: ## help
	@grep -E '^[[:alnum:]_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
