package calculator

import (
	"testing"
)

func TestPreviousOperationDisplay(t *testing.T) {
	m := New()

	// Helper function to simulate button presses
	press := func(button string) model {
		updatedModel, _ := m.handleButtonPress(button)
		m = updatedModel.(model)
		return m
	}

	// 1. Initial state
	if m.previousDisplay != "" {
		t.Errorf("Initial previousDisplay should be empty, got '%s'", m.previousDisplay)
	}

	// 2. Pressing a number
	m = press("2")
	if m.previousDisplay != "" {
		t.Errorf("previousDisplay should be empty after pressing a number, got '%s'", m.previousDisplay)
	}

	// 3. Pressing an operator
	m = press("➕")
	expected := "2 ➕"
	if m.previousDisplay != expected {
		t.Errorf("previousDisplay should be '%s' after pressing an operator, got '%s'", expected, m.previousDisplay)
	}

	// 4. Pressing another number
	m = press("3")
	if m.previousDisplay != expected {
		t.Errorf("previousDisplay should remain '%s' after pressing the second operand, got '%s'", expected, m.previousDisplay)
	}

	// 5. Pressing equals
	m = press("=")
	expected = "2 ➕ 3"
	if m.previousDisplay != expected {
		t.Errorf("previousDisplay should be '%s' after pressing equals, got '%s'", expected, m.previousDisplay)
	}
	if m.display != "5" {
		t.Errorf("display should be '5', got '%s'", m.display)
	}

	// 6. Start a new calculation
	m = press("✖")
	expected = "5 ✖"
	if m.previousDisplay != expected {
		t.Errorf("previousDisplay should be '%s' for new calculation, got '%s'", expected, m.previousDisplay)
	}

	m = press("4")
	m = press("=")
	expected = "5 ✖ 4"
	if m.previousDisplay != expected {
		t.Errorf("previousDisplay should be '%s', got '%s'", expected, m.previousDisplay)
	}
	if m.display != "20" {
		t.Errorf("display should be '20', got '%s'", m.display)
	}

	// 7. Test AC
	m = press("AC")
	if m.previousDisplay != "" {
		t.Errorf("previousDisplay should be empty after pressing AC, got '%s'", m.previousDisplay)
	}
	if m.display != "0" {
		t.Errorf("display should be '0' after AC, got '%s'", m.display)
	}

	// 8. Division
	press("1")
	press("0")
	press("÷")
	press("2")
	press("=")
	expected = "10 ÷ 2"
	if m.previousDisplay != expected {
		t.Errorf("previousDisplay should be '%s' for division, got '%s'", expected, m.previousDisplay)
	}
	if m.display != "5" {
		t.Errorf("display should be '5' for division, got '%s'", m.display)
	}
}

func TestCalculationLogic(t *testing.T) {
	m := New()

	// Helper function to simulate a sequence of button presses
	calculate := func(buttons ...string) model {
		var updatedModel interface{}
		for _, btn := range buttons {
			btn, _ = mapKeyToButton(btn)
			updatedModel, _ = m.handleButtonPress(btn)
			m = updatedModel.(model)
		}
		return m
	}

	testCases := []struct {
		name     string
		buttons  []string
		expected string
	}{
		{"Addition", []string{"2", "+", "3", "="}, "5"},
		{"Subtraction", []string{"5", "-", "2", "="}, "3"},
		{"Multiplication", []string{"4", "*", "3", "="}, "12"},
		{"Division", []string{"1", "0", "/", "2", "="}, "5"},
		{"Chained operations", []string{"2", "+", "3", "=", "+", "5", "="}, "10"},
		{"Division by zero", []string{"5", "/", "0", "="}, "Error"},
		{"Clear after error", []string{"5", "/", "0", "=", "c"}, "0"},
		{"Percentage", []string{"5", "0", "%"}, "0.5"},
		{"Sign toggle", []string{"5", "~"}, "-5"},
		{"Sign toggle twice", []string{"5", "~", "~"}, "5"},
		{"Decimal input", []string{".", "5", "+", "1", "="}, "1.5"},
		{"HONK button", []string{"3", "+", "3", "H"}, "6"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m = New() // Reset model for each test case
			result := calculate(tc.buttons...)
			if result.display != tc.expected {
				t.Errorf("Expected display to be '%s', but got '%s'", tc.expected, result.display)
			}
		})
	}
}
