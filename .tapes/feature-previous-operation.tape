Output .tapes/assets/feature-previous-operation.gif
Set FontSize 18
Set Width 1000
Set Height 700

Env TERM xterm-256color
Env COLORTERM truecolor

# Pre-build binary for stability (cold compile can be slow on CI)
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Demonstrate the previous operation feature
Type "./calc"
Enter
Sleep 2000ms

# Type "2 + 3"
Type "2"
Sleep 500ms
Type "+"
Sleep 800ms # Pause to show "2 +"
Type "3"
Sleep 800ms # Pause to show "3" with "2 +" above

# Calculate
Type "="
Sleep 1500ms # Pause to show the result "2 + 3 = 5" and "5"

# Start a new calculation
Type "7"
Sleep 1000ms # Pause to show the previous operation is cleared

# Exit
Type "q"
Sleep 600ms