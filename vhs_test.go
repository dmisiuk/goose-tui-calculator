package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/charmbracelet/x/exp/golden"
	"github.com/dmisiuk/goose-tui-calculator/internal/testutil"
)

// TestVHSBasicDemo validates VHS demo generation and output
func TestVHSBasicDemo(t *testing.T) {
	// Check if VHS is installed
	if _, err := exec.LookPath("vhs"); err != nil {
		t.Skip("VHS not installed, skipping integration test")
	}

	// Set environment variables to force colors
	env := append(os.Environ(),
		"CLICOLOR_FORCE=1",
		"FORCE_COLOR=1",
		"TERM=xterm-256color",
	)

	// Create temporary directory for test output
	tempDir := t.TempDir()
	tapeFile := filepath.Join(tempDir, "test-demo.tape")
	txtOutput := filepath.Join(tempDir, "test-demo.txt")
	gifOutput := filepath.Join(tempDir, "test-demo.gif")

	// Create a test tape file
	tapeContent := `Output ` + gifOutput + `
Output ` + txtOutput + `
Set FontSize 18
Set Width 1200
Set Height 800

# Pre-build binary for stability
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Demo calculation: 2 + 3 = 5
Type "./calc"
Enter
Sleep 2000ms
Type "2"
Sleep 300ms
Type "+"
Sleep 300ms
Type "3"
Sleep 300ms
Type "="
Sleep 800ms
Type "q"
Sleep 600ms
`

	// Write tape file
	if err := os.WriteFile(tapeFile, []byte(tapeContent), 0644); err != nil {
		t.Fatalf("Failed to write tape file: %v", err)
	}

	// Run VHS with the tape file
	cmd := exec.Command("vhs", tapeFile)
	cmd.Env = env
	cmd.Dir = tempDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("VHS failed: %v, output: %s", err, string(output))
	}

	// Verify that both GIF and TXT files were created
	if _, err := os.Stat(gifOutput); os.IsNotExist(err) {
		t.Error("GIF file was not created")
	}

	if _, err := os.Stat(txtOutput); os.IsNotExist(err) {
		t.Error("TXT file was not created")
	}

	// Read and validate the TXT output
	txtContent, err := os.ReadFile(txtOutput)
	if err != nil {
		t.Fatalf("Failed to read TXT output: %v", err)
	}

	// Validate against golden file
	golden.RequireEqual(t, txtContent)

	// Run validation tests
	validationResults := testutil.ValidateTerminalOutput(txtContent)

	// Print validation results for debugging
	t.Log(testutil.PrintValidationResults(validationResults))

	// Check critical validations
	allPassed := true
	for _, result := range validationResults {
		if !result.Valid {
			allPassed = false
			t.Errorf("Validation failed: %s - %s", result.Name, result.Reason)
		}
	}

	if !allPassed {
		t.Error("One or more VHS output validations failed")
	}

	// Additional specific checks for our calculator
	outputStr := string(txtContent)

	// Check for calculation result
	if !strings.Contains(outputStr, "5") {
		t.Error("Expected calculation result '5' in VHS output")
	}

	// Check for goose logo
	if !strings.Contains(outputStr, "🪿") {
		t.Error("Expected goose emoji in VHS output")
	}

	// Check for calculator elements
	requiredElements := []string{"2", "+", "3", "=", "AC", "GOOSE"}
	for _, element := range requiredElements {
		if !strings.Contains(outputStr, element) {
			t.Errorf("Expected element '%s' in VHS output", element)
		}
	}
}

// TestVHSColorValidation specifically tests color output in VHS
func TestVHSColorValidation(t *testing.T) {
	// Check if VHS is installed
	if _, err := exec.LookPath("vhs"); err != nil {
		t.Skip("VHS not installed, skipping color validation test")
	}

	// Set environment variables to force colors
	env := append(os.Environ(),
		"CLICOLOR_FORCE=1",
		"FORCE_COLOR=1",
		"TERM=xterm-256color",
	)

	// Create temporary directory for test output
	tempDir := t.TempDir()
	tapeFile := filepath.Join(tempDir, "test-color.tape")
	txtOutput := filepath.Join(tempDir, "test-color.txt")

	// Create a simple tape file focused on color output
	tapeContent := `Output ` + txtOutput + `
Set FontSize 18
Set Width 800
Set Height 600

# Pre-build binary
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Start calculator briefly to capture colors
Type "./calc"
Enter
Sleep 1000ms
Type "q"
Sleep 500ms
`

	// Write tape file
	if err := os.WriteFile(tapeFile, []byte(tapeContent), 0644); err != nil {
		t.Fatalf("Failed to write tape file: %v", err)
	}

	// Run VHS with the tape file
	cmd := exec.Command("vhs", tapeFile)
	cmd.Env = env
	cmd.Dir = tempDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("VHS failed: %v, output: %s", err, string(output))
	}

	// Read and validate the TXT output
	txtContent, err := os.ReadFile(txtOutput)
	if err != nil {
		t.Fatalf("Failed to read TXT output: %v", err)
	}

	// Validate against golden file
	golden.RequireEqual(t, txtContent)

	// Specifically validate colors
	validationResults := testutil.ValidateTerminalOutput(txtContent)

	// Check that colors are present
	colorValidation := validationResults[0] // First validation is ANSI Colors
	if !colorValidation.Valid {
		t.Errorf("Color validation failed: %s", colorValidation.Reason)
	}

	// Additional color-specific checks
	outputStr := string(txtContent)

	// Look for ANSI escape sequences
	ansiPatterns := []string{"\x1b[", "\x1b[48;", "\x1b[38;"}
	colorFound := false
	for _, pattern := range ansiPatterns {
		if strings.Contains(outputStr, pattern) {
			colorFound = true
			break
		}
	}

	if !colorFound {
		t.Error("No ANSI color sequences found in VHS output despite color forcing")
	}
}

// TestVHSDimensions validates VHS output dimensions
func TestVHSDimensions(t *testing.T) {
	// Check if VHS is installed
	if _, err := exec.LookPath("vhs"); err != nil {
		t.Skip("VHS not installed, skipping dimensions test")
	}

	// Create temporary directory for test output
	tempDir := t.TempDir()
	tapeFile := filepath.Join(tempDir, "test-dimensions.tape")
	txtOutput := filepath.Join(tempDir, "test-dimensions.txt")

	// Create tape file with specific dimensions
	tapeContent := `Output ` + txtOutput + `
Set FontSize 18
Set Width 1200
Set Height 800

# Pre-build binary
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Start calculator
Type "./calc"
Enter
Sleep 1500ms
Type "q"
Sleep 500ms
`

	// Write tape file
	if err := os.WriteFile(tapeFile, []byte(tapeContent), 0644); err != nil {
		t.Fatalf("Failed to write tape file: %v", err)
	}

	// Run VHS
	cmd := exec.Command("vhs", tapeFile)
	cmd.Dir = tempDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("VHS failed: %v, output: %s", err, string(output))
	}

	// Read and validate the TXT output
	txtContent, err := os.ReadFile(txtOutput)
	if err != nil {
		t.Fatalf("Failed to read TXT output: %v", err)
	}

	// Validate against golden file
	golden.RequireEqual(t, txtContent)

	// Check that output has reasonable dimensions
	outputStr := string(txtContent)
	lines := strings.Split(outputStr, "\n")

	// Should have at least 20 lines for proper dimensions
	if len(lines) < 20 {
		t.Errorf("Output has only %d lines, expected at least 20", len(lines))
	}

	// Check that lines have reasonable width
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Should have at least 60 characters width
	if maxWidth < 60 {
		t.Errorf("Output max width is %d characters, expected at least 60", maxWidth)
	}
}