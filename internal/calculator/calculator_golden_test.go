package calculator

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/exp/golden"
	"github.com/dmisiuk/goose-tui-calculator/internal/testutil"
	"github.com/stretchr/testify/require"
)

// TestCalculatorInitialRender validates the initial UI state
func TestCalculatorInitialRender(t *testing.T) {
	m := New()
	output := m.View()
	outBytes := []byte(output)
	
	// Validate with testutil helpers
	results := testutil.ValidateTerminalOutput(outBytes)
	for _, result := range results {
		if !result.Passed {
			t.Errorf("Validation failed: %s - %s", result.Name, result.Message)
		}
	}
	
	// Check golden file
	golden.RequireEqual(t, outBytes)
}

// TestCalculatorBasicOperation validates "2 + 3 = 5" calculation
func TestCalculatorBasicOperation(t *testing.T) {
	m := New()
	
	// Perform calculation: 2 + 3 = 5
	updatedModel, _ := m.handleButtonPress("2")
	m = updatedModel.(model)
	updatedModel, _ = m.handleButtonPress("+")
	m = updatedModel.(model)
	updatedModel, _ = m.handleButtonPress("3")
	m = updatedModel.(model)
	updatedModel, _ = m.handleButtonPress("=")
	m = updatedModel.(model)
	
	output := m.View()
	outBytes := []byte(output)
	
	// Validate the result
	require.Contains(t, output, "5", "Output should contain result '5'")
	require.Contains(t, output, "2 + 3", "Output should show the operation")
	
	// Validate with testutil helpers
	results := testutil.ValidateTerminalOutput(outBytes)
	require.True(t, testutil.AllPassed(results), "All validations should pass:\n%s", testutil.FormatResults(results))
	
	// Check golden file
	golden.RequireEqual(t, outBytes)
}

// TestCalculatorUIElements validates all UI components are present
func TestCalculatorUIElements(t *testing.T) {
	m := New()
	output := m.View()
	
	// Check for GOOSE logo/title
	require.Contains(t, output, "GOOSE", "Output should contain GOOSE title")
	
	// Check for all buttons
	buttons := []string{"AC", "+/-", "%", "/", "7", "8", "9", "x", "4", "5", "6", "-", "1", "2", "3", "+", "0", ".", "="}
	for _, btn := range buttons {
		require.Contains(t, output, btn, "Output should contain button '%s'", btn)
	}
	
	// Check for border characters
	hasBorder := strings.ContainsAny(output, "─│┌┐└┘╭╮╯╰")
	require.True(t, hasBorder, "Output should contain border characters")
	
	// Check for display showing initial "0"
	require.Contains(t, output, "0", "Output should contain initial display value '0'")
	
	// Check golden file
	golden.RequireEqual(t, []byte(output))
}

// TestCalculatorColors validates ANSI color codes are present
func TestCalculatorColors(t *testing.T) {
	m := New()
	output := m.View()
	
	// Check for ANSI escape sequences (colors)
	require.Contains(t, output, "\x1b[", "Output should contain ANSI color escape sequences")
	
	// Verify color codes with regex pattern
	results := testutil.ValidateTerminalOutput([]byte(output))
	colorResult := results[0] // First result is ANSI Color Codes
	require.True(t, colorResult.Passed, "ANSI color codes validation should pass")
	
	// Check golden file
	golden.RequireEqual(t, []byte(output))
}
