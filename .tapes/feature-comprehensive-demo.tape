Output .tapes/assets/feature-comprehensive-demo.gif

Set FontSize 18
Set Width 1200
Set Height 1000

Env TERM xterm-256color
Env COLORTERM truecolor

# Pre-build binary for stability
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Launch calculator
Type "./calc"
Enter
Sleep 2500ms

# ====================================
# Flow 1: Direct keyboard input (number keys)
# ====================================
# Simple addition: 15 + 27 = 42
Type "1"
Sleep 400ms
Type "5"
Sleep 400ms
Type "+"
Sleep 600ms
Type "2"
Sleep 400ms
Type "7"
Sleep 400ms
Type "="
Sleep 1200ms

# ====================================
# Flow 2: Clean screen with AC button
# ====================================
Type "c"
Sleep 800ms

# ====================================
# Flow 3: Arrow key navigation + Enter
# ====================================
# Navigate to 8 button (row 0, col 2)
Down
Sleep 400ms
Right
Sleep 400ms
Right
Sleep 400ms
# Press 8
Enter
Sleep 500ms

# Navigate to * button (row 0, col 3)
Right
Sleep 400ms
# Press *
Enter
Sleep 500ms

# Navigate to 6 button (row 1, col 2)
Down
Sleep 400ms
Left
Sleep 400ms
# Press 6
Enter
Sleep 500ms

# Navigate to = button (row 4, col 2)
Down
Sleep 400ms
Down
Sleep 400ms
Down
Sleep 400ms
# Press =
Enter
Sleep 1200ms

# ====================================
# Flow 4: Clean screen again
# ====================================
# Navigate to AC button (row 0, col 0)
Up
Sleep 400ms
Up
Sleep 400ms
Up
Sleep 400ms
Left
Sleep 400ms
Left
Sleep 400ms
# Press AC
Enter
Sleep 800ms

# ====================================
# Flow 5: Show error (division by zero)
# ====================================
# Type 9 / 0 = to trigger error
Type "9"
Sleep 400ms
Type "/"
Sleep 600ms
Type "0"
Sleep 400ms
Type "="
Sleep 1500ms

# ====================================
# Flow 6: Clear error with AC
# ====================================
Type "c"
Sleep 800ms

# ====================================
# Flow 7: Mixed input methods
# ====================================
# Start with keyboard
Type "3"
Sleep 400ms
Type "+"
Sleep 600ms

# Switch to navigation for second operand
# Navigate to 7 button (row 0, col 1)
Down
Sleep 400ms
Right
Sleep 400ms
Enter
Sleep 500ms

# Back to keyboard for equals
Type "="
Sleep 1200ms

# Clean exit
Type "q"
Sleep 600ms
