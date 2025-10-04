package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/charmbracelet/x/exp/golden"
	"github.com/dmisiuk/goose-tui-calculator/internal/testutil"
	"github.com/stretchr/testify/require"
)

// TestVHSBasicDemo validates VHS demo output
func TestVHSBasicDemo(t *testing.T) {
	// Check if VHS is installed
	_, err := exec.LookPath("vhs")
	if err != nil {
		t.Skip("VHS not installed, skipping integration test")
	}

	// Set environment variables for color forcing
	os.Setenv("CLICOLOR_FORCE", "1")
	os.Setenv("FORCE_COLOR", "1")
	os.Setenv("COLORTERM", "truecolor")
	os.Setenv("TERM", "xterm-256color")

	// Run VHS on calculator-basic.tape
	tapeFile := ".tapes/calculator-basic.tape"
	cmd := exec.Command("vhs", tapeFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run VHS: %v\nOutput: %s", err, string(output))
	}

	// Read the generated .txt output
	txtOutputPath := ".tapes/golden/calculator-basic.txt"
	txtOutput, err := os.ReadFile(txtOutputPath)
	if err != nil {
		t.Fatalf("Failed to read VHS output file: %v", err)
	}

	// Validate output is not empty
	require.NotEmpty(t, txtOutput, "VHS output should not be empty")

	// Validate with testutil helpers
	results := testutil.ValidateTerminalOutput(txtOutput)
	for _, result := range results {
		if !result.Passed {
			t.Logf("Validation warning: %s - %s", result.Name, result.Message)
		}
	}

	// Check golden file (with -update flag this will create/update the golden file)
	// Save the golden file with a custom name
	goldenTestName := "TestVHSBasicDemo"
	t.Run(goldenTestName, func(t *testing.T) {
		golden.RequireEqual(t, txtOutput)
	})

	// Verify GIF was also generated
	gifPath := ".tapes/assets/calculator-basic.gif"
	_, err = os.Stat(gifPath)
	require.NoError(t, err, "GIF file should be generated")
}
