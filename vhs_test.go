package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/charmbracelet/x/exp/golden"
	"github.com/stretchr/testify/require"

	"github.com/dmisiuk/goose-tui-calculator/internal/testutil"
)

func TestVHSBasicDemo(t *testing.T) {
	vhsPath, err := exec.LookPath("vhs")
	if err != nil {
		t.Skip("vhs not installed; skipping integration test")
	}

	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("FORCE_COLOR", "1")
	t.Setenv("COLORTERM", "truecolor")

	cmd := exec.Command(vhsPath, ".tapes/calculator-basic.tape")
	cmd.Env = append(os.Environ(),
		"CLICOLOR_FORCE=1",
		"FORCE_COLOR=1",
		"COLORTERM=truecolor",
	)

	var combined bytes.Buffer
	cmd.Stdout = &combined
	cmd.Stderr = &combined

	if err := cmd.Run(); err != nil {
		t.Fatalf("vhs execution failed: %v\n%s", err, combined.String())
	}

	output := combined.Bytes()
	if len(output) == 0 {
		textPath := filepath.Join(".tapes", "golden", "calculator-basic.txt")
		var readErr error
		output, readErr = os.ReadFile(textPath)
		require.NoError(t, readErr, "expected fallback output from %s", textPath)
	}

	results := testutil.ValidateTerminalOutput(output)
	for _, result := range results {
		require.Truef(t, result.Passed, "validation %s failed: %s", result.Check, result.Details)
	}

	expected, err := os.ReadFile(filepath.Join(".tapes", "golden", "calculator-basic.txt"))
	require.NoError(t, err)
	require.Equal(t, string(expected), string(output))

	golden.RequireEqual(t, output)
}
