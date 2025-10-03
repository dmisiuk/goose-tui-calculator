# Goose TUI Calculator Requirements (Initial Draft)

Issue: #4
Status: Draft (v0.1)
Owner: @dmisiuk

## 1. Core Application & Technology
- Language: Go (Golang) — single static binary distribution target.
- TUI Framework: Bubble Tea (MVU architecture) for event-driven terminal UI.
- Styling: Lipgloss for theming; goal: retro Casio calculator aesthetic.
- Packaging: Go modules (go.mod). Standard project layout: /cmd, /internal, /pkg (future optional), /docs.
- Target Platforms: Linux, macOS, Windows (x86_64 + planned arm64 where supported).
- Minimum Go Version: 1.22 (adjust if project go.mod differs).

## 2. Functional Features (Initial Scope)
- Interactive calculator UI with display + button grid.
- Keyboard input for digits 0–9, decimal point, arithmetic operators (+, -, *, /), equals, clear (C / AC), backspace.
- Mouse interaction: clicking buttons triggers the same actions (where terminal supports mouse reporting).
- Visual feedback: highlight (active/pressed) state on key/button interaction.
- Audio feedback: terminal bell (configurable, can be muted via a flag or config later).
- Arithmetic operations: standard left-to-right with proper operator precedence or (initial MVP) immediate-execution model (define final approach in implementation notes).
- Supports chaining operations (e.g., 2 + 3 * 4 = ...). Exact evaluation model to be documented in code comments.
- Error handling: division by zero shows an error state (e.g., "ERR" or similar) and prevents crash.

## 3. Non-Functional Requirements
- Performance: Near-instant input handling (<16ms per event under normal usage).
- Accessibility: High-contrast default theme; avoid relying solely on color for state changes.
- Configuration: (Future) CLI flags or config file for theme, sound, precision.
- Precision: Use Go big.Float or decimal strategy later for high precision; MVP may use float64 with rounding.
- Deterministic builds: Utilize Go module sums; CI must verify `go mod tidy` produces no diff.

## 4. Project Structure & Governance
- License: MIT (root LICENSE file).
- CONTRIBUTING.md: Define workflow (branch naming, commit style, PR expectations, test & demo requirements).
- Branch Strategy: `main` (stable), feature branches (`feat/*`), docs branches (`docs/*`), fix branches (`fix/*`). Optional: `develop` if a staging branch is desired.
- Code Style: `go fmt` enforced via CI.

## 5. Testing Strategy
- Unit Tests: Core arithmetic logic and state transitions.
- Integration Tests: Interaction between model, update, and view functions.
- Visual Demo Validation: Each `.tape` (vhs) script must execute successfully in CI to detect regressions.
- Coverage: Track coverage; upload to Codecov (badge in README). Failing threshold may be introduced later (>70%).
- Race Detection: `go test -race ./...` in CI on supported OS.

## 6. E2E Visual Demos as Code
- Tool: `vhs` scripts stored in `/.tapes/*.tape`.
- Generated Assets: `/.tapes/assets/*.gif` committed for documentation & README embedding.
- Contribution Rule: New feature PRs must include updated or new tape + regenerated GIF if UI behavior changes.
- Regression Guard: CI runs all tapes to ensure they still complete.
- Automated Generation: `vhs-demo.yml` workflow uses @charmbracelet/vhs-action to auto-generate demos on PRs; artifacts uploaded for review before committing to repo.

## 7. CI/CD Workflows
- test-and-verify.yml:
  - Trigger: push + pull_request.
  - Steps: fmt check, vet, unit+integration tests, race detection, run vhs demos, collect coverage, upload to Codecov.
- vhs-demo.yml:
  - Trigger: pull_request (on main/develop branches).
  - Steps: Build calculator, run VHS action to generate demos from all `.tape` files, upload GIFs as artifacts (30-day retention), post PR comment with download links.
  - Purpose: Automate demo generation for PR review; contributors download approved demos and commit to `.tapes/assets/`.
- release.yml:
  - Trigger: git tag push matching `v*`.
  - Steps: cross-compile (Linux, macOS, Windows; add arm64 where available), attach binaries to GitHub Release, generate checksums.
- Future: Add SBOM (Syft) + provenance (SLSA) if needed.

## 8. Documentation
- README.md: Badges (build, coverage), feature list, quickstart (install/run), embedded GIF demos, contribution guide link.
- Requirements Doc: This file lives in `docs/requirements.md` and evolves (tracked via PRs referencing Issue #4 until stable v1).
- Changelog: Introduce `CHANGELOG.md` once first release candidate is cut (Keep a Changelog format).

## 9. Logging & Telemetry (Future)
- MVP: Minimal logging (only fatal or debug behind flag).
- Future: Add optional structured logging for debugging.

## 10. Release & Distribution
- Versioning: Semantic Versioning (MAJOR.MINOR.PATCH).
- Artifacts: Prebuilt binaries + checksums + optionally Homebrew tap / Scoop manifest later.
- Roadmap Placeholder: Add backlog section in README or GitHub Projects.

## 11. Security & Quality Gates (Future Enhancements)
- Static Analysis: `golangci-lint` integration later.
- Security: `gosec` optional scan.
- Supply Chain: Verify dependencies (Dependabot alerts enabled in repo settings).

## 12. Backlog / Future Ideas
- Scientific functions (sin, cos, memory store/recall).
- Expression history panel.
- Configurable themes + dark/light variants.
- Plugin system for custom operations.
- WASM build for web-embedded demo.

## 13. Acceptance Criteria for This Document (Issue #4)
- File added under `docs/requirements.md`.
- Linked from README.
- Reviewed & approved via PR.
- Future changes require PR referencing the issue (until closed).

---
Initial draft created to capture the baseline scope used to build the first iteration of the application.
