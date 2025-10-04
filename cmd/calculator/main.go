package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dmisiuk/goose-tui-calculator/internal/calculator"
	"github.com/muesli/termenv"
)

func main() {
	configureColorProfile()

	m := calculator.New()
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}

func configureColorProfile() {
	if os.Getenv("NO_COLOR") != "" {
		return
	}

	termProgram := strings.ToLower(os.Getenv("TERM_PROGRAM"))

	if os.Getenv("GOOSE_FORCE_COLOR") != "" ||
		os.Getenv("VHS") != "" ||
		os.Getenv("FORCE_COLOR") != "" ||
		os.Getenv("CLICOLOR_FORCE") != "" ||
		strings.EqualFold(os.Getenv("COLORTERM"), "truecolor") ||
		strings.Contains(termProgram, "vhs") {
		lipgloss.SetColorProfile(termenv.TrueColor)
	}
}
