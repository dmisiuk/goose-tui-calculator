# Golden Files

This directory contains golden files for automated testing of the calculator's terminal output.

## What are Golden Files?

Golden files are reference outputs that capture the expected behavior of our application. They serve as a baseline for regression testing - if the application's output changes unexpectedly, the tests will fail.

## File Types

### Calculator Golden Files
- **Location**: `testdata/` (project root)
- **Purpose**: Bubble Tea unit test output validation
- **Content**: Terminal UI snapshots with ANSI colors and formatting

### VHS Golden Files
- **Location**: `.tapes/golden/` (this directory)
- **Purpose**: VHS demo output validation
- **Content**: Text output from VHS recordings with colors and UI elements

## File Naming Convention

```
testdata/TestCalculatorInitialRender.golden          # Initial UI state
testdata/TestCalculatorBasicOperation.golden        # Basic calculation test
testdata/TestCalculatorUIElements.golden            # UI elements validation
testdata/TestCalculatorColors.golden                # Color validation

.tapes/golden/TestVHSBasicDemo.golden              # VHS basic demo
.tapes/golden/TestVHSColorValidation.golden        # VHS color validation
.tapes/golden/calculator-basic.txt                 # Calculator basic demo output
```

## How Golden Testing Works

1. **Test Generation**: When tests run with `-update` flag, golden files are created/updated
2. **Test Validation**: Normal test runs compare current output against golden files
3. **Failure Detection**: Any difference between current and expected output causes test failure
4. **Manual Review**: Failed tests require manual investigation to determine if change is intentional

## Updating Golden Files

### After Intentional Changes
When you make intentional changes to the calculator UI or behavior:

```bash
# Update all golden files
make test-update-golden

# Or using go test directly
go test -v -update ./...
```

### Review Changes
Always review what changed before committing:

```bash
# See what would change
git diff testdata/
git diff .tapes/golden/

# Check if changes are expected
make test
```

## Validation Criteria

Our golden tests validate:

### UI Elements
- ✅ Goose logo (🪿 GOOSE 🪿)
- ✅ All calculator buttons (AC, 0-9, operators, etc.)
- ✅ Display with numbers and results
- ✅ Help text and borders

### Colors
- ✅ ANSI escape sequences present
- ✅ Multiple colors for different button types
- ✅ LCD display colors (green background, light text)
- ✅ Visual feedback colors (highlight, pressed states)

### Functionality
- ✅ Initial state shows "0"
- ✅ Basic calculations work (2 + 3 = 5)
- ✅ Error handling (division by zero)
- ✅ Navigation and keyboard input

### Dimensions
- ✅ Output has sufficient lines (>5)
- ✅ Output has sufficient width (>60 chars)
- ✅ Border characters present
- ✅ Proper spacing and alignment

## Troubleshooting

### Tests Fail After Changes
1. **Determine if change is intentional**: Review the failing output
2. **If intentional**: Update golden files with `make test-update-golden`
3. **If unintentional**: Fix the code that caused the regression

### Color Issues in CI
Golden files capture color output. If colors are missing in CI:
- Check environment variables: `CLICOLOR_FORCE=1`, `FORCE_COLOR=1`
- Verify TERM setting: `TERM=xterm-256color`
- Ensure VHS tape has color settings

### Line Ending Issues
Golden files are marked as binary in `.gitattributes` to prevent line ending issues across platforms.

## Best Practices

1. **Commit golden files**: Always commit updated golden files with your changes
2. **Review diff changes**: Check `git diff` before committing golden file updates
3. **Document intentional changes**: Mention UI changes in commit messages
4. **Run tests locally**: Run `make test` before pushing changes
5. **Use descriptive names**: Golden file names should clearly indicate their purpose

## Integration with CI

The CI pipeline:
1. Runs unit tests against golden files
2. Generates VHS demos and validates output
3. Fails if golden file comparisons fail
4. Uploads test artifacts for manual review

This ensures that any UI regression is caught automatically before merge.