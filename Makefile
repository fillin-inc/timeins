test: ## test
	go test -v -cover ./...

lint: ## lint
	golangci-lint run ./...

benchmark: ## benchmark
	go test -bench . -benchmem

# Docker関連コマンド
docker-build: ## Build Docker image
	docker-compose build

docker-test: ## Run tests in Docker container
	docker-compose run --rm test

docker-lint: ## Run lint in Docker container
	docker-compose run --rm lint

docker-benchmark: ## Run benchmark in Docker container
	docker-compose run --rm benchmark

docker-dev: ## Start development container
	docker-compose run --rm dev /bin/sh

docker-up: ## Start all services
	docker-compose up

docker-down: ## Stop all services
	docker-compose down

docker-clean: ## Clean up Docker images and containers
	docker-compose down --rmi all --volumes --remove-orphans

help: ## help
	@grep -E '^[[:alnum:]_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
