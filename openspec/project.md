# Project Context

## Purpose
Goose TUI Calculator is a Casio-style terminal user interface calculator built with Go and Bubble Tea. The project emphasizes visual clarity, testability, and treats **visual demos as versioned artifacts**. The calculator features authentic retro LCD design with dual Goose (🪿) branding and provides a delightful terminal-based calculation experience.

## Tech Stack
- **Language**: Go 1.24.0 - single static binary distribution
- **TUI Framework**: Bubble Tea (MVU architecture) for event-driven terminal UI
- **UI Components**: Charmbracelet Bubbles for reusable TUI components
- **Styling**: Lipgloss for theming and retro Casio calculator aesthetic
- **Demo Tool**: VHS for creating visual demo tapes (.tape scripts → .gif outputs)
- **Package Management**: Go modules (go.mod)
- **Target Platforms**: Linux, macOS, Windows (x86_64 + planned arm64)

## Project Conventions

### Code Style
- **Formatting**: Enforced via `go fmt` - all code must be formatted before commit
- **Linting**: `go vet` enforced in CI
- **File Organization**: Standard Go project layout:
  - `/cmd/calculator` - main application entry point
  - `/internal/calculator` - internal calculator logic and TUI model
  - `/docs` - documentation files
  - `/.tapes` - VHS demo scripts and generated GIF assets
- **Naming**: Standard Go conventions (PascalCase for exports, camelCase for internal)
- **No Debug Code**: No stray debug prints or commented-out code in PRs
- **Dependencies**: `go mod tidy` must produce no diff

### Architecture Patterns
- **MVU Pattern**: Bubble Tea's Model-View-Update architecture
  - Model: Calculator state (display, buttons, selected index, operator state)
  - Update: Event handlers for keyboard/mouse input
  - View: Rendering display and button grid with Lipgloss styles
- **Pure Functions**: Calculation logic prefers pure functions for easier testing
- **State Management**: Internal state machine for calculator operations
- **Visual Feedback System**: Distinguishes between navigation (gold), activation (orange-red), and direct keyboard input (blue/purple)
- **Timed Effects**: Visual feedback auto-clears after 300ms using tea.Cmd

### Testing Strategy
- **Unit Tests**: Core arithmetic logic and state transitions (`*_test.go`)
- **Integration Tests**: Interaction between model, update, and view functions (`*_integration_test.go`)
- **State Tests**: Calculator state machine validation (`calculator_state_test.go`)
- **Race Detection**: `go test -race ./...` in CI
- **VHS Demo Validation**: Each `.tape` script executes successfully in CI to detect visual regressions
- **Coverage**: Tracked and uploaded to Codecov (future threshold: >70%)
- **Edge Cases**: Division by zero, percent semantics, multiple decimals, sign toggling, operator chaining

### Git Workflow
- **Branch Strategy**: `main` (stable) + feature branches
- **Branch Naming Conventions**:
  - `feat/<slug>` - New features
  - `fix/<slug>` - Bug fixes
  - `docs/<slug>` - Documentation
  - `chore/<slug>` - Infrastructure/tooling
  - `refactor/<slug>` - Code refactoring
  - `test/<slug>` - Test additions
- **Commit Style**: Conventional commits (enforced socially)
  - `feat: add memory store/recall`
  - `fix: handle division by zero gracefully`
  - `docs: add contributing guide`
  - `chore: update ci workflow`
  - `test: expand percent operation cases`
  - `refactor: extract calculation engine`
- **PR Requirements**:
  - Reference issue with `Closes #X`
  - Include Before/After GIF sections (or mark N/A)
  - List updated VHS tapes if UI changes
  - All CI checks must pass
- **Merge Strategy**: Squash merge preferred

## Domain Context
- **Calculator Semantics**: Immediate-execution model (left-to-right evaluation) rather than algebraic precedence
- **Visual Demo as Code**: VHS tapes are first-class artifacts, versioned and enforced in CI
- **Casio-Style UI**: Authentic retro design with specific color palette:
  - Dark green LCD display (#1B5E4F) with light green text (#D5F5E3)
  - Red AC button, gray numbers, orange operators, bright orange equals
  - 24-character alignment for display and button grid
- **Dual Input Methods**: Support both arrow key navigation + Enter/Space and direct keyboard typing
- **Previous Operation Display**: Shows calculation history on second line above current input

## Important Constraints
- **Performance**: Near-instant input handling (<16ms per event)
- **Accessibility**: High-contrast default theme; avoid relying solely on color for state
- **Deterministic Builds**: Go module sums verified; `go mod tidy` enforced
- **Precision**: Currently uses float64; future may use big.Float or decimal for high precision
- **Single Binary**: Must produce standalone static binary for distribution
- **Visual Demos Required**: UI/behavior changes MUST update or add VHS tape + regenerate GIF(s)

## External Dependencies
- **Charmbracelet Ecosystem**:
  - `bubbletea` v1.3.9 - TUI framework
  - `bubbles` v0.18.0 - TUI components
  - `lipgloss` v1.1.0 - styling and layout
- **VHS**: Command-line tool for generating terminal GIFs from tape scripts
- **GitHub Actions**: CI/CD workflows for testing, VHS demo generation, and releases
- **Codecov**: Code coverage tracking (future integration)
- **GoReleaser**: Binary release automation (planned)
