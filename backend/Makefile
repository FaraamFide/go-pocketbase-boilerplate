# Makefile for PocketBase Go application
BINARY_NAME = app
CMD_PATH = ./cmd/app/main.go
OUTPUT_DIR = ./tmp
BINARY_PATH = $(OUTPUT_DIR)/$(BINARY_NAME)

# Configuration file paths
DEV_CONFIG = internal/config/config.dev.yaml
PROD_CONFIG = internal/config/config.prod.yaml

# --- Main Targets ---
.PHONY: help
help: ## Show this help message
	@printf "Usage: make <target>\n\n"
	@printf "Targets:\n"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: run
run: ## Run the application for development using the dev config
	@echo "Running with development config: $(DEV_CONFIG)"
	@go run $(CMD_PATH) --config=$(DEV_CONFIG)

.PHONY: dev
dev: ## Run with hot-reload using Air
	@if ! command -v air >/dev/null 2>&1; then \
		printf "\033[31m✘ Error: 'air' not installed!\033[0m\n"; \
		printf "  Install: go install github.com/air-verse/air@latest\n"; \
		printf "  Docs:    https://github.com/air-verse/air\n"; \
		exit 1; \
	fi
	@printf "\033[32m✓ Starting Air (hot reload)...\033[0m\n"
	@env PB_DATA_DIR=$(PB_DATA_DIR) air


.PHONY: build
build: ## Build production-ready binary
	@printf "Building production binary...\n"
	@mkdir -p $(OUTPUT_DIR)
	@CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o $(BINARY_PATH) $(CMD_PATH)
	@printf "Build complete: %s\n" "$(BINARY_PATH)"

.PHONY: run-prod
run-prod: build ## Build and run production binary
	@printf "Starting production server...\n"
	@$(BINARY_PATH) --config=$(PROD_CONFIG)

.PHONY: clean
clean: ## Remove build artifacts
	@printf "Cleaning build artifacts...\n"
	@rm -rf $(OUTPUT_DIR)
	@printf "Done.\n"

.DEFAULT_GOAL := help