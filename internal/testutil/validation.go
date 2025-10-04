package testutil

import (
	"regexp"
	"strings"
)

// ValidationResult represents the result of a validation check
type ValidationResult struct {
	Name    string
	Passed  bool
	Message string
}

// ValidateTerminalOutput performs comprehensive validation of terminal output
func ValidateTerminalOutput(output []byte) []ValidationResult {
	outputStr := string(output)
	results := []ValidationResult{}

	// Check for ANSI color codes
	ansiPattern := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	hasColors := ansiPattern.MatchString(outputStr)
	results = append(results, ValidationResult{
		Name:    "ANSI Color Codes",
		Passed:  hasColors,
		Message: "Output should contain ANSI color escape sequences",
	})

	// Check output is not empty (>5 lines)
	lines := strings.Split(outputStr, "\n")
	hasContent := len(lines) > 5
	results = append(results, ValidationResult{
		Name:    "Output Length",
		Passed:  hasContent,
		Message: "Output should contain more than 5 lines",
	})

	// Check for UI border characters
	borderChars := []string{"─", "│", "┌", "┐", "└", "┘", "╭", "╮", "╯", "╰"}
	hasBorders := false
	for _, char := range borderChars {
		if strings.Contains(outputStr, char) {
			hasBorders = true
			break
		}
	}
	results = append(results, ValidationResult{
		Name:    "UI Borders",
		Passed:  hasBorders,
		Message: "Output should contain UI border characters",
	})

	// Check for calculator title "GooseCalc" or "GOOSE"
	hasTitle := strings.Contains(outputStr, "GooseCalc") || strings.Contains(outputStr, "GOOSE")
	results = append(results, ValidationResult{
		Name:    "Calculator Title",
		Passed:  hasTitle,
		Message: "Output should contain 'GooseCalc' or 'GOOSE' title",
	})

	return results
}

// AllPassed returns true if all validation results passed
func AllPassed(results []ValidationResult) bool {
	for _, result := range results {
		if !result.Passed {
			return false
		}
	}
	return true
}

// FormatResults formats validation results as a human-readable string
func FormatResults(results []ValidationResult) string {
	var builder strings.Builder
	for _, result := range results {
		status := "✓"
		if !result.Passed {
			status = "✗"
		}
		builder.WriteString(status + " " + result.Name + ": " + result.Message + "\n")
	}
	return builder.String()
}
