# Golden Test Files

This directory contains golden test files for the GooseCalc calculator.

## What are Golden Files?

Golden files are reference outputs that tests compare against. They capture the expected terminal output of the calculator, including:

- ANSI color codes
- UI layout and borders
- Button labels and display values
- Complete visual state

## Workflow

### Running Tests

```bash
# Run all tests
make test

# Run unit tests only
make test-unit

# Run VHS integration tests
make test-integration
```

### Updating Golden Files

When you intentionally change the UI or output, update the golden files:

```bash
# Update all golden files
make test-update-golden

# Or use go test directly
CLICOLOR_FORCE=1 FORCE_COLOR=1 COLORTERM=truecolor go test -update ./...
```

### Golden File Locations

- `internal/calculator/testdata/*.golden` - Unit test golden files
- `.tapes/golden/*.txt` - VHS integration test outputs
- `testdata/*.golden` - Root-level test golden files

## When to Update

Update golden files when:
- ✅ UI layout changes (intentional)
- ✅ Color scheme updates
- ✅ Button labels change
- ✅ Display formatting improves

**Do NOT update** if:
- ❌ Tests fail due to regression
- ❌ Colors disappear unexpectedly
- ❌ UI elements are clipped

## CI/CD

Golden files are committed to git and validated in CI:
- Unit tests run on every PR
- VHS integration tests validate demo recordings
- Failures indicate visual regressions

## Troubleshooting

### Test fails with "golden file not found"

Run: `make test-update-golden`

### Colors not appearing in tests

Ensure environment variables are set:
```bash
export CLICOLOR_FORCE=1
export FORCE_COLOR=1
export COLORTERM=truecolor
```

### Golden file has wrong line endings

Check `.gitattributes` marks golden files as binary.

## Learn More

- [Golden File Testing Guide](https://ro-che.info/articles/2017-12-04-golden-tests)
- [teatest Documentation](https://pkg.go.dev/github.com/charmbracelet/x/exp/teatest)
