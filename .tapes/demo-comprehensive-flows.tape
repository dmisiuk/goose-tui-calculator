Output .tapes/assets/demo-comprehensive-flows.gif

Set FontSize 18
Set Width 1200
Set Height 700

Set Theme "Molokai"
Set TypingSpeed 100ms

# Pre-build binary for stability
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Start comprehensive demo
Type "./calc"
Enter
Sleep 2000ms

# Part 1: Number keys usage demo
# Type "123 + 456 =" using direct number keys
Type "1"
Sleep 300ms
Type "2"
Sleep 300ms
Type "3"
Sleep 400ms

Type "+"
Sleep 400ms

Type "4"
Sleep 300ms
Type "5"
Sleep 300ms
Type "6"
Sleep 400ms

Type "="
Sleep 1000ms

# Clear for next demo
Type "c"
Sleep 500ms

# Part 2: Arrow keys + space usage demo
# Navigate to "9" using arrow keys and press space
Down
Sleep 200ms
Down
Sleep 200ms
Right
Sleep 200ms
Right
Sleep 200ms
Space
Sleep 400ms

# Navigate to "x" (multiply)
Right
Sleep 200ms
Right
Sleep 200ms
Space
Sleep 400ms

# Navigate to "7"
Up
Sleep 200ms
Left
Sleep 200ms
Left
Sleep 200ms
Space
Sleep 400ms

# Navigate to "=" and get result
Down
Sleep 200ms
Down
Sleep 200ms
Right
Sleep 200ms
Right
Sleep 200ms
Space
Sleep 1000ms

# Clear for next demo
Type "c"
Sleep 500ms

# Part 3: Error handling demo
# Try to divide by zero
Type "5"
Sleep 300ms
Type "/"
Sleep 300ms
Type "0"
Sleep 300ms
Type "="
Sleep 1000ms

# Show error state, then clear
Type "c"
Sleep 500ms

# Try another invalid operation (very large number)
Type "999999999"
Sleep 600ms
Type "*"
Sleep 300ms
Type "999999999"
Sleep 600ms
Type "="
Sleep 1000ms

# Clear error
Type "c"
Sleep 500ms

# Part 4: Screen cleaning demo
# Do a complex calculation to fill the screen
Type "123.45"
Sleep 600ms
Type "+"
Sleep 300ms
Type "67.89"
Sleep 600ms
Type "*"
Sleep 300ms
Type "2"
Sleep 300ms
Type "="
Sleep 1000ms

# Clear screen with AC
Type "c"
Sleep 500ms

# Show another calculation
Type "25"
Sleep 300ms
Type "%"
Sleep 300ms
Type "="
Sleep 1000ms

# Final clear
Type "c"
Sleep 500ms

# Quit demo
Type "q"
Sleep 600ms