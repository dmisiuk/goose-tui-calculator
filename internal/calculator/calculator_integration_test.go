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
