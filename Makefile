.PHONY: help

help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build-project:  ## Build the project
	@echo "Building..."
	# Create bin directory if it doesn't exist
	mkdir -p bin
	go build -o bin/exercises

run-build:  ## Build and run the project
	@echo "Running..."
	./bin/exercises

run:  ## Run the project
	@echo "Running..."
	go run ./main.go $(ARGS)
