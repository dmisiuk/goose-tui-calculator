# Development Workflow (Detailed)

This document expands on the high‑level process in `CONTRIBUTING.md`.

## Philosophy
Demos are **first-class artifacts**. A pull request should let a reviewer *see* the change without pulling code locally.

## Lifecycle of a Change
1. Ideation → create Issue.
2. Discussion / refinement (acceptance criteria, demo expectations).
3. **Branch created** (`feature/<short-name>`, `bugfix/<short-name>`).
4. **Implementation with tests**.
5. **Demo script updated or created** (VHS).
6. **GIF(s) regenerated**.
7. **Documentation updated** (README.md, demo-history.md).
8. **PR opened** with Before / After visual evidence.
9. Review / refine.
10. Merge (squash) → optional tag → release automation.

## Feature Development Checklist

When working on a new feature, ensure all of the following are completed:

### 🔄 **Git Workflow**
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

## Visual Demo Design Guidelines
- Keep tapes short (< 15 seconds runtime) unless necessary.
- Show only the *minimal steps* to demonstrate the feature.
- Use a consistent font size (e.g., 18) for readability.
- Avoid randomness or timing dependencies.
- If timing is required, use `Sleep 300ms` intervals—too long slows CI.

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

## Pull Request Visual Sections
Example PR snippet:
```markdown
## Before
![Before](https://raw.githubusercontent.com/dmisiuk/goose-tui-calculator/<commit-or-branch>/.tapes/assets/calculator-basic.gif)

## After
![After](https://raw.githubusercontent.com/dmisiuk/goose-tui-calculator/<branch>/.tapes/assets/feature-sign-toggle.gif)
```
For new features: mark Before as `N/A`.

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

## FAQ
**Q: Do I commit generated GIFs?**  Yes, they are reviewed artifacts.
**Q: Can I combine multiple features in one tape?** Prefer one tape per discrete behavior unless tightly coupled.
**Q: What if a feature is experimental?** Mark the Issue and optionally prefix tape with `exp-` (may be removed before release).

---
This workflow ensures every behavioral change is testable *and* reviewable visually.
