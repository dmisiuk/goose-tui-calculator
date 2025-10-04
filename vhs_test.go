package main_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/charmbracelet/x/exp/golden"
	"github.com/dmisiuk/goose-tui-calculator/internal/testutil"
	"github.com/stretchr/testify/require"
)

// TestVHSBasicDemo runs VHS and validates the generated output
func TestVHSBasicDemo(t *testing.T) {
	// Check if VHS is installed
	_, err := exec.LookPath("vhs")
	if err != nil {
		t.Skip("VHS not installed, skipping integration test")
	}

	// Set environment variables for color forcing
	os.Setenv("CLICOLOR_FORCE", "1")
	os.Setenv("FORCE_COLOR", "1")

	// Get the tape file path
	tapeFile := ".tapes/calculator-basic.tape"
	require.FileExists(t, tapeFile, "Tape file should exist")

	// Run VHS
	cmd := exec.Command("vhs", tapeFile)
	cmd.Env = append(os.Environ(), "CLICOLOR_FORCE=1", "FORCE_COLOR=1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("VHS output: %s", output)
		require.NoError(t, err, "VHS should run successfully")
	}

	// Read the generated .txt output
	txtFile := ".tapes/assets/calculator-basic.txt"
	require.FileExists(t, txtFile, "VHS should generate .txt output")

	txtContent, err := os.ReadFile(txtFile)
	require.NoError(t, err, "Should be able to read .txt output")

	// Validate ANSI colors present
	results := testutil.ValidateTerminalOutput(txtContent)
	require.True(t, testutil.AllValidationsPassed(results),
		"VHS output validation failed:\n%s", testutil.FormatValidationResults(results))

	// Check for specific VHS demo elements
	txtStr := string(txtContent)
	require.Contains(t, txtStr, "2", "Demo should show '2' being entered")
	require.Contains(t, txtStr, "3", "Demo should show '3' being entered")
	require.Contains(t, txtStr, "5", "Demo should show result '5'")

	// Copy to golden directory for comparison
	goldenFile := filepath.Join(".tapes", "golden", "calculator-basic.txt")
	goldenDir := filepath.Dir(goldenFile)
	err = os.MkdirAll(goldenDir, 0755)
	require.NoError(t, err, "Should be able to create golden directory")

	// Compare against golden file (or create it with -update flag)
	golden.RequireEqual(t, txtContent)
}

// TestVHSDemoColors specifically validates color presence
func TestVHSDemoColors(t *testing.T) {
	// Check if VHS is installed
	_, err := exec.LookPath("vhs")
	if err != nil {
		t.Skip("VHS not installed, skipping integration test")
	}

	// This test assumes TestVHSBasicDemo has already run
	txtFile := ".tapes/assets/calculator-basic.txt"
	if _, err := os.Stat(txtFile); os.IsNotExist(err) {
		t.Skip("VHS output not found, run TestVHSBasicDemo first")
	}

	txtContent, err := os.ReadFile(txtFile)
	require.NoError(t, err)

	// Validate colors
	results := testutil.ValidateTerminalOutput(txtContent)
	for _, result := range results {
		if result.Name == "ANSI Colors Present" {
			require.True(t, result.Passed,
				"VHS demo should contain ANSI colors. Run VHS with CLICOLOR_FORCE=1")
		}
	}
}

// TestVHSDemoUICompleteness validates UI elements are present
func TestVHSDemoUICompleteness(t *testing.T) {
	// Check if VHS is installed
	_, err := exec.LookPath("vhs")
	if err != nil {
		t.Skip("VHS not installed, skipping integration test")
	}

	// This test assumes TestVHSBasicDemo has already run
	txtFile := ".tapes/assets/calculator-basic.txt"
	if _, err := os.Stat(txtFile); os.IsNotExist(err) {
		t.Skip("VHS output not found, run TestVHSBasicDemo first")
	}

	txtContent, err := os.ReadFile(txtFile)
	require.NoError(t, err)

	// Validate UI completeness
	results := testutil.ValidateTerminalOutput(txtContent)
	for _, result := range results {
		if result.Name == "UI Borders Present" {
			require.True(t, result.Passed,
				"VHS demo should show complete UI with borders")
		}
		if result.Name == "Sufficient Content" {
			require.True(t, result.Passed,
				"VHS demo should have sufficient content (not clipped)")
		}
	}
}
