# Goose TUI Calculator

A retro-styled terminal user interface calculator built with Go and Bubble Tea.

*Work in progress.*

![Baseline Demo](.tapes/assets/feature-display-alignment.gif)

## Features

### Display Alignment & Professional UI
Perfect visual alignment and professional appearance with rounded border containment:

![Display Alignment Demo](.tapes/assets/feature-display-alignment.gif)

- **Perfect alignment** - Display container matches keyboard width exactly (20 characters)
- **Consistent button rows** - All keyboard rows have uniform alignment with no overflow
- **Professional border** - Rounded border around entire calculator for visual containment  
- **Optimized layout** - Help text split into efficient 2-line format
- **Space efficiency** - No wasted empty space on right side of interface
- **Maintains functionality** - All existing features preserved with improved presentation

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

## Contributing

We follow an **Issue → Branch → Code + Tests → VHS Demo → PR** workflow.

Key points:
- Every feature or bug fix starts with an issue.
- UI or behavior changes require updating/adding a VHS tape and regenerated GIF.
- PRs must include Before / After GIF sections (or mark Before as N/A).
- See [CONTRIBUTING.md](CONTRIBUTING.md) and detailed [development workflow](docs/development-workflow.md).

If you want to propose a feature, open a *Feature Request* issue (template provided).

---
Visual demos are treated as versioned artifacts to keep reviews fast and transparent.
See demo history: [docs/demo-history.md](docs/demo-history.md)
