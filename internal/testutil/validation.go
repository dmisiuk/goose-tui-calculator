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

	// Check 1: ANSI color codes present
	ansiColorRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	hasColors := ansiColorRegex.MatchString(outputStr)
	results = append(results, ValidationResult{
		Name:    "ANSI Colors Present",
		Passed:  hasColors,
		Message: "Terminal output should contain ANSI color codes",
	})

	// Check 2: Output is not empty (more than 5 lines)
	lines := strings.Split(outputStr, "\n")
	hasContent := len(lines) > 5
	results = append(results, ValidationResult{
		Name:    "Sufficient Content",
		Passed:  hasContent,
		Message: "Output should have more than 5 lines",
	})

	// Check 3: UI border characters present
	borderChars := []string{"─", "│", "┌", "┐", "└", "┘", "╭", "╮", "╰", "╯"}
	hasBorder := false
	for _, char := range borderChars {
		if strings.Contains(outputStr, char) {
			hasBorder = true
			break
		}
	}
	results = append(results, ValidationResult{
		Name:    "UI Borders Present",
		Passed:  hasBorder,
		Message: "Output should contain UI border characters",
	})

	// Check 4: Calculator title "GooseCalc" or "Goose" present
	hasTitle := strings.Contains(outputStr, "GooseCalc") || strings.Contains(outputStr, "Goose")
	results = append(results, ValidationResult{
		Name:    "Calculator Title Present",
		Passed:  hasTitle,
		Message: "Output should contain calculator title (GooseCalc or Goose)",
	})

	return results
}

// AllValidationsPassed checks if all validation results passed
func AllValidationsPassed(results []ValidationResult) bool {
	for _, result := range results {
		if !result.Passed {
			return false
		}
	}
	return true
}

// FormatValidationResults returns a formatted string of validation results
func FormatValidationResults(results []ValidationResult) string {
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
