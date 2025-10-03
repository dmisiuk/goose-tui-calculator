# Contributing to Goose TUI Calculator

Thank you for helping build a high‑quality, visual‑first terminal calculator.

This project emphasizes: clarity, testability, and **visual demos as code**.

## Quick Start (Local Development)
```bash
git clone git@github.com:dmisiuk/goose-tui-calculator.git
cd goose-tui-calculator
go run ./cmd/calculator
```
Run tests:
```bash
go test ./...
```

## Workflow Overview (Mandatory)
1. Open an Issue (feature / bug / chore) BEFORE starting (exceptions: trivial docs / CI tweaks).
2. Create a branch using naming convention (see below).
3. Implement code + tests.
4. Update or add a VHS tape (.tapes/*.tape) if user-visible behavior changes.
5. Regenerate GIF(s) into `.tapes/assets/` (locally or by downloading the `vhs-demos` CI artifact).
6. Open a Pull Request:
   - Reference the issue (e.g., `Closes #42`).
   - Include Before / After GIF sections (or mark N/A).
   - List updated tapes.
7. Ensure CI passes (tests + demo scripts + formatting checks).
8. Address review feedback.
9. Squash merge (preferred) once approved.

## Branch Naming
| Type | Pattern | Example |
|------|---------|---------|
| Feature | feat/<slug> | feat/sign-toggle |
| Bug fix | fix/<slug> | fix/div-zero-edge |
| Documentation | docs/<slug> | docs/workflow-initial |
| Chore/Infra | chore/<slug> | chore/add-goreleaser |
| Refactor | refactor/<slug> | refactor/engine-separation |
| Tests only | test/<slug> | test/engine-percent |

Avoid very long slugs. Use lowercase and dashes.

## Commit Style
Conventional commits recommended (enforced socially, not by tooling yet):
- feat: add memory store/recall
- fix: handle division by zero gracefully
- docs: add contributing guide
- chore: update ci workflow
- test: expand percent operation cases
- refactor: extract calculation engine

## Definition of Done (Feature / Bug PR)
All must be satisfied unless marked N/A intentionally:
- [ ] Issue exists and PR references it (Closes #X)
- [ ] Branch name follows convention
- [ ] Code formatted (`go fmt`), passes `go vet`
- [ ] Tests added/updated (logic changes)
- [ ] VHS tape updated or added (user-visible changes)
- [ ] GIF(s) regenerated and embedded in PR
- [ ] README/docs updated (if user-facing behavior changes)
- [ ] No stray debug prints / commented out code
- [ ] go mod tidy produces no diff

## Visual Demo Policy (VHS)
We treat **demos as versioned artifacts**:
- Scripts live in `.tapes/*.tape`
- Generated GIFs in `.tapes/assets/*.gif`
- At least one baseline tape must always run in CI (e.g., `calculator-basic.tape`).
- Every pull request automatically re-renders `.tapes/calculator.tape` via `@charmbracelet/vhs-action` and uploads the GIF bundle as the `vhs-demos` workflow artifact for reviewers.
- Any UI/behavior change MUST:
  - Update an existing tape OR add a new one.
  - Regenerate affected GIF(s).
  - Embed GIF(s) in PR (Before / After). If no prior behavior existed, label Before: N/A.

### Creating / Updating a Tape
Example basic tape (`.tapes/calculator-basic.tape`):
```
Output calculator-basic.gif
Set FontSize 18
Type "go run ./cmd/calculator\n"
Sleep 500ms
Type "2"
Type "+"
Type "3"
Type "="
Sleep 800ms
Type "q"
```
Generate manually:
```bash
vhs .tapes/calculator-basic.tape
```
Result GIF will appear (by default) in working directory or specified path—move or configure output to `.tapes/assets/`.

### When to Add New vs Update Existing
- Minor extension of existing behavior: update baseline tape.
- New feature with distinct workflow: add `feature-<slug>.tape`.
- Visual bug reproduction/regression: add `bugfix-<slug>.tape` (optional but encouraged).

## Tests
- Calculation logic should prefer pure functions (refactor path) for easier coverage.
- Edge cases: division by zero, percent semantics, multiple decimals, sign toggling.
- Future: repeated equals behavior, operator replacement.

## CI Expectations
Current CI runs:
- Build & test
- Vet
- VHS tape(s) rendered via `@charmbracelet/vhs-action` with artifacts attached to pull requests
Future enhancements may add: race detector, coverage upload, lint, demo enforcement script.

## Demo Enforcement (Planned Option)
A CI script may fail PRs if Go code changes without `.tapes/` modifications (unless labeled `no-demo-needed`). Contributors should assume enforcement even if not yet active.

## Labels (Suggested)
- type:feature, type:bug, type:docs, type:chore
- needs-demo / no-demo-needed
- blocked
- ready-for-review

## Releases
Versioning: Semantic Versioning. Release workflow uses GoReleaser once `.goreleaser.yml` is added. Tag with `vX.Y.Z` to trigger.

## Questions / Discussion
Open an issue or start a discussion thread for architectural or workflow changes before large refactors.

## Code of Conduct
(If needed later—currently omitted; may adopt Contributor Covenant.)

---
Thank you for contributing! Visual clarity + reliable behavior are our top priorities.
