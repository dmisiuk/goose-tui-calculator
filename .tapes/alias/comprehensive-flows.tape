# Comprehensive flows demo tape
# Demonstrates number key input, arrow navigation, error handling, and clearing screen.

Output .tapes/assets/comprehensive-flows.gif
Set FontSize 18
Set Width 1000
Set Height 600

# Ensure calculator binary is built for fast startup
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

########################################
# Flow 1: Number keys – 2 + 3 = 5
########################################
Type "./calc"
Enter
Sleep 2000ms

# Direct number and operator keys
Type "2"
Sleep 300ms
Type "+"
Sleep 300ms
Type "3"
Sleep 300ms
Type "="
Sleep 800ms

# Quit first session
Type "q"
Sleep 600ms

########################################
# Flow 2: Arrow keys + Space – 9 x 4 = 36
########################################
Type "./calc"
Enter
Sleep 2000ms

# Navigate to 9 (row 1 col 2) and press
Down
Sleep 400ms
Down
Sleep 400ms
Right
Sleep 400ms
Space
Sleep 800ms

# Navigate to x operator (row 1 col 3)
Right
Sleep 400ms
Space
Sleep 800ms

# Navigate to 4
Left
Sleep 400ms
Left
Sleep 400ms
Up
Sleep 400ms
Space
Sleep 800ms

# Navigate to =
Down
Sleep 400ms
Right
Sleep 400ms
Space
Sleep 1000ms

# Quit second session
Type "q"
Sleep 600ms

########################################
# Flow 3: Error handling – division by zero
########################################
Type "./calc"
Enter
Sleep 2000ms

Type "5"
Sleep 300ms
Type "/"
Sleep 300ms
Type "0"
Sleep 300ms
Type "="
Sleep 1000ms

# Quit third session
Type "q"
Sleep 600ms

########################################
# Flow 4: Clear screen (AC)
########################################
Type "./calc"
Enter
Sleep 2000ms

# Enter some number then clear
Type "1"
Sleep 300ms
Type "2"
Sleep 300ms
Type "3"
Sleep 300ms
Type "c"   # 'c' maps to AC
Sleep 1000ms

# Quit
Type "q"
Sleep 600ms
