Output .tapes/assets/comprehensive-flow.gif
Set Theme "Cyberdyne"
Set FontSize 18
Set Width 1000
Set Height 800

# Pre-build binary for stability (cold compile can be slow on CI)
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Start the calculator
Type "./calc"
Enter
Sleep 2000ms

# Flow 1: Calculation using direct number key input (10 + 5 = 15)
Type "10+5="
Sleep 1500ms

# Flow 2: Clearing the screen using the AC button (mapped to 'c' or 'C')
Type "c"
Sleep 1000ms

# Flow 3: Calculation using arrow key navigation and spacebar (8 * 2 = 16)
# From AC (0,0) to 8 (1,1)
Down
Sleep 300ms
Right
Sleep 300ms
Space
Sleep 500ms
# From 8 (1,1) to x (1,3)
Right
Sleep 300ms
Right
Sleep 300ms
Space
Sleep 500ms
# From x (1,3) to 2 (3,1)
Down
Sleep 300ms
Down
Sleep 300ms
Left
Sleep 300ms
Left
Sleep 300ms
Space
Sleep 500ms
# From 2 (3,1) to = (4,2)
Down
Sleep 300ms
Right
Sleep 300ms
Space
Sleep 1500ms

# Flow 4: An operation that results in an error (9 / 0)
# Clear screen first
Type "c"
Sleep 1000ms
# Enter 9 / 0 =
Type "9/0="
Sleep 1500ms

# Flow 5: Clearing the error message with the AC button
Type "c"
Sleep 1000ms

# Quit the application
Type "q"
Sleep 500ms