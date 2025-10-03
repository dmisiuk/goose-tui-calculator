# VHS Demo Recording Workflow

This document describes the automated VHS demo recording system for the Goose TUI Calculator project.

## Overview

The VHS demo recording workflow automatically generates demo GIFs for all pull requests, ensuring that visual changes are captured and reviewed before merging.

## How It Works

### Workflow File
Location: `.github/workflows/vhs-demo.yml`

### Triggers
The workflow runs automatically when:
- A pull request is opened to `main` or `develop` branches
- A pull request is updated (new commits pushed)
- Changes affect any of these paths:
  - `**.go` - Any Go source files
  - `.tapes/*.tape` - VHS tape scripts
  - `go.mod` or `go.sum` - Go dependencies

### Workflow Steps

1. **Checkout Code** - Gets the PR branch code
2. **Setup Go** - Installs Go 1.24
3. **Build Calculator** - Compiles the calculator binary (`calc`)
4. **Discover Tapes** - Finds all `.tape` files in `.tapes/` directory
5. **Run VHS** - Uses `@charmbracelet/vhs-action@v2` to record demos
6. **Upload Artifacts** - Stores generated GIFs as workflow artifacts
7. **Comment on PR** - Posts a summary comment with download links

## Storage Architecture

### Three-Tier Storage Model

#### 1. Repository Storage (`.tapes/assets/`)
**Purpose:** Source of truth for committed demos

**Characteristics:**
- Permanent storage
- Version controlled
- Used in README and documentation
- Reviewed and approved demos only

**When to commit here:**
- After verifying generated demos are correct
- When visual behavior changes are finalized
- Before merging PR to main

#### 2. Workflow Artifacts
**Purpose:** Temporary validation and review

**Characteristics:**
- 30-day retention period
- Automatically generated on every PR
- No manual intervention needed
- Cleaned up automatically

**Use cases:**
- Review demos before committing
- Compare with existing demos
- Verify visual correctness
- Iterate on tape scripts

#### 3. PR Descriptions
**Purpose:** Visual evidence for code review

**Characteristics:**
- Embedded using GitHub raw URLs
- Links to committed assets
- Before/After comparisons
- Reviewer-facing documentation

**Format:**
```markdown
## Before
![Before Demo](https://raw.githubusercontent.com/USER/REPO/BRANCH/.tapes/assets/feature-old.gif)

## After
![After Demo](https://raw.githubusercontent.com/USER/REPO/BRANCH/.tapes/assets/feature-new.gif)
```

## Using the Workflow

### For Contributors

#### Step 1: Make Your Changes
```bash
# Create feature branch
git checkout -b feat/my-feature

# Make code changes
# Update or create .tape files if needed
```

#### Step 2: Open Pull Request
```bash
git push origin feat/my-feature
# Open PR on GitHub
```

#### Step 3: Wait for Workflow
The VHS Demo Recording workflow will:
- Start automatically
- Take 2-5 minutes to complete
- Post a comment when done

#### Step 4: Review Generated Demos
1. Find the workflow comment on your PR
2. Click the artifact download link
3. Extract and review each GIF
4. Verify:
   - Visual appearance is correct
   - Timing is appropriate
   - No artifacts or glitches
   - Demonstrates the intended behavior

#### Step 5: Commit Approved Demos
If the demos look good:
```bash
# Download artifact and extract
unzip vhs-demo-recordings.zip -d /tmp/demos

# Copy to repository
cp /tmp/demos/*.gif .tapes/assets/

# Commit
git add .tapes/assets/*.gif
git commit -m "chore: update demo recordings"
git push
```

#### Step 6: Update PR Description
Add Before/After sections using the committed GIFs:
```markdown
## Before
![Before](https://raw.githubusercontent.com/dmisiuk/goose-tui-calculator/feat/my-feature/.tapes/assets/before.gif)

## After
![After](https://raw.githubusercontent.com/dmisiuk/goose-tui-calculator/feat/my-feature/.tapes/assets/after.gif)
```

### For Reviewers

#### What to Check
1. **PR Comment** - Check the automated comment for demo count and list
2. **Artifacts** - Download and review if needed
3. **PR Description** - View embedded Before/After GIFs
4. **Visual Correctness** - Verify demos match code changes
5. **Completeness** - Ensure all relevant behaviors are demonstrated

#### When to Request Changes
- Demos don't match described changes
- Missing demos for visual changes
- Timing issues (too fast/slow)
- Visual glitches or artifacts
- Incomplete demonstration of features

## Best Practices

### When to Update Tapes

**Always update tapes when:**
- Adding new UI features
- Changing visual behavior
- Modifying user interaction flows
- Fixing visual bugs
- Changing keyboard shortcuts or navigation

**No need to update tapes when:**
- Pure refactoring (no visual changes)
- Backend logic changes (no UI impact)
- Documentation-only updates
- Internal code structure changes
- Add `no-demo-needed` label to PR

### Tape Naming Conventions

Follow these patterns:
- `calculator-basic.tape` - Core baseline functionality
- `feature-<name>.tape` - New features
- `bugfix-<name>.tape` - Bug fixes with visual impact

### Demo Quality Guidelines

**Good demos:**
- Clear and focused on specific behavior
- Appropriate timing (not too fast)
- Show complete user workflows
- Demonstrate edge cases
- Professional appearance

**Avoid:**
- Rushed demos (too fast to follow)
- Incomplete workflows
- Unrelated actions
- Excessive length
- Poor timing

## Troubleshooting

### Workflow Fails

**Problem:** VHS workflow fails to complete

**Possible causes:**
- Invalid tape syntax
- Missing dependencies
- Build failures
- Timeout issues

**Solution:**
1. Check workflow logs for error messages
2. Test tape locally: `vhs .tapes/your-tape.tape`
3. Ensure calculator builds: `go build ./cmd/calculator`
4. Fix issues and push updates

### No GIFs Generated

**Problem:** Workflow completes but no artifacts uploaded

**Possible causes:**
- Tape files don't generate output
- Output path is incorrect
- VHS action failed silently

**Solution:**
1. Verify tape `Output` directive: `Output .tapes/assets/name.gif`
2. Check tape syntax
3. Test locally before pushing
4. Review workflow logs

### Demos Look Wrong

**Problem:** Generated GIFs don't show correct behavior

**Possible causes:**
- Timing issues (app not ready)
- Wrong keypresses
- Binary not up to date
- VHS environment differences

**Solution:**
1. Increase sleep times after app launch
2. Verify keypress sequences
3. Test locally first
4. Ensure binary is rebuilt in workflow

### PR Comment Not Posted

**Problem:** Workflow completes but no comment appears

**Possible causes:**
- GitHub API rate limiting
- Permissions issue
- Script error

**Solution:**
- Check workflow logs for script errors
- Verify GitHub Actions has comment permissions
- Re-run workflow if transient issue

## Configuration

### Workflow Settings

**Artifact retention:**
```yaml
retention-days: 30
```
Can be adjusted from 1-90 days based on needs.

**Workflow triggers:**
```yaml
on:
  pull_request:
    branches: [main, develop]
    paths:
      - '**.go'
      - '.tapes/*.tape'
      - 'go.mod'
      - 'go.sum'
```

**VHS action version:**
```yaml
uses: charmbracelet/vhs-action@v2
```
Update when new versions are released.

### Customization

To modify workflow behavior:

1. Edit `.github/workflows/vhs-demo.yml`
2. Test changes in a PR
3. Document changes in this file
4. Update CONTRIBUTING.md if user-facing

## Reference

### Official Documentation
- [VHS GitHub Action](https://github.com/charmbracelet/vhs-action)
- [VHS Documentation](https://github.com/charmbracelet/vhs)
- [GitHub Actions Artifacts](https://docs.github.com/en/actions/using-workflows/storing-workflow-data-as-artifacts)

### Related Documentation
- [CONTRIBUTING.md](../CONTRIBUTING.md) - Main contribution guide
- [development-workflow.md](development-workflow.md) - Detailed development process
- [vhs-best-practices.md](vhs-best-practices.md) - VHS tape writing guide

## FAQs

**Q: Do I need to commit GIFs to the repository?**
A: Yes, finalized GIFs should be committed to `.tapes/assets/` as the source of truth.

**Q: Can I skip running VHS locally?**
A: Not recommended. Always test locally before pushing to catch issues early.

**Q: What if my PR doesn't change visual behavior?**
A: Add the `no-demo-needed` label and explain in PR description why demos aren't needed.

**Q: How long are artifacts kept?**
A: 30 days. Download and commit important demos before they expire.

**Q: Can I run demos for specific tapes only?**
A: Currently all tapes run. To skip a tape temporarily, move it out of `.tapes/` or rename it.

**Q: What if the workflow is slow?**
A: Recording demos takes time. Optimize by reducing tape duration and using appropriate sleep times.

---

*Last updated: 2024 - This workflow improves CI visibility and ensures visual changes are always documented.*
