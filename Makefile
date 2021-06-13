.PHONY: help
help: ## Prints out the options available in this makefile
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: solve
solve: ## Run the solver for all puzzles
	go run .

.PHONY: solveall
solveall: ## Run the solver for all puzzles
	go run . -all

.PHONY: test
test: ## Run unit tests
	go test -timeout=30s -short ./...

.PHONY: lint
lint: ## Runs linters (via golangci-lint) on golang code
	golangci-lint run -v ./...

.PHONY: puzzle
puzzle: ## generates the template for the puzzle
	# See https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet
	# Or  https://golang.org/pkg/text/template/
	go run tool/main.go