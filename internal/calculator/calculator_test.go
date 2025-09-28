package calculator

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func TestButtonHighlightState(t *testing.T) {
	m := New()

	// Position cursor on first button (AC)
	m.cursorX = 0
	m.cursorY = 0

	output := m.View()

	// Check that highlight style is applied to the AC button
	highlightMarker := highlightStyle.Render("AC")
	if !strings.Contains(output, highlightMarker) {
		t.Errorf("Expected output to contain highlighted AC button, got: %s", output)
	}

	// Verify other buttons are not highlighted
	normalMarker := buttonStyle.Render("+/-")
	if !strings.Contains(output, normalMarker) {
		t.Errorf("Expected other buttons to have normal styling")
	}
}

func TestButtonPressedState(t *testing.T) {
	m := New()

	// Position cursor and set pressed state on same button
	m.cursorX = 1
	m.cursorY = 0
	m.pressedX = 1
	m.pressedY = 0

	output := m.View()

	// Pressed style should take precedence over highlight
	pressedMarker := pressedStyle.Render("+/-")
	if !strings.Contains(output, pressedMarker) {
		t.Errorf("Expected output to contain pressed +/- button")
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
		name     string
		x, y     int
		button   string
		expected lipgloss.Style
	}{
		{"AC special function", 0, 0, "AC", specialFuncsStyle},
		{"percent special function", 2, 0, "%", specialFuncsStyle},
		{"division operator", 3, 0, "/", operatorStyle},
		{"equals button", 2, 4, "=", equalsStyle},
		{"number button", 0, 1, "7", buttonStyle},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.cursorX = tt.x
			m.cursorY = tt.y

			output := m.View()

			// When highlighted, should use highlight style instead
			highlightMarker := highlightStyle.Render(tt.button)
			if !strings.Contains(output, highlightMarker) {
				t.Errorf("Expected highlighted %s button in output", tt.button)
			}
		})
	}
}

func TestZeroButtonSpecialWidth(t *testing.T) {
	m := New()
	m.cursorX = 0
	m.cursorY = 4 // Position on "0" button

	output := m.View()

	// Zero button should have special width (11 instead of 5)
	zeroHighlight := highlightStyle.Copy().Width(11).Render("0")
	if !strings.Contains(output, zeroHighlight) {
		t.Errorf("Expected zero button to have special width highlighting")
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

	output := m.View()

	// Should contain pressed style, not highlight style
	pressedMarker := pressedStyle.Render("AC")
	highlightMarker := highlightStyle.Render("AC")

	if !strings.Contains(output, pressedMarker) {
		t.Errorf("Expected pressed style to be applied")
	}

	// Should not contain just the highlight style when pressed
	if strings.Contains(output, highlightMarker) && !strings.Contains(output, pressedMarker) {
		t.Errorf("Pressed style should take precedence over highlight style")
	}
}

func TestNoHighlightWhenCursorNotOnButton(t *testing.T) {
	m := New()

	// Set cursor to an invalid position (shouldn't highlight anything)
	m.cursorX = -1
	m.cursorY = -1

	output := m.View()

	// All buttons should use their normal styles
	normalAC := specialFuncsStyle.Render("AC")
	if !strings.Contains(output, normalAC) {
		t.Errorf("Expected AC button to have normal special function style when not highlighted")
	}
}
