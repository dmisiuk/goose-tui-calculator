Output .tapes/assets/feature-display-alignment.gif
Set FontSize 18
Set Width 1000
Set Height 700

Env TERM xterm-256color
Env COLORTERM truecolor

# Pre-build binary for stability (cold compile can be slow on CI)
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Demonstrate the aligned display with keyboard
Type "./calc"
Enter
Sleep 2000ms

# Show basic calculation to demonstrate the alignment
Type "4"
Sleep 500ms
Type "-"
Sleep 800ms # Pause to show "4 -"
Type "5"
Sleep 800ms # Pause to show "5" with "4 -" above

# Calculate to show result with previous operation
Type "="
Sleep 1500ms # Pause to show the result "4 - 5 = -1" and "-1"

# Show keyboard navigation alignment
Down
Sleep 400ms
Right
Sleep 400ms
Right
Sleep 400ms
Enter # Press 9
Sleep 800ms

# Exit
Type "q"
Sleep 600ms
