package audio

import "testing"

func TestGetButtonType(t *testing.T) {
	tests := []struct {
		button   string
		expected ButtonType
	}{
		// Number buttons
		{"0", ButtonTypeNumber},
		{"1", ButtonTypeNumber},
		{"2", ButtonTypeNumber},
		{"3", ButtonTypeNumber},
		{"4", ButtonTypeNumber},
		{"5", ButtonTypeNumber},
		{"6", ButtonTypeNumber},
		{"7", ButtonTypeNumber},
		{"8", ButtonTypeNumber},
		{"9", ButtonTypeNumber},

		// Special action buttons
		{"AC", ButtonTypeSpecialAction},
		{"=", ButtonTypeSpecialAction},

		// Functional buttons
		{"+", ButtonTypeFunctional},
		{"-", ButtonTypeFunctional},
		{"x", ButtonTypeFunctional},
		{"/", ButtonTypeFunctional},
		{"%", ButtonTypeFunctional},
		{"+/-", ButtonTypeFunctional},
		{".", ButtonTypeFunctional},
	}

	for _, tt := range tests {
		t.Run(tt.button, func(t *testing.T) {
			result := GetButtonType(tt.button)
			if result != tt.expected {
				t.Errorf("GetButtonType(%q) = %v, expected %v", tt.button, result, tt.expected)
			}
		})
	}
}

func TestPlayButtonSound(t *testing.T) {
	// Test that PlayButtonSound doesn't panic for different button types
	buttons := []string{"0", "1", "AC", "=", "+", "-", "x", "/", "%", "+/-", "."}

	for _, button := range buttons {
		t.Run(button, func(t *testing.T) {
			// Should not panic
			PlayButtonSound(button)
		})
	}
}

func TestPlayNumberSound(t *testing.T) {
	// Should not panic
	PlayNumberSound()
}

func TestPlayFunctionalSound(t *testing.T) {
	// Should not panic
	PlayFunctionalSound()
}

func TestPlaySpecialActionSound(t *testing.T) {
	// Should not panic
	PlaySpecialActionSound()
}
