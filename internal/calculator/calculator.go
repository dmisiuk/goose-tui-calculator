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
	bgColor         = lipgloss.Color("#1c1c1c")
	displayColor    = lipgloss.Color("#2E8B57") // SeaGreen
	resultColor     = lipgloss.Color("#32CD32") // LimeGreen
	exprColor       = lipgloss.Color("#708090") // SlateGray
	titleColor      = lipgloss.Color("#FFD700") // Gold
	borderColor     = lipgloss.Color("#8A2BE2") // BlueViolet
	buttonTextColor = lipgloss.Color("#FFFFFF")

	// Button colors
	acColor       = lipgloss.Color("#FF4500") // OrangeRed
	numColor      = lipgloss.Color("#6495ED") // CornflowerBlue
	opColor       = lipgloss.Color("#FFA500") // Orange
	funcColor     = lipgloss.Color("#B0C4DE") // LightSteelBlue
	equalsColor   = lipgloss.Color("#FFD700") // Gold
	honkColor     = lipgloss.Color("#FF69B4") // HotPink
	zeroColor     = numColor
	inactiveColor = lipgloss.Color("#444444")

	// Logo
	logoStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(titleColor).
			Align(lipgloss.Center)

	// Display
	displayContainerStyle = lipgloss.NewStyle().
				Background(displayColor).
				Border(lipgloss.ThickBorder(), false, false, false, true).
				BorderTop(false).
				BorderBottom(true).
				BorderLeft(false).
				BorderRight(false).
				BorderForeground(borderColor).
				Padding(1, 2)

	displayStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(resultColor).
			Align(lipgloss.Right)

	previousDisplayStyle = lipgloss.NewStyle().
				Foreground(exprColor).
				Align(lipgloss.Right)

	// Buttons
	baseButtonStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(buttonTextColor).
			Align(lipgloss.Center).
			Width(7).
			Height(2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(inactiveColor)

	acButtonStyle         = baseButtonStyle.Copy().Background(acColor)
	numberButtonStyle     = baseButtonStyle.Copy().Background(numColor)
	operatorButtonStyle   = baseButtonStyle.Copy().Background(opColor)
	functionalButtonStyle = baseButtonStyle.Copy().Background(funcColor).Foreground(lipgloss.Color("#000000"))
	equalsButtonStyle     = baseButtonStyle.Copy().Background(equalsColor).Foreground(lipgloss.Color("#000000"))
	honkButtonStyle       = baseButtonStyle.Copy().Background(honkColor)
	zeroButtonStyle       = baseButtonStyle.Copy().Background(zeroColor)

	// Visual feedback
	highlightStyle = baseButtonStyle.Copy().
			Border(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("#FFFFFF"))

	pressedStyle = baseButtonStyle.Copy().
			Background(lipgloss.Color("#FFFFFF")).
			Foreground(lipgloss.Color("#000000"))

	directKeyboardStyle = baseButtonStyle.Copy().
				Background(lipgloss.Color("#FF00FF")) // Magenta for direct key press

	// Calculator body
	calculatorBodyStyle = lipgloss.NewStyle().
				Border(lipgloss.DoubleBorder()).
				BorderForeground(borderColor).
				Padding(1).
				Background(bgColor)
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
			{"AC", "+/-", "%", "÷"},
			{"7", "8", "9", "✖"},
			{"4", "5", "6", "−"},
			{"1", "2", "3", "➕"},
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
					// NOTE: This layout calculation is simplified. A more robust TUI would
					// not have hardcoded layout logic in the update function.
					const buttonGridYOffset = 9
					const buttonWidth = 7
					const buttonSpacing = 1
					gridX := x * (buttonWidth + buttonSpacing)
					gridY := y + buttonGridYOffset

					// Clicks are registered on the top line of the 2-line-height buttons
					if msg.Y == gridY && msg.X >= gridX && msg.X < gridX+buttonWidth {
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
			case "➕":
				result = val1 + val2
			case "−":
				result = val1 - val2
			case "✖":
				result = val1 * val2
			case "÷":
				if val2 == 0 {
					m.display = "Error"
					m.isError = true
					break
				}
				result = val1 / val2
			}
			if !m.isError {
				m.previousDisplay = fmt.Sprintf("%s %s %s", m.operand1, m.operator, operand2)
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

func isOperator(s string) bool { return s == "➕" || s == "−" || s == "✖" || s == "÷" }

func mapKeyToButton(k string) (string, bool) {
	if isNumber(k) {
		return k, true
	}
	switch k {
	case "+":
		return "➕", true
	case "-":
		return "−", true
	case "/":
		return "÷", true
	case "*", "x":
		return "✖", true
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
	case "h", "H":
		return "HONK", true
	}
	return "", false
}

func (m model) View() string {
	if m.isQuitting {
		return "Thanks for using the Goose Calculator!\n"
	}

	var b strings.Builder
	totalWidth := 31 // 4 buttons * 7 width + 3 spaces

	// Title
	b.WriteString(logoStyle.Width(totalWidth).Render("🦢  GOOSE CALC  🦢"))
	b.WriteString("\n")

	// Display
	isFinished := m.lastButton == "=" || m.lastButton == "HONK"

	exprLine := "Expr: " + m.previousDisplay
	resLine := "Result: " + m.display

	if isFinished {
		resLine += " ✅"
	} else if m.isOperand2 {
		resLine = "Result: "
	}

	if m.previousDisplay == "" {
		exprLine = "Expr: "
	}

	separator := lipgloss.NewStyle().Foreground(borderColor).Render(strings.Repeat("─", totalWidth-4))
	displayContent := lipgloss.JoinVertical(lipgloss.Left,
		previousDisplayStyle.Render(exprLine),
		separator,
		displayStyle.Render(resLine),
	)
	b.WriteString(displayContainerStyle.Width(totalWidth).Render(displayContent))
	b.WriteString("\n\n")

	// Button grid
	for y, row := range m.buttons {
		var rowButtons []string
		for x, val := range row {
			var style lipgloss.Style

			if val == "AC" {
				style = acButtonStyle
			} else if val == "=" {
				style = equalsButtonStyle
			} else if val == "HONK" {
				style = honkButtonStyle
			} else if isOperator(val) {
				style = operatorButtonStyle
			} else if val == "+/-" || val == "%" || val == "." {
				style = functionalButtonStyle
			} else if isNumber(val) {
				style = numberButtonStyle
			} else {
				style = baseButtonStyle // Fallback
			}

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

			rowButtons = append(rowButtons, style.Render(val))
		}
		b.WriteString(lipgloss.JoinHorizontal(lipgloss.Center, strings.Join(rowButtons, " ")))
		b.WriteString("\n")
	}

	// Help
	b.WriteString("\n")
	helpText := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#95A5A6")).
		Width(totalWidth).
		Align(lipgloss.Center).
		Render("(h for help · q/esc to quit)")
	b.WriteString(helpText)

	return calculatorBodyStyle.Render(b.String())
}

func isSpecialFunc(s string) bool { return s == "AC" || s == "+/-" || s == "%" }
