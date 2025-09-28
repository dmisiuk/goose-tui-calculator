Output .tapes/assets/feature-enhanced-visual-feedback.gif

Set FontSize 18
Set Width 1000
Set Height 600

# Pre-build binary for stability
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Demo showcasing enhanced visual feedback for different input methods
Type "./calc"
Enter
Sleep 2000ms

# Part 1: Direct keyboard input (blue/purple feedback)
# Type "2" directly - should show blue feedback
Type "2"
Sleep 500ms
# Type "+" directly - should show blue feedback
Type "+"
Sleep 500ms
# Type "3" directly - should show blue feedback
Type "3"
Sleep 500ms
# Type "=" directly - should show blue feedback
Type "="
Sleep 800ms

# Clear for next demo
Type "c"
Sleep 500ms

# Part 2: Navigation input (orange feedback)
# Navigate to "5" button (row 1, col 1) and press with Enter
Down
Sleep 300ms
Right
Sleep 300ms
Enter
Sleep 500ms

# Navigate to "*" button (row 1, col 3)
Right
Sleep 300ms
Right
Sleep 300ms
Enter
Sleep 500ms

# Navigate to "4" button (row 2, col 0)
Down
Sleep 300ms
Left
Sleep 300ms
Left
Sleep 300ms
Left
Sleep 300ms
Enter
Sleep 500ms

# Navigate to "=" button (row 4, col 2)
Down
Sleep 300ms
Down
Sleep 300ms
Right
Sleep 300ms
Right
Sleep 300ms
Enter
Sleep 800ms

# Quit
Type "q"
Sleep 600ms
