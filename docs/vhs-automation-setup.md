# VHS Demo Automation Setup

## Overview
This document describes the automated VHS demo recording system configured for pull requests.

## Implementation Summary

### New GitHub Workflow
**File**: `.github/workflows/vhs-demo.yml`

The workflow automatically:
1. Triggers on pull requests to `main` and `develop` branches when Go code or VHS tapes change
2. Builds the calculator application
3. Runs all VHS tapes using [@charmbracelet/vhs-action](https://github.com/charmbracelet/vhs-action)
4. Uploads generated GIFs as GitHub Actions artifacts (30-day retention)
5. Posts a comment on the PR with download links and review checklist

### Workflow Triggers
The workflow runs when PRs modify:
- `**.go` - Any Go source files
- `.tapes/**` - VHS tape scripts
- `cmd/**` - Command directory
- `internal/**` - Internal packages

### Demo Storage Strategy
**Two-tier approach:**
1. **Automatic (GitHub Artifacts)**: All generated demos uploaded as artifacts for review
2. **Manual (Repository)**: Contributors download and commit approved demos to `.tapes/assets/`

This approach allows:
- Quick iteration without polluting the repository
- Review before committing to maintain quality
- Retention of all generated demos for 30 days

## Updated Documentation

### 1. CONTRIBUTING.md
Added section: "Automated Demo Generation (VHS Action)"
- Explains how the workflow works
- Describes the demo storage strategy
- Provides step-by-step usage instructions

### 2. docs/development-workflow.md
Added section: "Automated VHS Demo Recording"
- Detailed explanation of the workflow
- Best practices for using automated demos
- Storage decision guidance
- Integration with existing development process

### 3. docs/requirements.md
Updated sections:
- **Section 6**: Added automated generation to E2E Visual Demos
- **Section 7**: Added `vhs-demo.yml` workflow documentation

### 4. .github/pull_request_template.md
Enhanced with:
- Information box about automated demo generation
- Demo Quality Checklist for contributors using automated demos
- Updated checklist item about GIF regeneration

### 5. README.md
Added brief mention in Contributing section about automatic demo generation

## How Contributors Use This

### For PR Authors
1. Create/update VHS tapes in `.tapes/` directory
2. Push changes to PR branch
3. Wait for "VHS Demo Recording" workflow to complete
4. Review generated demos in Actions artifacts
5. Download approved demos
6. Commit demos to `.tapes/assets/` directory

### For Reviewers
1. Check PR for automated workflow results
2. Download artifacts to review demos
3. Verify demo quality matches code changes
4. Request updates if demos need adjustment

## Benefits

1. **Automation**: Eliminates manual VHS installation and execution
2. **Consistency**: All demos generated in standardized CI environment
3. **Visibility**: Immediate feedback on demo generation success/failure
4. **Quality Control**: Review before committing to repository
5. **Iteration Support**: Easy to regenerate and review multiple versions

## Technical Details

### Workflow Steps
```yaml
1. Checkout PR branch
2. Set up Go 1.24
3. Build calculator application
4. Run VHS action on .tapes directory
5. Upload artifacts (30-day retention)
6. Post PR comment with details
```

### Dependencies
- `actions/checkout@v4`
- `actions/setup-go@v4`
- `charmbracelet/vhs-action@v2`
- `actions/upload-artifact@v4`
- `actions/github-script@v7`

## Future Enhancements

Potential improvements mentioned in documentation:
- Script: `scripts/update-demos.sh` for bulk regeneration
- Badge: "Visual demos passing" custom status
- Lint: Ensure PR template is properly filled
- CI enforcement: Require demo updates for Go code changes

## References

- [VHS Action Documentation](https://github.com/charmbracelet/vhs-action)
- [VHS Best Practices](./vhs-best-practices.md)
- [Development Workflow](./development-workflow.md)
- [Contributing Guide](../CONTRIBUTING.md)
