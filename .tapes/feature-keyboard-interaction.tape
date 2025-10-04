Output .tapes/assets/feature-keyboard-interaction.gif
Set FontSize 18
Set Width 1000
Set Height 700

Env TERM xterm-256color
Env COLORTERM truecolor

# Prepare and build
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Run the calculator
Type "./calc"
Enter
Sleep 2000ms

# Input sequence by keyboard keys:

# Number keys '1', '0'
Type "1"
Sleep 200ms
Type "0"
Sleep 200ms

# Operator '+'
Type "+"
Sleep 200ms

# Number key '5'
Type "5"
Sleep 200ms

# Press space (simulate 'Enter' alternative)
Type " "
Sleep 200ms

# Press '=' explicitly with Enter
Type "="
Sleep 500ms

# Clear with mapped 'c' key for AC
Type "c"
Sleep 400ms

# Multiplication '*' mapped to 'x'
Type "*"
Sleep 200ms
Type "3"
Sleep 200ms

# Equals by keyboard Enter key
Enter
Sleep 800ms

# Quit application
Type "q"
Sleep 600ms
