# Development Workflow (Detailed)

This document expands on the high‑level process in `CONTRIBUTING.md`.

## Philosophy
Demos are **first-class artifacts**. A pull request should let a reviewer *see* the change without pulling code locally.

## Lifecycle of a Change
1. **Ideation → create Issue** with clear problem description and acceptance criteria.
2. Discussion / refinement (acceptance criteria, demo expectations).
3. **Branch created** (`feature/<short-name>`, `bugfix/<short-name>`).
4. **Implementation with tests**.
5. **Demo script updated or created** (VHS).
6. **GIF(s) regenerated** (locally or retrieved from the `vhs-demos` CI artifact).
7. **Documentation updated** (README.md, demo-history.md).
8. **PR opened** with Before / After visual evidence and reference to Issue.
9. Review / refine.
10. Merge (squash) → optional tag → release automation.

## Feature Development Checklist

When working on a new feature, ensure all of the following are completed:

### 🔄 **Git Workflow**
- [ ] **Create GitHub Issue** describing the problem, proposed solution, and acceptance criteria
- [ ] Create feature branch from main: `git checkout -b feature/<short-name>`
- [ ] Verify `.gitignore` excludes binary files (`calc`, `calculator`, etc.)
- [ ] Remove any accidentally committed binaries: `rm -f calc && git status`
- [ ] Commit changes with descriptive messages following conventional commits
- [ ] Push feature branch: `git push -u origin feature/<branch-name>`

### 💻 **Implementation**
- [ ] Implement feature following existing code patterns and conventions
- [ ] Add necessary imports and dependencies
- [ ] Handle errors appropriately
- [ ] Maintain thread-safe state management
- [ ] Ensure backward compatibility (no breaking changes)

### 🧪 **Testing Requirements**
- [ ] **Unit tests pass**: `go test ./...`
- [ ] **Build verification**: `go build -o calc ./cmd/calculator`
- [ ] **Integration testing**: Manual verification of all interaction methods
- [ ] **Edge case testing**: Test error conditions and boundary cases
- [ ] **Regression testing**: Verify existing functionality still works

### 📚 **Documentation Requirements**
- [ ] **README baseline demo**: Update main demo to showcase richest functionality
- [ ] **Feature documentation**: Add feature description with bullet points
- [ ] **Demo history**: Ensure `docs/demo-history.md` reflects the change timeline
- [ ] **VHS tape quality**: Regenerate demos after any visual fixes

### 🔍 **Pre-PR Verification**
- [ ] **All workflow steps completed**: Go through this entire checklist
- [ ] **Demo selection correct**: Use most recent feature as Before demo
- [ ] **Visual alignment verified**: No extra spacing or misalignment
- [ ] **GitHub URLs working**: Test that PR images display correctly
- [ ] **Branch up to date**: Push all commits including visual fixes
- [ ] **Issue reference ready**: Include "Fixes #<issue-number>" in PR description

## Visual Demo Design Guidelines
- Keep tapes short (< 15 seconds runtime) unless necessary.
- Show only the *minimal steps* to demonstrate the feature.
- Use a consistent font size (e.g., 18) for readability.
- Avoid randomness or timing dependencies.
- If timing is required, use `Sleep 300ms` intervals—too long slows CI.

### Visual Quality Checks
- [ ] **Alignment verification**: Check for proper alignment, no empty lines or extra spacing
- [ ] **Container sizing**: Ensure display containers fit content without extra space
- [ ] **Multi-line displays**: Verify line positioning and spacing in multi-line features
- [ ] **Manual testing**: Run the calculator locally to verify visual appearance
- [ ] **Demo accuracy**: Ensure GIF accurately represents the actual application behavior

## Tape Naming Conventions
| Purpose | Pattern | Example |
|---------|---------|---------|
| Baseline core use | calculator-basic.tape | calculator-basic.tape |
| Feature | feature-<short>.tape | feature-sign-toggle.tape |
| Bug reproduction | bugfix-<short>.tape | bugfix-decimal-rounding.tape |

## Recommended Repository Layout
```
.tapes/
  calculator-basic.tape
  feature-*.tape
  bugfix-*.tape
  assets/
    calculator-basic.gif
    feature-sign-toggle.gif
```

## Regenerating All GIFs
(Planned script — to be added later)
```bash
for t in .tapes/*.tape; do
  vhs "$t" || exit 1
done
mv ./*.gif .tapes/assets/ 2>/dev/null || true
```
A future helper script may formalize this.

## CI-rendered demo artifacts

- Every pull request triggers `@charmbracelet/vhs-action` against `.tapes/calculator.tape` inside the **Test and Verify** workflow.
- Generated GIFs in `.tapes/assets/` are uploaded automatically as the `vhs-demos` artifact when the workflow runs on pull requests.
- Reviewers and contributors can download the artifact from the workflow run page to obtain the freshly rendered demo without rebuilding locally.
- Artifacts follow the default GitHub retention policy (currently 90 days); commit the GIF to the repository if you need a permanent record.

## Pull Request Visual Sections

### Demo Selection Guidelines
- **Before demo**: Use the **most recent feature demo** from `docs/demo-history.md`, NOT the basic calculator
- **After demo**: Use your new feature demo from your feature branch
- **Proper progression**: Show logical feature evolution (e.g., enhanced-visual-feedback → enhanced-visual-feedback + previous-operation)

### GitHub Raw URL Format
Use proper GitHub raw URLs with embedded markdown images:
```markdown
## Before
![Before](https://raw.githubusercontent.com/dmisiuk/goose-tui-calculator/main/.tapes/assets/feature-[previous-feature].gif)

## After
![After](https://raw.githubusercontent.com/dmisiuk/goose-tui-calculator/feature/[your-branch]/.tapes/assets/feature-[your-feature].gif)
```

### PR Image Requirements
- [ ] **Embedded GIFs**: Use `![Alt](URL)` syntax, not just file paths
- [ ] **Branch-specific URLs**: Before from `main`, After from your feature branch
- [ ] **Demo history reference**: Check `docs/demo-history.md` for proper Before demo
- [ ] **Visual verification**: Ensure GIFs actually display in PR preview

Example PR snippet:
```markdown
## Before
Enhanced visual feedback calculator without previous operation display:
![Before](https://raw.githubusercontent.com/dmisiuk/goose-tui-calculator/main/.tapes/assets/feature-enhanced-visual-feedback.gif)

## After
Calculator with both enhanced visual feedback AND previous operation display:
![After](https://raw.githubusercontent.com/dmisiuk/goose-tui-calculator/feature/previous-operation/.tapes/assets/feature-previous-operation.gif)
```

## Demo Update Decision Tree
| Change Type | Update Tape? | Notes |
|-------------|--------------|-------|
| UI behavior change | Yes | Add or modify relevant tape |
| Pure refactor (no visible change) | No (label no-demo-needed) | Must not alter output |
| Style/theme change | Yes | Visual diff expected |
| Logic change affecting output | Yes | Show new result sequence |
| Docs only | No | |

## Optional CI Enforcement Strategy
1. Detect Go changes vs main.
2. Check for .tapes/ changes.
3. If none and no `no-demo-needed` label → fail.
4. Override allowed by maintainer via label.

## Future Enhancements
- Script: `scripts/update-demos.sh`
- Lint: ensure PR template filled (GitHub Action)
- Badge: “Visual demos passing” (custom status)

## Common Issues and Corrective Actions

### PR Demo Issues
**Problem**: PR shows file paths instead of embedded GIFs
**Solution**: Use proper markdown image syntax `![Alt](https://raw.githubusercontent.com/...)` with GitHub raw URLs

**Problem**: Using basic calculator as Before demo instead of latest feature
**Solution**: Always check `docs/demo-history.md` for the most recent feature demo

### Visual Alignment Issues
**Problem**: Extra empty lines or misaligned display elements
**Solution**: 
1. Check container height vs content needs
2. Verify padding and spacing calculations  
3. Test manually before regenerating demos
4. Compare Before/After visually for regressions

### Workflow Verification
**Problem**: Missing steps in development process
**Solution**: Use the checklist systematically - don't skip verification steps

**Problem**: Incomplete documentation updates
**Solution**: Update README baseline demo to always show richest functionality

## FAQ
**Q: Do I commit generated GIFs?**  Yes, they are reviewed artifacts.
**Q: Can I combine multiple features in one tape?** Prefer one tape per discrete behavior unless tightly coupled.
**Q: What if a feature is experimental?** Mark the Issue and optionally prefix tape with `exp-` (may be removed before release).
**Q: How do I know which demo to use as Before?** Check `docs/demo-history.md` for the chronological order and use the most recent feature.
**Q: My display looks misaligned - what should I check?** Verify container height, padding, and line spacing. Test manually before committing.

---
This workflow ensures every behavioral change is testable *and* reviewable visually.
