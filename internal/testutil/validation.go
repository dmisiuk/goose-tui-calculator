package testutil

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidationResult represents the result of a validation check
type ValidationResult struct {
	Valid  bool   `json:"valid"`
	Name   string `json:"name"`
	Reason string `json:"reason,omitempty"`
}

// ValidateTerminalOutput performs comprehensive validation of terminal output
func ValidateTerminalOutput(output []byte) []ValidationResult {
	var results []ValidationResult
	outputStr := string(output)

	// 1. Check for ANSI color codes
	results = append(results, validateAnsiColors(outputStr))

	// 2. Check output is not empty (>5 lines)
	results = append(results, validateOutputNotEmpty(outputStr))

	// 3. Check for UI border characters
	results = append(results, validateUIBorders(outputStr))

	// 4. Check for calculator title "GooseCalc" or "GOOSE"
	results = append(results, validateCalculatorTitle(outputStr))

	// 5. Check for display elements
	results = append(results, validateDisplayElements(outputStr))

	// 6. Check for button elements
	results = append(results, validateButtonElements(outputStr))

	// 7. Check for help text
	results = append(results, validateHelpText(outputStr))

	return results
}

// validateAnsiColors checks for ANSI color codes in the output
func validateAnsiColors(output string) ValidationResult {
	// ANSI escape sequence patterns
	ansiPatterns := []string{
		`\x1b\[([0-9]{1,3}(;[0-9]{1,3})*)?[mH]`, // SGR and cursor movement
		`\x1b\[48;5;\d+m`,                       // 256-color background
		`\x1b\[38;5;\d+m`,                       // 256-color foreground
		`\x1b\[48;2;[\d;]+m`,                    // RGB background
		`\x1b\[38;2;[\d;]+m`,                    // RGB foreground
	}

	colorFound := false
	for _, pattern := range ansiPatterns {
		if matched, _ := regexp.MatchString(pattern, output); matched {
			colorFound = true
			break
		}
	}

	return ValidationResult{
		Valid:  colorFound,
		Name:   "ANSI Colors",
		Reason: conditionReason(colorFound, "ANSI color codes found", "No ANSI color codes detected"),
	}
}

// validateOutputNotEmpty checks that output has more than 5 lines
func validateOutputNotEmpty(output string) ValidationResult {
	lines := strings.Split(output, "\n")
	lineCount := 0

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			lineCount++
		}
	}

	hasEnoughLines := lineCount > 5

	return ValidationResult{
		Valid:  hasEnoughLines,
		Name:   "Output Not Empty",
		Reason: conditionReason(hasEnoughLines,
			fmt.Sprintf("Output has %d non-empty lines", lineCount),
			fmt.Sprintf("Output only has %d non-empty lines (need >5)", lineCount)),
	}
}

// validateUIBorders checks for UI border characters
func validateUIBorders(output string) ValidationResult {
	borderChars := []string{"─", "│", "┌", "┐", "└", "┘", "├", "┤", "┬", "┴", "┼"}
	bordersFound := 0

	for _, char := range borderChars {
		if strings.Contains(output, char) {
			bordersFound++
		}
	}

	hasBorders := bordersFound >= 3 // Require at least 3 different border characters

	return ValidationResult{
		Valid:  hasBorders,
		Name:   "UI Borders",
		Reason: conditionReason(hasBorders,
			fmt.Sprintf("Found %d border characters", bordersFound),
			fmt.Sprintf("Found only %d border characters (need ≥3)", bordersFound)),
	}
}

// validateCalculatorTitle checks for calculator title
func validateCalculatorTitle(output string) ValidationResult {
	titleVariants := []string{"🪿 GOOSE 🪿", "GOOSE", "GooseCalc"}
	titleFound := false
	var foundTitle string

	for _, title := range titleVariants {
		if strings.Contains(output, title) {
			titleFound = true
			foundTitle = title
			break
		}
	}

	return ValidationResult{
		Valid:  titleFound,
		Name:   "Calculator Title",
		Reason: conditionReason(titleFound,
			fmt.Sprintf("Found calculator title: %s", foundTitle),
			"No calculator title found (expected: 🪿 GOOSE 🪿, GOOSE, or GooseCalc)"),
	}
}

// validateDisplayElements checks for calculator display elements
func validateDisplayElements(output string) ValidationResult {
	displayIndicators := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."}
	displayFound := 0

	for _, indicator := range displayIndicators {
		if strings.Contains(output, indicator) {
			displayFound++
		}
	}

	hasDisplay := displayFound >= 5 // Require at least 5 display characters

	return ValidationResult{
		Valid:  hasDisplay,
		Name:   "Display Elements",
		Reason: conditionReason(hasDisplay,
			fmt.Sprintf("Found %d display characters", displayFound),
			fmt.Sprintf("Found only %d display characters (need ≥5)", displayFound)),
	}
}

// validateButtonElements checks for calculator button elements
func validateButtonElements(output string) ValidationResult {
	requiredButtons := []string{"AC", "+", "-", "x", "/", "=", "0"}
	buttonsFound := 0

	for _, button := range requiredButtons {
		if strings.Contains(output, button) {
			buttonsFound++
		}
	}

	hasButtons := buttonsFound >= len(requiredButtons)-1 // Allow missing 1 button

	return ValidationResult{
		Valid:  hasButtons,
		Name:   "Button Elements",
		Reason: conditionReason(hasButtons,
			fmt.Sprintf("Found %d/%d required buttons", buttonsFound, len(requiredButtons)),
			fmt.Sprintf("Found only %d/%d required buttons", buttonsFound, len(requiredButtons))),
	}
}

// validateHelpText checks for help text
func validateHelpText(output string) ValidationResult {
	helpVariants := []string{"Press q or esc to quit", "Press q", "quit", "help"}
	helpFound := false
	var foundHelp string

	for _, help := range helpVariants {
		if strings.Contains(output, help) {
			helpFound = true
			foundHelp = help
			break
		}
	}

	return ValidationResult{
		Valid:  helpFound,
		Name:   "Help Text",
		Reason: conditionReason(helpFound,
			fmt.Sprintf("Found help text: %s", foundHelp),
			"No help text found"),
	}
}

// conditionReason returns appropriate reason based on condition
func conditionReason(condition bool, trueReason, falseReason string) string {
	if condition {
		return trueReason
	}
	return falseReason
}

// PrintValidationResults prints validation results in a readable format
func PrintValidationResults(results []ValidationResult) string {
	var sb strings.Builder
	sb.WriteString("Validation Results:\n")
	sb.WriteString("==================\n")

	passed := 0
	total := len(results)

	for _, result := range results {
		status := "❌ FAIL"
		if result.Valid {
			status = "✅ PASS"
			passed++
		}

		sb.WriteString(fmt.Sprintf("%s %s: %s\n", status, result.Name, result.Reason))
	}

	sb.WriteString("==================\n")
	sb.WriteString(fmt.Sprintf("Summary: %d/%d tests passed\n", passed, total))

	if passed == total {
		sb.WriteString("🎉 All validations passed!\n")
	} else {
		sb.WriteString("❌ Some validations failed\n")
	}

	return sb.String()
}