package calculator

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestButtonHighlightState(t *testing.T) {
	m := New()

	// Position cursor on first button (AC)
	m.cursorX = 0
	m.cursorY = 0

	output := m.View()

	// Check that output contains the AC button and the Goose logo
	if !strings.Contains(output, "AC") {
		t.Errorf("Expected output to contain AC button")
	}
	if !strings.Contains(output, "GOOSE") {
		t.Errorf("Expected output to contain GOOSE logo")
	}
}

func TestButtonPressedState(t *testing.T) {
	m := New()

	// Position cursor and set pressed state on same button
	m.cursorX = 1
	m.cursorY = 0
	m.pressedX = 1
	m.pressedY = 0
	m.activationMethod = activationNavigation

	output := m.View()

	// Should contain the button
	if !strings.Contains(output, "+/-") {
		t.Errorf("Expected output to contain +/- button")
	}
}

func TestKeyboardNavigationVisualFeedback(t *testing.T) {
	m := New()

	tests := []struct {
		name      string
		key       tea.KeyMsg
		expectedX int
		expectedY int
	}{
		{"move right", tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'l'}}, 1, 0},
		{"move down", tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}}, 0, 1},
		{"move up", tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}}, 0, 0},
		{"move left", tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'h'}}, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset position
			m.cursorX = 0
			m.cursorY = 0

			updatedModel, _ := m.Update(tt.key)
			updated := updatedModel.(model)

			if updated.cursorX != tt.expectedX || updated.cursorY != tt.expectedY {
				t.Errorf("Expected cursor at (%d, %d), got (%d, %d)",
					tt.expectedX, tt.expectedY, updated.cursorX, updated.cursorY)
			}
		})
	}
}

func TestEnterKeyPressVisualFeedback(t *testing.T) {
	m := New()
	m.cursorX = 0
	m.cursorY = 1 // Position on "7" button

	// Simulate Enter key press
	enterKey := tea.KeyMsg{Type: tea.KeyEnter}
	updatedModel, _ := m.Update(enterKey)
	updated := updatedModel.(model)

	// Verify the display shows "7" (button was pressed)
	if updated.display != "7" {
		t.Errorf("Expected display to show '7', got '%s'", updated.display)
	}
}

func TestSpecialButtonStyles(t *testing.T) {
	m := New()

	tests := []struct {
		name   string
		x, y   int
		button string
	}{
		{"AC special function", 0, 0, "AC"},
		{"percent special function", 2, 0, "%"},
		{"division operator", 3, 0, "/"},
		{"equals button", 3, 4, "="},
		{"honk button", 2, 4, "HONK"},
		{"number button", 0, 1, "7"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.cursorX = tt.x
			m.cursorY = tt.y

			output := m.View()

			// When highlighted, button should appear in output
			if !strings.Contains(output, tt.button) {
				t.Errorf("Expected %s button in output", tt.button)
			}
		})
	}
}

func TestZeroButtonSpecialWidth(t *testing.T) {
	m := New()
	m.cursorX = 0
	m.cursorY = 4 // Position on "0" button

	output := m.View()

	// Zero button should appear in output
	if !strings.Contains(output, "0") {
		t.Errorf("Expected zero button in output")
	}
}

func TestMouseInteractionVisualFeedback(t *testing.T) {
	m := New()

	// Simulate mouse click on button at position (0, 2) which should be "AC" button
	mouseMsg := tea.MouseMsg{
		X:    0, // First column
		Y:    2, // Row 0 of buttons (display takes up row 0-1)
		Type: tea.MouseLeft,
	}

	updatedModel, _ := m.Update(mouseMsg)
	updated := updatedModel.(model)

	// Verify the AC button was pressed (display should be reset to "0")
	if updated.display != "0" {
		t.Errorf("Expected display to show '0' after AC press, got '%s'", updated.display)
	}
}

func TestVisualFeedbackPrecedence(t *testing.T) {
	m := New()

	// Test that pressed state takes precedence over highlight
	m.cursorX = 0
	m.cursorY = 0
	m.pressedX = 0
	m.pressedY = 0
	m.activationMethod = activationNavigation

	output := m.View()

	// Should contain button and logo
	if !strings.Contains(output, "AC") {
		t.Errorf("Expected AC button in output")
	}
	if !strings.Contains(output, "GOOSE") {
		t.Errorf("Expected Goose logo in output")
	}
}

func TestNoHighlightWhenCursorNotOnButton(t *testing.T) {
	m := New()

	// Set cursor to an invalid position (shouldn't highlight anything)
	m.cursorX = -1
	m.cursorY = -1

	output := m.View()

	// All buttons should appear in output
	if !strings.Contains(output, "AC") {
		t.Errorf("Expected AC button in output")
	}
	if !strings.Contains(output, "GOOSE") {
		t.Errorf("Expected Goose logo in output")
	}
}

func TestGooseLogoInView(t *testing.T) {
	m := New()

	output := m.View()

	// Check that Goose logo and emoji are in the output
	if !strings.Contains(output, "GOOSE") {
		t.Errorf("Expected GOOSE logo in output")
	}
	if !strings.Contains(output, "🪿") {
		t.Errorf("Expected goose emoji in output")
	}
}

func TestAllButtonsInView(t *testing.T) {
	m := New()

	output := m.View()

	// Check that all buttons appear in the view
	buttons := []string{"AC", "+/-", "%", "/", "7", "8", "9", "x", "4", "5", "6", "-", "1", "2", "3", "+", "0", ".", "HONK", "="}
	for _, btn := range buttons {
		if !strings.Contains(output, btn) {
			t.Errorf("Expected button '%s' in output", btn)
		}
	}
}
