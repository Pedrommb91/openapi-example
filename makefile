.DEFAULT_GOAL := help

export LC_ALL=en_US.UTF-8

.PHONY: help
help: ## help command
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: run
run: ## run command
	go run cmd/app/main.go

.PHONY: generate
generate: ## generate command
	go generate ./...