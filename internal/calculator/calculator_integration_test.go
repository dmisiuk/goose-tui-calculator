package calculator_test

import (
	"testing"

	"github.com/dmisiuk/goose-tui-calculator/internal/calculator"
)

func TestIntegrationExample(t *testing.T) {
	m := calculator.New()
	m, _ = m.HandleButtonPress("1")
	m, _ = m.HandleButtonPress("➕")
	m, _ = m.HandleButtonPress("2")
	m, _ = m.HandleButtonPress("=")
	if m.Display() != "3" {
		t.Errorf("Expected display 3, got %s", m.Display())
	}
}

func TestHonkButtonIntegration(t *testing.T) {
	m := calculator.New()
	m, _ = m.HandleButtonPress("7")
	m, _ = m.HandleButtonPress("✖")
	m, _ = m.HandleButtonPress("3")
	m, _ = m.HandleButtonPress("HONK")
	if m.Display() != "21" {
		t.Errorf("Expected display 21 after HONK, got %s", m.Display())
	}
}