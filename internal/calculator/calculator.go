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
	// Colors - using terminal defaults for backgrounds
	displayBackground = lipgloss.Color("#1B5E4F") // Dark green LCD
	displayText       = lipgloss.Color("#D5F5E3") // Light green LCD text
	displayTextDim    = lipgloss.Color("#82C9B5") // Dim green

	acButtonColor      = lipgloss.Color("#C0392B") // Red AC
	numberButtonColor  = lipgloss.Color("#5D6D7E") // Gray numbers
	operatorColor      = lipgloss.Color("#D68910") // Orange operators
	functionalKeyColor = lipgloss.Color("#7F8C8D") // Light gray functional keys (+/-, %, .)
	zeroButtonColor    = lipgloss.Color("#34495E") // Darker blue-gray for 0
	equalsButtonColor  = lipgloss.Color("#E67E22") // Bright orange for equals

	buttonTextColor = lipgloss.Color("#FFFFFF") // White text
	logoTextColor   = lipgloss.Color("#FFFFFF") // White logo

	// Logo
	logoStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(logoTextColor).
			Align(lipgloss.Center)

	// LCD Display
	displayContainerStyle = lipgloss.NewStyle().
				Background(displayBackground).
				Padding(1, 2)

	displayStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(displayText).
			Align(lipgloss.Right)

	previousDisplayStyle = lipgloss.NewStyle().
				Foreground(displayTextDim).
				Align(lipgloss.Right)

	// Buttons - wider for better look
	baseButtonStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(buttonTextColor).
			Align(lipgloss.Center).
			Width(6).
			Height(2)

	numberButtonStyle     = baseButtonStyle.Copy().Background(numberButtonColor)
	acButtonStyle         = baseButtonStyle.Copy().Background(acButtonColor)
	operatorButtonStyle   = baseButtonStyle.Copy().Background(operatorColor)
	functionalButtonStyle = baseButtonStyle.Copy().Background(functionalKeyColor)
	zeroButtonStyle       = baseButtonStyle.Copy().Background(zeroButtonColor)
	equalsButtonStyle     = baseButtonStyle.Copy().Background(equalsButtonColor)

	// Visual feedback
	highlightBackground      = lipgloss.Color("#FFD700")
	pressedBackground        = lipgloss.Color("#FF4500")
	directKeyboardBackground = lipgloss.Color("#6A5ACD")

	highlightStyle      = baseButtonStyle.Copy().Background(highlightBackground).Foreground(lipgloss.Color("#000000"))
	pressedStyle        = baseButtonStyle.Copy().Background(pressedBackground).Foreground(buttonTextColor)
	directKeyboardStyle = baseButtonStyle.Copy().Background(directKeyboardBackground).Foreground(buttonTextColor)

	// Calculator body - NO background, use terminal default
	calculatorBodyStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#95A5A6")).
				Padding(1, 2)
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
			{"0", ".", "HONK", "="},
		},
		keys: defaultKeyMap,
	}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.activationMethod != activationNone && time.Since(m.activationStartTime) > time.Millisecond*300 {
			m.activationMethod = activationNone
			m.pressedX = -1
			m.pressedY = -1
		}
		return m, tick()
	case tea.KeyMsg:
		if btn, ok := mapKeyToButton(msg.String()); ok {
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
						m.activationMethod = activationNavigation
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
	case button == "=" || button == "HONK":
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

	// Logo - match button grid width (4 buttons × 6 chars = 24)
	b.WriteString(logoStyle.Width(24).Render("🪿 GOOSE 🪿"))
	b.WriteString("\n\n")

	// Display - width matches 4 buttons at 6 chars each = 24
	displayWidth := 24
	var combinedDisplay string
	if m.previousDisplay != "" {
		prev := previousDisplayStyle.Width(displayWidth - 4).Render(m.previousDisplay)
		curr := displayStyle.Width(displayWidth - 4).Render(m.display)
		combinedDisplay = lipgloss.JoinVertical(lipgloss.Right, prev, curr)
	} else {
		empty := previousDisplayStyle.Width(displayWidth - 4).Render("")
		curr := displayStyle.Width(displayWidth - 4).Render(m.display)
		combinedDisplay = lipgloss.JoinVertical(lipgloss.Right, empty, curr)
	}
	b.WriteString(displayContainerStyle.Width(displayWidth).Render(combinedDisplay))
	b.WriteString("\n\n")

	// Button grid
	for y, row := range m.buttons {
		var rowButtons []string
		for x, val := range row {
			var style lipgloss.Style

			if val == "AC" {
				style = acButtonStyle
			} else if val == "=" || val == "HONK" {
				style = equalsButtonStyle
			} else if isOperator(val) {
				style = operatorButtonStyle
			} else if val == "+/-" || val == "%" || val == "." {
				style = functionalButtonStyle
			} else if val == "0" {
				style = zeroButtonStyle
			} else {
				style = numberButtonStyle
			}

			// Apply feedback
			if m.pressedX == x && m.pressedY == y {
				switch m.activationMethod {
				case activationDirectKeyboard:
					style = directKeyboardStyle.Copy().Width(style.GetWidth())
				case activationNavigation:
					style = pressedStyle.Copy().Width(style.GetWidth())
				}
			} else if m.cursorY == y && m.cursorX == x {
				style = highlightStyle.Copy().Width(style.GetWidth())
			}

			// Previous design widened 0 when row had 3 columns; with 4-column layout keep standard width.
			rowButtons = append(rowButtons, style.Render(val))
		}
		b.WriteString(lipgloss.JoinHorizontal(lipgloss.Left, rowButtons...))
		b.WriteString("\n")
	}

	// Help - centered to match button grid width
	b.WriteString("\n")
	helpText := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#95A5A6")).
		Width(24).
		Align(lipgloss.Center).
		Render("Press q or esc to quit")
	b.WriteString(helpText)

	return calculatorBodyStyle.Render(b.String())
}

func isSpecialFunc(s string) bool { return s == "AC" || s == "+/-" || s == "%" || s == "HONK" }
