package calculator_test

import (
	"github.com/dmisiuk/goose-tui-calculator/internal/calculator"
	"testing"
)

func TestIntegrationExample(t *testing.T) {
	m := calculator.New()
	m, _ = m.HandleButtonPress("1")
	m, _ = m.HandleButtonPress("+")
	m, _ = m.HandleButtonPress("2")
	m, _ = m.HandleButtonPress("=")
	if m.Display() != "3" {
		t.Errorf("Expected display 3, got %s", m.Display())
	}
}

func TestAudioIntegration(t *testing.T) {
	// Test that button presses don't panic with audio integration
	m := calculator.New()

	// Test number buttons
	numberButtons := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, btn := range numberButtons {
		m, _ = m.HandleButtonPress(btn)
	}

	// Test functional buttons
	functionalButtons := []string{"+", "-", "x", "/", "%", "+/-", "."}
	for _, btn := range functionalButtons {
		m, _ = m.HandleButtonPress(btn)
	}

	// Test special action buttons
	specialButtons := []string{"AC", "="}
	for _, btn := range specialButtons {
		m, _ = m.HandleButtonPress(btn)
	}
}

func TestAudioWithCalculations(t *testing.T) {
	// Test that audio doesn't interfere with calculations
	m := calculator.New()

	// Test multiple calculations with audio
	tests := []struct {
		buttons  []string
		expected string
	}{
		{[]string{"5", "+", "3", "="}, "8"},
		{[]string{"AC", "9", "-", "4", "="}, "5"},
		{[]string{"AC", "6", "x", "7", "="}, "42"},
		{[]string{"AC", "8", "/", "2", "="}, "4"},
	}

	for _, tt := range tests {
		for _, btn := range tt.buttons {
			m, _ = m.HandleButtonPress(btn)
		}
		if m.Display() != tt.expected {
			t.Errorf("Expected %s, got %s for buttons %v", tt.expected, m.Display(), tt.buttons)
		}
	}
}
