# VHS Workflow Setup Instructions

## Overview

This guide explains how to set up VHS demo recording for pull requests using the workflow files provided in `dotgithub/workflows/`.

## Files Created

### 1. `dotgithub/workflows/vhs-demo-generation.yml`
A standalone workflow that can be used to generate VHS demos on demand or automatically in PRs.

**Features:**
- Automatic demo generation for changed tape files
- Fallback to basic demo when no tapes are changed
- Manual workflow dispatch support
- PR comment generation
- Artifact upload for 30 days

**Triggers:**
- Pull requests (when `.tapes/` or Go files change)
- Manual workflow dispatch

### 2. `dotgithub/workflows/vhs-demo-integration.yml`
Integration guide for adding VHS demo generation to existing workflows.

Contains copy-paste ready code blocks for:
- VHS setup
- Demo generation
- Artifact upload
- PR comments

## Setup Instructions

### Option A: Standalone Workflow (Easy)

1. **Move the workflow file:**
   ```bash
   mv dotgithub/workflows/vhs-demo-generation.yml .github/workflows/
   ```

2. **Update workflow permissions** (if needed):
   - Go to repository Settings → Actions → General
   - Under "Workflow permissions", ensure:
     - Read and write permissions for contents
     - Read permissions for pull requests
     - Allow GitHub Actions to create and approve pull requests

3. **Test the workflow:**
   - Create a PR that changes a `.tape` file
   - Or manually trigger the workflow from Actions tab

### Option B: Integration with Existing Workflow

1. **Open your existing workflow:**
   ```bash
   # Edit .github/workflows/test-and-verify.yml
   ```

2. **Add VHS setup** (after Go setup):
   ```yaml
   - name: Setup VHS
     run: |
       curl -L https://github.com/charmbracelet/vhs/releases/latest/download/vhs_0.7.1_linux_amd64.tar.gz | tar xz
       sudo mv vhs /usr/local/bin/
       vhs version
   ```

3. **Add demo generation** (after tests):
   ```yaml
   - name: Generate VHS Demos
     if: github.event_name == 'pull_request'
     run: |
       # Copy the demo generation script from vhs-demo-integration.yml
   ```

4. **Add artifact upload:**
   ```yaml
   - name: Upload demo artifacts
     if: github.event_name == 'pull_request'
     uses: actions/upload-artifact@v3
     with:
       name: vhs-demos-${{ github.event.number }}
       path: .tapes/assets/
       retention-days: 30
   ```

### Option C: Hybrid Approach

Use the standalone workflow for demo generation and add artifact upload to your existing workflow.

## Configuration

### Environment Variables

- `GITHUB_TOKEN`: Automatically provided by GitHub Actions
- `DEMO_TYPE`: Can be set to `changed`, `all`, or `basic`

### Workflow Permissions

The workflow requires these permissions:
- `contents: read` - To checkout code
- `pull-requests: write` - To post PR comments
- `actions: read` - To upload artifacts

## Demo Generation Modes

### Changed Mode (Default)
- Detects changed `.tape` files in the PR
- Generates demos only for modified tapes
- Falls back to basic demo if no tapes changed

### All Mode
- Generates demos for all tape files
- Useful for full demo regeneration

### Basic Mode
- Only generates `calculator-basic.tape`
- Fastest option for quick validation

## Output

### Artifacts
- Generated GIF files are uploaded as workflow artifacts
- Available for 30 days
- Named `vhs-demos` or `vhs-demos-{pr-number}`

### PR Comments
- Automatic comments on PRs with demo summary
- Includes file sizes and download links
- Only posted when demos are generated

## Troubleshooting

### Common Issues

1. **VHS installation fails:**
   - Check VHS version compatibility
   - Ensure sufficient disk space

2. **No demos generated:**
   - Check if `.tape` files exist
   - Verify file permissions
   - Check VHS command syntax

3. **Artifacts not uploaded:**
   - Verify `.tapes/assets/` directory exists
   - Check file sizes (large files may exceed limits)

4. **PR comments not posted:**
   - Check GITHUB_TOKEN permissions
   - Verify PR number is available

### Debug Mode

Add this step to your workflow for debugging:
```yaml
- name: Debug Demo Generation
  run: |
    echo "Listing tape files:"
    find .tapes -name "*.tape" -type f
    echo "Checking git diff:"
    git diff --name-only origin/main...HEAD | grep tape || echo "No tape changes"
```

## Next Steps

1. **Choose setup option** (A, B, or C)
2. **Move workflow files** to `.github/workflows/`
3. **Update permissions** if needed
4. **Test with a PR** containing tape changes
5. **Monitor first run** and adjust as needed

## Support

If you encounter issues:
1. Check the workflow logs in the Actions tab
2. Review the troubleshooting section above
3. Ensure all prerequisites are met
4. Test with a simple tape file first