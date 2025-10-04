.PHONY: help test test-unit test-integration test-update-golden clean build lint format

# Default target
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Development targets
test: ## Run all tests (unit + integration)
	@echo "🧪 Running all tests..."
	@$(MAKE) test-unit
	@$(MAKE) test-integration
	@echo "✅ All tests completed!"

test-unit: ## Run unit tests only
	@echo "🔬 Running unit tests..."
	@go test -v -short ./...
	@go test -v ./internal/calculator -run TestCalculator
	@echo "✅ Unit tests completed!"

test-integration: ## Run VHS integration tests only
	@echo "🎥 Running VHS integration tests..."
	@if ! command -v vhs &> /dev/null; then \
		echo "❌ VHS not found. Please install VHS first:"; \
		echo "   brew install vhs"; \
		echo "   Or visit: https://github.com/charmbracelet/vhs"; \
		exit 1; \
	fi
	@CLICOLOR_FORCE=1 FORCE_COLOR=1 TERM=xterm-256color go test -v -run TestVHS ./...
	@echo "✅ Integration tests completed!"

test-update-golden: ## Update golden files (run after making changes)
	@echo "🔄 Updating golden files..."
	@go test -v -update ./...
	@echo "✅ Golden files updated!"

test-ci: ## Run tests in CI mode (no colors, no interaction)
	@echo "🤖 Running tests in CI mode..."
	@go test -v -short ./...
	@echo "✅ CI tests completed!"

# Build targets
build: ## Build the calculator binary
	@echo "🔨 Building calculator..."
	@go build -o bin/calculator ./cmd/calculator
	@echo "✅ Built bin/calculator"

build-all: ## Build for multiple platforms
	@echo "🔨 Building for multiple platforms..."
	@mkdir -p bin
	@GOOS=linux GOARCH=amd64 go build -o bin/calculator-linux-amd64 ./cmd/calculator
	@GOOS=darwin GOARCH=amd64 go build -o bin/calculator-darwin-amd64 ./cmd/calculator
	@GOOS=darwin GOARCH=arm64 go build -o bin/calculator-darwin-arm64 ./cmd/calculator
	@GOOS=windows GOARCH=amd64 go build -o bin/calculator-windows-amd64.exe ./cmd/calculator
	@echo "✅ Built for all platforms"

# Code quality targets
lint: ## Run linter
	@echo "🔍 Running linter..."
	@if command -v golangci-lint &> /dev/null; then \
		golangci-lint run; \
	else \
		echo "❌ golangci-lint not found. Install with: curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b \$$(go env GOPATH)/bin v1.54.2"; \
	fi

format: ## Format Go code
	@echo "📝 Formatting code..."
	@go fmt ./...
	@goimports -w .
	@echo "✅ Code formatted!"

# VHS/Demo targets
demo: ## Generate VHS demo GIFs
	@echo "🎥 Generating VHS demos..."
	@if ! command -v vhs &> /dev/null; then \
		echo "❌ VHS not found. Please install VHS first:"; \
		echo "   brew install vhs"; \
		echo "   Or visit: https://github.com/charmbracelet/vhs"; \
		exit 1; \
	fi
	@mkdir -p .tapes/assets .tapes/golden
	@CLICOLOR_FORCE=1 FORCE_COLOR=1 TERM=xterm-256color vhs .tapes/calculator-basic.tape
	@echo "✅ Demo generated in .tapes/assets/calculator-basic.gif"

demo-all: ## Generate all VHS demos
	@echo "🎥 Generating all VHS demos..."
	@if ! command -v vhs &> /dev/null; then \
		echo "❌ VHS not found. Please install VHS first"; \
		exit 1; \
	fi
	@mkdir -p .tapes/assets .tapes/golden
	@CLICOLOR_FORCE=1 FORCE_COLOR=1 TERM=xterm-256color vhs .tapes/*.tape
	@echo "✅ All demos generated!"

# Utility targets
clean: ## Clean build artifacts and generated files
	@echo "🧹 Cleaning..."
	@rm -rf bin/
	@rm -f calculator calc
	@rm -f .tapes/assets/*.gif
	@rm -f .tapes/golden/*.txt
	@rm -f test-data/
	@echo "✅ Cleaned!"

deps: ## Download and verify dependencies
	@echo "📦 Managing dependencies..."
	@go mod download
	@go mod verify
	@go mod tidy
	@echo "✅ Dependencies updated!"

# Development workflow targets
dev-setup: ## Set up development environment
	@echo "🛠️  Setting up development environment..."
	@$(MAKE) deps
	@if ! command -v vhs &> /dev/null; then \
		echo "💡 Consider installing VHS for demo generation:"; \
		echo "   brew install vhs"; \
		echo "   Or visit: https://github.com/charmbracelet/vhs"; \
	fi
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "💡 Consider installing golangci-lint for code quality:"; \
		echo "   curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b \$$(go env GOPATH)/bin v1.54.2"; \
	fi
	@echo "✅ Development setup complete!"

check: ## Run all checks (tests + lint + build)
	@echo "🔍 Running all checks..."
	@$(MAKE) test-unit
	@$(MAKE) lint
	@$(MAKE) build
	@echo "✅ All checks passed!"

# Quick targets for common tasks
run: build ## Build and run the calculator
	@./bin/calculator

quick-test: ## Quick test run (unit tests only)
	@go test -short ./...

update: test-update-golden ## Alias for test-update-golden

# Release targets
release-check: ## Run all checks for release
	@echo "🚀 Running release checks..."
	@$(MAKE) test
	@$(MAKE) lint
	@$(MAKE) build-all
	@$(MAKE) demo-all
	@echo "✅ Release checks complete!"