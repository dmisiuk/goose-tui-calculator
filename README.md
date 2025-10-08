# Goose TUI Calculator 🪿

A Casio-style terminal user interface calculator built with Go and Bubble Tea, featuring the Goose logo and authentic retro LCD design.

*Work in progress.*

## Comprehensive Demo

See all calculator features in action - direct keyboard input, arrow navigation, error handling, and screen clearing:

![Comprehensive Demo](.tapes/assets/feature-comprehensive-demo.gif)

*The above demo showcases:*
- **Direct keyboard input** - Type numbers and operators directly
- **Arrow key navigation** - Navigate with arrow keys and press Enter/Space
- **Error handling** - Division by zero displays error message
- **Screen clearing** - AC button clears display and history
- **Mixed input methods** - Seamlessly switch between keyboard and navigation

![Baseline Demo](.tapes/assets/feature-no-battery.gif)

## Features

### Casio-Style Design with Goose Branding
Authentic calculator appearance inspired by classic Casio calculators:

![Goose Logo Casio UI Demo](.tapes/assets/feature-no-battery.gif)

- **Dual Goose Logo 🪿** - Distinctive dual goose branding (🪿 GOOSE 🪿) at the top
- **Green LCD Display** - Authentic dark green background with light green text
- **Perfect Alignment** - Display and button grid precisely aligned at 24 chars
- **Rich Color Scheme** - Distinct colors for different key types:
  - AC: Red for clear action
  - Numbers (1-9): Dark gray
  - Zero (0): Darker blue-gray for emphasis
  - Operators (+, -, x, /): Orange
  - Equals (=): Bright orange to highlight action
  - Functional keys (+/-, %, .): Light gray
- **Casio Aesthetics** - Clean layout with proper borders
- **Simplified Help** - Essential information only: "Press q or esc to quit"

### Previous Operation Display
The calculator shows your previous operation on a second line above the current input:

![Previous Operation Demo](.tapes/assets/feature-previous-operation.gif)

- **Operation tracking** - See "2 +" while entering the second operand
- **Full history** - View complete calculation "2 + 3 = 5" after pressing equals
- **Right alignment** - Both history and current values align to the right
- **Clear integration** - AC button clears both lines for fresh start

### Enhanced Visual Feedback
Advanced visual feedback system that distinguishes between different input methods:

![Enhanced Visual Feedback Demo](.tapes/assets/feature-enhanced-visual-feedback.gif)

- **Navigation highlighting** - Gold background when navigating with arrow keys
- **Navigation activation** - Orange-red background when pressing Enter/Space or clicking
- **Direct keyboard input** - Blue/purple background when typing numbers/operators directly
- **Timed feedback** - Visual feedback automatically clears after 300ms
- **Complete accessibility** - Clear visual distinction between all interaction methods
- **Multi-input support** - Seamless switching between navigation and direct input

## Project Requirements

An initial requirements and scope document is maintained in [docs/requirements.md](docs/requirements.md). This captures:
- Core technology choices
- Functional & non-functional scope
- Testing & CI/CD strategy
- Visual demo (vhs) workflow
- Roadmap / backlog placeholders

Issue tracking the requirements document: See Issue #4.

## Testing & Golden Files

Automated tests cover both the Bubble Tea model and the VHS recording workflow.

- **Run calculator golden tests**
  ```bash
  go test ./internal/calculator -run TestCalculator
  ```
- **Regenerate golden snapshots** (after intentional UI changes)
  ```bash
  CLICOLOR_FORCE=1 FORCE_COLOR=1 go test ./internal/calculator -run TestCalculator -update
  ```
- **Execute the VHS integration test** (skips automatically when VHS is not installed)
  ```bash
  go test -run TestVHSBasicDemo ./...
  ```

Golden snapshots for the calculator live in `internal/calculator/testdata/`. The VHS integration test compares its output to `.tapes/golden/calculator-basic.txt` so remember to regenerate those assets when updating the demo scripts.

## Contributing

We follow an **Issue → Branch → Code + Tests → VHS Demo → PR** workflow.

Key points:
- Every feature or bug fix starts with an issue.
- UI or behavior changes require updating/adding a VHS tape and regenerated GIF.
- PRs must include Before / After GIF sections (or mark Before as N/A).
- See [CONTRIBUTING.md](CONTRIBUTING.md) and detailed [development workflow](docs/development-workflow.md).

If you want to propose a feature, open a *Feature Request* issue (template provided).

## Claude Integration

This repository is integrated with the Anthropic Claude code assistant using the Z.ai provider. Claude can be triggered to help with development tasks such as fixing bugs or implementing new features.

### How to Use

To trigger the Claude agent, you can use one of the following methods:

1.  **Create an issue with the `claude` label:** When you create a new issue, add the `claude` label to it.
2.  **Assign an issue to `claude-bot`:** Assign an existing issue to the user `claude-bot`.
3.  **Mention `@claude` in a comment:** In any issue or pull request, create a comment that starts with `@claude`.

### Configuration

For the Claude integration to work, a `ZAI_API_KEY` secret must be configured in the repository's settings. This secret should contain a valid API key from Z.ai.

---
Visual demos are treated as versioned artifacts to keep reviews fast and transparent.
See demo history: [docs/demo-history.md](docs/demo-history.md)
