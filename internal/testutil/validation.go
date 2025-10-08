package testutil

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type ValidationResult struct {
	Check   string
	Passed  bool
	Details string
}

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func ValidateTerminalOutput(output []byte) []ValidationResult {
	var results []ValidationResult

	lineCount := bytes.Count(output, []byte("\n"))
	results = append(results, ValidationResult{
		Check:   "non-empty",
		Passed:  lineCount >= 5,
		Details: fmt.Sprintf("expected at least 5 lines, got %d", lineCount),
	})

	hasANSI := ansiRegex.Find(output) != nil
	results = append(results, ValidationResult{
		Check:   "ansi-colors",
		Passed:  hasANSI,
		Details: "expected ANSI color escape sequences in output",
	})

	borderRunes := []string{"─", "│", "┌", "┐", "└", "┘", "╭", "╮", "╰", "╯"}
	borderPassed := false
	for _, r := range borderRunes {
		if strings.Contains(string(output), r) {
			borderPassed = true
			break
		}
	}
	results = append(results, ValidationResult{
		Check:   "border-characters",
		Passed:  borderPassed,
		Details: "expected box drawing characters in output",
	})

	titles := []string{"GooseCalc", "Goose Calculator", "GOOSE"}
	titlePassed := false
	for _, t := range titles {
		if strings.Contains(string(output), t) {
			titlePassed = true
			break
		}
	}
	results = append(results, ValidationResult{
		Check:   "title-present",
		Passed:  titlePassed,
		Details: fmt.Sprintf("expected title keywords %v in output", titles),
	})

	return results
}
