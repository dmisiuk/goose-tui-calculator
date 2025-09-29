package calculator

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	displayColor = lipgloss.Color("#3C4043") // A dark gray for the display background
	buttonColor  = lipgloss.Color("#BDC1C6") // A light gray for buttons
	specialFuncs = lipgloss.Color("#F28B82") // A reddish color for AC, +/-, %
	operators    = lipgloss.Color("#F9AB00") // An orange for operators
	equalsColor  = lipgloss.Color("#81C995") // A greenish color for equals
	textColor    = lipgloss.Color("#FFFFFF") // White text
	altTextColor = lipgloss.Color("#9E9E9E") // A dimmer white text

	// Styles
	containerStyle = lipgloss.NewStyle().
			Background(displayColor).
			Width(20).
			Height(4).
			Padding(1, 2)

	displayStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(textColor).
			Align(lipgloss.Right).
			Width(16)

	previousDisplayStyle = lipgloss.NewStyle().
				Foreground(altTextColor).
				Align(lipgloss.Right).
				Width(16)

	buttonStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(textColor).
			Background(buttonColor).
			Align(lipgloss.Center).
			Width(5).
			Height(1)

	specialFuncsStyle = buttonStyle.Copy().Background(specialFuncs)
	operatorStyle     = buttonStyle.Copy().Background(operators)
	equalsStyle       = buttonStyle.Copy().Background(equalsColor)

	// Highlight and pressed button styles
	highlightBackground      = lipgloss.Color("#FFD700") // Gold color for highlight
	pressedBackground        = lipgloss.Color("#FF4500") // OrangeRed for pressed
	directKeyboardBackground = lipgloss.Color("#6A5ACD") // SlateBlue for direct keyboard input

	highlightStyle      = buttonStyle.Copy().Background(highlightBackground).Foreground(lipgloss.Color("#000000"))
	pressedStyle        = buttonStyle.Copy().Background(pressedBackground).Foreground(lipgloss.Color("#FFFFFF"))
	directKeyboardStyle = buttonStyle.Copy().Background(directKeyboardBackground).Foreground(lipgloss.Color("#FFFFFF"))

	// Calculator border style
	calculatorBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#9E9E9E")).
				Padding(1)
)

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*300, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type activationMethod int

const (
	activationNone activationMethod = iota
	activationNavigation
	activationDirectKeyboard
)

type model struct {
	display             string
	previousDisplay     string
	buttons             [][]string
	cursorX             int
	cursorY             int
	operator            string
	operand1            string
	isOperand2          bool
	lastButton          string
	isError             bool
	keys                keyMap
	mouseEvent          tea.MouseMsg
	isQuitting          bool
	lastOperator        string
	pressedX            int
	pressedY            int
	activationMethod    activationMethod
	activationStartTime time.Time
}

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Enter key.Binding
	Quit  key.Binding
	Esc   key.Binding
}

var defaultKeyMap = keyMap{
	Up:    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "move up")),
	Down:  key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "move down")),
	Left:  key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("←/h", "move left")),
	Right: key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("→/l", "move right")),
	Enter: key.NewBinding(key.WithKeys("enter", " "), key.WithHelp("enter", "press button")),
	Quit:  key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
	Esc:   key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "quit")),
}

func New() model {
	return model{
		display:         "0",
		previousDisplay: "",
		buttons: [][]string{
			{"AC", "+/-", "%", "/"},
			{"7", "8", "9", "x"},
			{"4", "5", "6", "-"},
			{"1", "2", "3", "+"},
			{"0", ".", "="},
		},
		keys: defaultKeyMap,
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		// Clear activation visual feedback after timing expires
		if m.activationMethod != activationNone && time.Since(m.activationStartTime) > time.Millisecond*300 {
			m.activationMethod = activationNone
			m.pressedX = -1
			m.pressedY = -1
		}
		return m, tick()
	case tea.KeyMsg:
		// Direct mapped keys (digits/operators/etc.) first
		if btn, ok := mapKeyToButton(msg.String()); ok {
			// Find button position for visual feedback
			for y, row := range m.buttons {
				for x, val := range row {
					if val == btn {
						m.pressedX = x
						m.pressedY = y
						m.activationMethod = activationDirectKeyboard
						m.activationStartTime = time.Now()
						break
					}
				}
			}
			updatedModel, cmd := m.handleButtonPress(btn)
			if updated, ok := updatedModel.(model); ok {
				return updated, tea.Batch(cmd, tick())
			}
			return updatedModel, tea.Batch(cmd, tick())
		}
		switch {
		case key.Matches(msg, m.keys.Quit), key.Matches(msg, m.keys.Esc):
			m.isQuitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Up):
			if m.cursorY > 0 {
				m.cursorY--
				if m.cursorX >= len(m.buttons[m.cursorY]) {
					m.cursorX = len(m.buttons[m.cursorY]) - 1
				}
			}
		case key.Matches(msg, m.keys.Down):
			if m.cursorY < len(m.buttons)-1 {
				m.cursorY++
				if m.cursorX >= len(m.buttons[m.cursorY]) {
					m.cursorX = len(m.buttons[m.cursorY]) - 1
				}
			}
		case key.Matches(msg, m.keys.Left):
			if m.cursorX > 0 {
				m.cursorX--
			}
		case key.Matches(msg, m.keys.Right):
			if m.cursorX < len(m.buttons[m.cursorY])-1 {
				m.cursorX++
			}
		case key.Matches(msg, m.keys.Enter):
			button := m.buttons[m.cursorY][m.cursorX]
			m.pressedX = m.cursorX
			m.pressedY = m.cursorY
			m.activationMethod = activationNavigation
			m.activationStartTime = time.Now()
			updatedModel, cmd := m.handleButtonPress(button)
			if updated, ok := updatedModel.(model); ok {
				return updated, tea.Batch(cmd, tick())
			}
			return updatedModel, tea.Batch(cmd, tick())
		}
	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			for y, row := range m.buttons {
				for x, val := range row {
					if msg.Y == y+2 && msg.X >= x*6 && msg.X < x*6+5 {
						m.pressedX = x
						m.pressedY = y
						m.activationMethod = activationNavigation // Mouse clicks use navigation style
						m.activationStartTime = time.Now()
						updatedModel, cmd := m.handleButtonPress(val)
						if updated, ok := updatedModel.(model); ok {
							return updated, tea.Batch(cmd, tick())
						}
						return updatedModel, tea.Batch(cmd, tick())
					}
				}
			}
		}
	}
	return m, nil
}
func (m model) HandleButtonPress(button string) (model, tea.Cmd) {
	updatedModel, cmd := m.handleButtonPress(button)
	if updated, ok := updatedModel.(model); ok {
		return updated, cmd
	}
	return m, cmd
}

func (m model) Display() string {
	return m.display
}

func (m model) handleButtonPress(button string) (tea.Model, tea.Cmd) {
	m.lastButton = button
	m.isError = false

	switch {
	case isNumber(button):
		if m.isOperand2 {
			if m.operator == "" {
				m.previousDisplay = ""
			}
			m.display = button
			m.isOperand2 = false
		} else if m.display == "0" {
			m.display = button
		} else {
			m.display += button
		}
	case button == ".":
		if !strings.Contains(m.display, ".") {
			m.display += "."
		}
	case isOperator(button):
		m.operand1 = m.display
		m.operator = button
		m.isOperand2 = true
		m.previousDisplay = m.operand1 + " " + m.operator
	case button == "AC":
		m.display = "0"
		m.previousDisplay = ""
		m.operand1 = ""
		m.operator = ""
		m.isOperand2 = false
	case button == "+/-":
		if m.display != "0" {
			if strings.HasPrefix(m.display, "-") {
				m.display = strings.TrimPrefix(m.display, "-")
			} else {
				m.display = "-" + m.display
			}
		}
	case button == "%":
		val, _ := strconv.ParseFloat(m.display, 64)
		m.display = fmt.Sprintf("%g", val/100)
	case button == "=":
		if m.operand1 != "" && m.operator != "" {
			operand2 := m.display
			val1, err1 := strconv.ParseFloat(m.operand1, 64)
			val2, err2 := strconv.ParseFloat(operand2, 64)
			if err1 != nil || err2 != nil {
				m.display = "Error"
				m.isError = true
				break
			}
			var result float64
			switch m.operator {
			case "+":
				result = val1 + val2
			case "-":
				result = val1 - val2
			case "x":
				result = val1 * val2
			case "/":
				if val2 == 0 {
					m.display = "Error"
					m.isError = true
					break
				}
				result = val1 / val2
			}
			if !m.isError {
				m.previousDisplay = fmt.Sprintf("%s %s %s = %g", m.operand1, m.operator, operand2, result)
				m.display = fmt.Sprintf("%g", result)
			}
			m.operand1 = ""
			m.operator = ""
			m.isOperand2 = true
		}
	}

	return m, func() tea.Msg { fmt.Print("\a"); return nil }
}

func isNumber(s string) bool { _, err := strconv.Atoi(s); return err == nil }

func isOperator(s string) bool { return s == "+" || s == "-" || s == "x" || s == "/" }

// mapKeyToButton allows direct keyboard entry of calculator buttons without navigation.
// Supports: digits, + - * / x . = c (AC), %, and ~ as sign toggle.
func mapKeyToButton(k string) (string, bool) {
	if isNumber(k) {
		return k, true
	}
	switch k {
	case "+", "-", "/":
		return k, true
	case "*", "x":
		return "x", true
	case ".":
		return ".", true
	case "=":
		return "=", true
	case "c", "C":
		return "AC", true
	case "%":
		return "%", true
	case "~":
		return "+/-", true
	}
	return "", false
}

func (m model) View() string {
	if m.isQuitting {
		return "Thanks for using the Goose Calculator!\n"
	}
	var b strings.Builder

	// Render previous and current display
	var combinedDisplay string
	if m.previousDisplay != "" {
		previousDisplayStr := previousDisplayStyle.Render(m.previousDisplay)
		displayStr := displayStyle.Render(m.display)
		combinedDisplay = lipgloss.JoinVertical(lipgloss.Right, previousDisplayStr, displayStr)
	} else {
		// To maintain alignment and position, we create an empty line with the same style
		// and join it with the main display. This ensures the main display stays at the bottom.
		emptyLine := previousDisplayStyle.Render("")
		displayStr := displayStyle.Render(m.display)
		combinedDisplay = lipgloss.JoinVertical(lipgloss.Right, emptyLine, displayStr)
	}
	b.WriteString(containerStyle.Render(combinedDisplay))

	b.WriteString("\n\n")
	for y, row := range m.buttons {
		var rowStr []string
		for x, val := range row {
			style := buttonStyle
			if isSpecialFunc(val) {
				style = specialFuncsStyle
			} else if isOperator(val) {
				style = operatorStyle
			} else if val == "=" {
				style = equalsStyle
			}
			// Handle visual feedback based on activation method and cursor position
			if m.pressedX == x && m.pressedY == y {
				// Button is currently being pressed - use activation-specific style
				switch m.activationMethod {
				case activationDirectKeyboard:
					style = directKeyboardStyle
				case activationNavigation:
					style = pressedStyle
				}
			} else if m.cursorY == y && m.cursorX == x {
				// Button is highlighted by cursor navigation
				style = highlightStyle
			}
			if val == "0" {
				style = style.Copy().Width(10)
				rowStr = append(rowStr, style.Render(val))
				if x+1 < len(row) {
					x++
				}
			} else {
				rowStr = append(rowStr, style.Render(val))
			}
		}
		b.WriteString(lipgloss.JoinHorizontal(lipgloss.Left, rowStr...))
		b.WriteString("\n")
	}
	b.WriteString("\nType numbers/operators or use arrows + space/enter.\nPress q to quit.")

	// Wrap the entire calculator in a border
	return calculatorBorderStyle.Render(b.String())
}

func isSpecialFunc(s string) bool { return s == "AC" || s == "+/-" || s == "%" }
