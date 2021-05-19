.PHONY: help
help: ## Prints out the options available in this makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: solve
solve: ## Run the solver for all puzzles
	go run .

.PHONY: test
test: ## Run unit tests
	go test -timeout=30s -short ./...

.PHONY: lint
lint: ## Runs linters (via golangci-lint) on golang code
	golangci-lint run -v ./...
