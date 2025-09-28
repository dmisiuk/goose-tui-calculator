package calculator

import (
	"testing"
)

func TestHandleButtonPress(t *testing.T) {
	tests := []struct {
		name          string
		initialModel  model
		button        string
		expectedDisplay string
	}{
		{"Initial number press", New(), "7", "7"},
		{"Append number", model{display: "7"}, "8", "78"},
		{"Add operation", model{display: "78"}, "+", "78"},
		{"Second operand", model{display: "78", operand1: "78", operator: "+", isOperand2: true}, "9", "9"},
		{"Calculation", model{display: "9", operand1: "78", operator: "+"}, "=", "87"},
		{"Clear", model{display: "87"}, "AC", "0"},
		{"Division", model{display: "2", operand1: "10", operator: "/"}, "=", "5"},
		{"Division by zero", model{display: "0", operand1: "5", operator: "/"}, "=", "Error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := tt.initialModel.handleButtonPress(tt.button)
			if m.(model).display != tt.expectedDisplay {
				t.Errorf("expected display to be %s, but got %s", tt.expectedDisplay, m.(model).display)
			}
		})
	}
}
