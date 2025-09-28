Output .tapes/assets/feature-keyboard-navigation.gif
Set FontSize 18
Set Width 1000
Set Height 600

# Pre-build binary for stability
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Run calculator
Type "./calc"
Enter
Sleep 2000ms

# Navigate to '2'
Type "{{down}}"
Sleep 150ms
Type "{{down}}"
Sleep 150ms
Type "{{down}}"
Sleep 150ms
Type "{{right}}"
Sleep 150ms
Type " "
Sleep 400ms

# Navigate to '+'
Type "{{right}}"
Sleep 150ms
Type "{{right}}"
Sleep 150ms
Type "{{right}}"
Sleep 150ms
Type " "
Sleep 400ms

# Navigate to '3'
Type "{{left}}"
Sleep 150ms
Type "{{left}}"
Sleep 150ms
Type " "
Sleep 400ms

# Navigate to '='
Type "{{right}}"
Sleep 150ms
Type "{{right}}"
Sleep 150ms
Type " "
Sleep 400ms

# Quit application
Type "q"
Sleep 600ms
