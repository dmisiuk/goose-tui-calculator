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
	colorterm := strings.ToLower(os.Getenv("COLORTERM"))
	termProgram := strings.ToLower(os.Getenv("TERM_PROGRAM"))

	switch {
	case strings.Contains(colorterm, "truecolor"), os.Getenv("VHS") != "", termProgram == "vhs":
		lipgloss.SetColorProfile(termenv.TrueColor)
	case strings.Contains(colorterm, "256"):
		lipgloss.SetColorProfile(termenv.ANSI256)
	}
}
