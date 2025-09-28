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
Sleep 400ms
Type "{{down}}"
Sleep 400ms
Type "{{down}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type " "
Sleep 1500ms

# Navigate to '+'
Type "{{right}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type " "
Sleep 1500ms

# Navigate to '3'
Type "{{left}}"
Sleep 400ms
Type "{{left}}"
Sleep 400ms
Type " "
Sleep 1500ms

# Navigate to '='
Type "{{right}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type " "
Sleep 1500ms

# Quit application
Type "q"
Sleep 1000ms
