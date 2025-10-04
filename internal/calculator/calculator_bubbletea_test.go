package calculator_test

import (
	"io"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/golden"
	"github.com/charmbracelet/x/exp/teatest"
	"github.com/dmisiuk/goose-tui-calculator/internal/calculator"
	"github.com/dmisiuk/goose-tui-calculator/internal/testutil"
	"github.com/stretchr/testify/require"
)

// TestCalculatorInitialRender validates the initial UI state
func TestCalculatorInitialRender(t *testing.T) {
	m := calculator.New()
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))
	
	// Wait for initial render
	time.Sleep(100 * time.Millisecond)
	
	// Quit the program
	tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	
	// Get the output
	outReader := tm.FinalOutput(t)
	outBytes, err := io.ReadAll(outReader)
	require.NoError(t, err)
	out := string(outBytes)
	
	// Validate output with testutil
	results := testutil.ValidateTerminalOutput(outBytes)
	require.True(t, testutil.AllValidationsPassed(results), 
		"Initial render validation failed:\n%s", testutil.FormatValidationResults(results))
	
	// Check for key UI elements
	require.Contains(t, out, "0", "Initial display should show '0'")
	require.Contains(t, out, "quit", "Should show quit instructions")
	
	// Update golden file if flag is set
	golden.RequireEqual(t, outBytes)
}

// TestCalculatorBasicOperation validates "2 + 3 = 5" calculation
func TestCalculatorBasicOperation(t *testing.T) {
	m := calculator.New()
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))
	
	// Simulate button presses: 2 + 3 =
	tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'2'}})
	time.Sleep(50 * time.Millisecond)
	
	tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'+'}})
	time.Sleep(50 * time.Millisecond)
	
	tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'3'}})
	time.Sleep(50 * time.Millisecond)
	
	tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'='}})
	time.Sleep(100 * time.Millisecond)
	
	// Quit
	tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	
	// Get the output
	outReader := tm.FinalOutput(t)
	outBytes, err := io.ReadAll(outReader)
	require.NoError(t, err)
	out := string(outBytes)
	
	// Validate the result shows "5"
	require.Contains(t, out, "5", "Calculation result should show '5'")
	
	// Update golden file if flag is set
	golden.RequireEqual(t, outBytes)
}

// TestCalculatorUIElements validates all UI components are present
func TestCalculatorUIElements(t *testing.T) {
	m := calculator.New()
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))
	
	// Wait for render
	time.Sleep(100 * time.Millisecond)
	
	// Quit
	tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	
	// Get the output
	outReader := tm.FinalOutput(t)
	outBytes, err := io.ReadAll(outReader)
	require.NoError(t, err)
	out := string(outBytes)
	
	// Check for essential UI elements
	requiredElements := []string{
		"AC",    // Clear button
		"+/-",   // Sign toggle
		"%",     // Percent
		"/",     // Division
		"x",     // Multiplication
		"-",     // Subtraction
		"+",     // Addition
		"=",     // Equals
		".",     // Decimal point
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", // All digits
	}
	
	for _, element := range requiredElements {
		require.Contains(t, out, element, "UI should contain button: %s", element)
	}
	
	// Validate with testutil
	results := testutil.ValidateTerminalOutput(outBytes)
	require.True(t, testutil.AllValidationsPassed(results),
		"UI elements validation failed:\n%s", testutil.FormatValidationResults(results))
	
	// Update golden file if flag is set
	golden.RequireEqual(t, outBytes)
}

// TestCalculatorColors validates ANSI color codes are present
func TestCalculatorColors(t *testing.T) {
	m := calculator.New()
	tm := teatest.NewTestModel(t, m, teatest.WithInitialTermSize(80, 24))
	
	// Wait for render
	time.Sleep(100 * time.Millisecond)
	
	// Quit
	tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	
	// Get the output
	outReader := tm.FinalOutput(t)
	outBytes, err := io.ReadAll(outReader)
	require.NoError(t, err)
	out := string(outBytes)
	
	// Check for ANSI escape sequences (color codes)
	require.True(t, strings.Contains(out, "\x1b["), 
		"Output should contain ANSI escape sequences for colors")
	
	// Count approximate number of color codes (should be many for a colorful UI)
	colorCodeCount := strings.Count(out, "\x1b[")
	require.Greater(t, colorCodeCount, 10, 
		"Should have multiple color codes (found %d)", colorCodeCount)
	
	// Validate with testutil
	results := testutil.ValidateTerminalOutput(outBytes)
	for _, result := range results {
		if result.Name == "ANSI Colors Present" {
			require.True(t, result.Passed, "ANSI colors validation failed: %s", result.Message)
		}
	}
	
	// Update golden file if flag is set
	golden.RequireEqual(t, outBytes)
}
