package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dmisiuk/goose-tui-calculator/internal/calculator"
	"github.com/muesli/termenv"
)

func main() {
	// Force TrueColor output when COLORTERM is set to truecolor
	// This ensures colors work in VHS recordings and CI environments
	// where auto-detection may fail
	if os.Getenv("COLORTERM") == "truecolor" {
		lipgloss.SetColorProfile(termenv.TrueColor)
	}

	m := calculator.New()
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
