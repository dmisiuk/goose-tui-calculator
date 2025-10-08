package calculator

import (
	"io"
	"testing"
	"time"

	"github.com/charmbracelet/x/exp/teatest"
	"github.com/stretchr/testify/require"

	"github.com/dmisiuk/goose-tui-calculator/internal/testutil"
)

func renderCalculator(t *testing.T, input func(tm *teatest.TestModel)) []byte {
	t.Helper()

	t.Setenv("CLICOLOR_FORCE", "1")
	t.Setenv("FORCE_COLOR", "1")
	t.Setenv("COLORTERM", "truecolor")

	tm := teatest.NewTestModel(t, New(), teatest.WithInitialTermSize(80, 24))

	time.Sleep(150 * time.Millisecond)

	if input != nil {
		input(tm)
		time.Sleep(150 * time.Millisecond)
	}

	// Quit the program gracefully to capture the final render
	tm.Type("q")
	tm.WaitFinished(t, teatest.WithFinalTimeout(3*time.Second))

	reader := tm.FinalOutput(t, teatest.WithFinalTimeout(3*time.Second))
	out, err := io.ReadAll(reader)
	require.NoError(t, err)

	return out
}

func assertValidation(t *testing.T, output []byte) {
	t.Helper()

	results := testutil.ValidateTerminalOutput(output)
	for _, result := range results {
		require.Truef(t, result.Passed, "validation %s failed: %s", result.Check, result.Details)
	}
}

func TestCalculatorInitialRender(t *testing.T) {
	output := renderCalculator(t, nil)
	assertValidation(t, output)
	teatest.RequireEqualOutput(t, output)
}

func TestCalculatorBasicOperation(t *testing.T) {
	output := renderCalculator(t, func(tm *teatest.TestModel) {
		tm.Type("2+3=")
	})

	assertValidation(t, output)
	require.Contains(t, string(output), "5", "expected addition result in output")
	teatest.RequireEqualOutput(t, output)
}

func TestCalculatorUIElements(t *testing.T) {
	output := renderCalculator(t, nil)

	assertValidation(t, output)

	expectedElements := []string{"AC", "+/-", "%", "/", "x", "-", "+", "=", "Press q or esc to quit"}
	for _, element := range expectedElements {
		require.Containsf(t, string(output), element, "expected UI element %q", element)
	}

	teatest.RequireEqualOutput(t, output)
}

func TestCalculatorColors(t *testing.T) {
	output := renderCalculator(t, nil)

	assertValidation(t, output)
	require.Contains(t, string(output), "\x1b[", "expected ANSI color codes in output")

	teatest.RequireEqualOutput(t, output)
}
