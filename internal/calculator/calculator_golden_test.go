package calculator

import (
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/x/exp/teatest"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/golden"
)

// TestCalculatorInitialRender validates the initial UI state
func TestCalculatorInitialRender(t *testing.T) {
	m := New()

	// Create a test model with teatest
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))

	// Get the final output after the model stabilizes
	finalOut := teatest.RequireFinalModel(t, tm).View()

	// Validate against golden file
	golden.RequireEqual(t, []byte(finalOut))

	// Additional validations
	if !strings.Contains(finalOut, "🪿 GOOSE 🪿") {
		t.Error("Expected goose logo in initial render")
	}

	if !strings.Contains(finalOut, "0") {
		t.Error("Expected initial display to show '0'")
	}

	// Check for border characters
	borderChars := []string{"─", "│", "┌", "┐", "└", "┘"}
	hasBorder := false
	for _, char := range borderChars {
		if strings.Contains(finalOut, char) {
			hasBorder = true
			break
		}
	}
	if !hasBorder {
		t.Error("Expected border characters in calculator UI")
	}
}

// TestCalculatorBasicOperation validates "2 + 3 = 5" calculation
func TestCalculatorBasicOperation(t *testing.T) {
	m := New()

	// Create a test model with teatest
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))

	// Simulate typing "2 + 3 ="
	teatest.Type(tm, "2")
	time.Sleep(100 * time.Millisecond)

	teatest.Type(tm, "+")
	time.Sleep(100 * time.Millisecond)

	teatest.Type(tm, "3")
	time.Sleep(100 * time.Millisecond)

	teatest.Type(tm, "=")
	time.Sleep(200 * time.Millisecond)

	// Get the final output
	finalOut := teatest.RequireFinalModel(t, tm).View()

	// Validate against golden file
	golden.RequireEqual(t, []byte(finalOut))

	// Check if the calculation result is displayed
	if !strings.Contains(finalOut, "5") {
		t.Error("Expected calculation result '5' in display")
	}

	// Check if the operation history is displayed
	if !strings.Contains(finalOut, "2 + 3 = 5") {
		t.Error("Expected operation history '2 + 3 = 5' in display")
	}
}

// TestCalculatorUIElements validates all UI components are present
func TestCalculatorUIElements(t *testing.T) {
	m := New()

	// Create a test model with teatest
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))

	// Get the final output
	finalOut := teatest.RequireFinalModel(t, tm).View()

	// Validate against golden file
	golden.RequireEqual(t, []byte(finalOut))

	// Check for all required UI elements
	requiredElements := []string{
		"🪿 GOOSE 🪿", // Logo
		"AC", "+/-", "%", "/", // First row
		"7", "8", "9", "x", // Second row
		"4", "5", "6", "-", // Third row
		"1", "2", "3", "+", // Fourth row
		"0", ".", "=", // Fifth row
		"Press q or esc to quit", // Help text
	}

	for _, element := range requiredElements {
		if !strings.Contains(finalOut, element) {
			t.Errorf("Expected UI element '%s' in output", element)
		}
	}

	// Check for display container
	if !strings.Contains(finalOut, "0") {
		t.Error("Expected display in UI")
	}

	// Check for calculator body border
	borderChars := []string{"─", "│", "┌", "┐", "└", "┘"}
	borderFound := false
	for _, char := range borderChars {
		if strings.Contains(finalOut, char) {
			borderFound = true
			break
		}
	}
	if !borderFound {
		t.Error("Expected calculator border in UI")
	}
}

// TestCalculatorColors validates ANSI color codes are present
func TestCalculatorColors(t *testing.T) {
	m := New()

	// Create a test model with teatest
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))

	// Get the final output
	finalOut := teatest.RequireFinalModel(t, tm).View()

	// Validate against golden file
	golden.RequireEqual(t, []byte(finalOut))

	// Check for ANSI color codes (ESC sequences)
	ansiSequences := []string{
		"\x1b[",      // Standard ANSI escape sequence
		"\x1b[48;",   // Background color
		"\x1b[38;",   // Foreground color
		"\x1b[0m",    // Reset sequence
	}

	colorFound := false
	for _, sequence := range ansiSequences {
		if strings.Contains(finalOut, sequence) {
			colorFound = true
			break
		}
	}

	if !colorFound {
		t.Error("Expected ANSI color codes in calculator UI output")
	}

	// Additional check for multiple colors (should have multiple different colors)
	colorCount := 0
	for _, sequence := range ansiSequences {
		colorCount += strings.Count(finalOut, sequence)
	}

	if colorCount < 3 {
		t.Errorf("Expected multiple color codes, found %d", colorCount)
	}
}

// TestCalculatorKeyboardNavigation validates keyboard navigation functionality
func TestCalculatorKeyboardNavigation(t *testing.T) {
	m := New()

	// Create a test model with teatest
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))

	// Test navigation keys
	teatest.Send(tm, tea.KeyMsg{Type: tea.KeyRight, Runes: []rune{'l'}})
	time.Sleep(50 * time.Millisecond)

	teatest.Send(tm, tea.KeyMsg{Type: tea.KeyDown, Runes: []rune{'j'}})
	time.Sleep(50 * time.Millisecond)

	teatest.Send(tm, tea.KeyMsg{Type: tea.KeyLeft, Runes: []rune{'h'}})
	time.Sleep(50 * time.Millisecond)

	teatest.Send(tm, tea.KeyMsg{Type: tea.KeyUp, Runes: []rune{'k'}})
	time.Sleep(50 * time.Millisecond)

	// Get the final output
	finalOut := teatest.RequireFinalModel(t, tm).View()

	// Validate against golden file
	golden.RequireEqual(t, []byte(finalOut))
}

// TestCalculatorErrorHandling validates error states
func TestCalculatorErrorHandling(t *testing.T) {
	m := New()

	// Create a test model with teatest
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))

	// Simulate division by zero: "5 / 0 ="
	teatest.Type(tm, "5")
	time.Sleep(100 * time.Millisecond)

	teatest.Type(tm, "/")
	time.Sleep(100 * time.Millisecond)

	teatest.Type(tm, "0")
	time.Sleep(100 * time.Millisecond)

	teatest.Type(tm, "=")
	time.Sleep(200 * time.Millisecond)

	// Get the final output
	finalOut := teatest.RequireFinalModel(t, tm).View()

	// Validate against golden file
	golden.RequireEqual(t, []byte(finalOut))

	// Check for error state
	if !strings.Contains(finalOut, "Error") {
		t.Error("Expected 'Error' in display for division by zero")
	}
}