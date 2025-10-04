.PHONY: test test-unit test-integration test-update-golden clean help

# Default target
help:
	@echo "Golden Test Framework Makefile"
	@echo ""
	@echo "Available targets:"
	@echo "  make test                 - Run all tests (unit + integration)"
	@echo "  make test-unit            - Run unit tests only"
	@echo "  make test-integration     - Run VHS integration tests"
	@echo "  make test-update-golden   - Update golden files"
	@echo "  make clean                - Remove generated files"
	@echo "  make help                 - Show this help message"

# Run all tests
test: test-unit test-integration

# Run unit tests only
test-unit:
	@echo "Running unit tests..."
	CLICOLOR_FORCE=1 FORCE_COLOR=1 COLORTERM=truecolor go test -v -short ./...

# Run VHS integration tests
test-integration:
	@echo "Running VHS integration tests..."
	@if command -v vhs >/dev/null 2>&1; then \
		CLICOLOR_FORCE=1 FORCE_COLOR=1 COLORTERM=truecolor go test -v -run TestVHS ./...; \
	else \
		echo "VHS not installed. Skipping integration tests."; \
		echo "Install VHS: brew install vhs"; \
	fi

# Update golden files
test-update-golden:
	@echo "Updating golden files..."
	CLICOLOR_FORCE=1 FORCE_COLOR=1 COLORTERM=truecolor go test -v -update ./...

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -f .tapes/assets/*.gif
	rm -f .tapes/golden/*.txt
	rm -f calc calculator
	go clean -testcache
