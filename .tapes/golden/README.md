# Golden Test Files

This directory contains golden files used for VHS demo validation.

## What are Golden Files?

Golden files are reference outputs that tests compare against to detect regressions. When a test runs, it compares the actual output with the golden file. If they differ, the test fails.

## Workflow

### Running Tests

```bash
# Run all tests (compares against golden files)
make test

# Run unit tests only
make test-unit

# Run VHS integration tests only
make test-integration
```

### Updating Golden Files

When you intentionally change the UI or demo behavior:

```bash
# Update all golden files
make test-update-golden

# Or use go test directly with -update flag
CLICOLOR_FORCE=1 FORCE_COLOR=1 go test -v -update ./...
```

**Important:** Always review changes to golden files before committing!

### What to Commit

✅ **DO commit:**
- Golden text files (`.txt` in `.tapes/golden/`)
- Golden test files (`.golden` in `testdata/`)

❌ **DON'T commit:**
- Generated GIFs (`.tapes/assets/*.gif`) - these are in `.gitignore`
- Temporary test files

## File Structure

```
.tapes/golden/
├── README.md                    # This file
└── TestVHSBasicDemo.golden      # VHS demo output reference
```

## Validation Checks

Golden files help validate:
1. **ANSI Colors Present** - Ensures colorful output
2. **Sufficient Content** - Checks output isn't empty/clipped
3. **UI Borders Present** - Validates complete UI rendering
4. **Calculator Title Present** - Checks for app branding

## Troubleshooting

### Test fails with "golden file mismatch"

This means the output changed. Either:
1. A bug was introduced → fix the bug
2. Intentional change → update golden files with `make test-update-golden`

### Colors missing in golden files

Ensure you run tests with color forcing:
```bash
CLICOLOR_FORCE=1 FORCE_COLOR=1 go test ./...
```

### Golden files have wrong line endings

The `.gitattributes` file marks golden files as binary to prevent line ending conversion.

## Related Files

- `calculator_test.go` - Unit tests using golden files
- `vhs_test.go` - VHS integration tests
- `internal/testutil/validation.go` - Validation helper functions
- `Makefile` - Convenient test commands
