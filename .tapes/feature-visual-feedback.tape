Output .tapes/assets/feature-visual-feedback.gif

Set FontSize 18
Set Width 1000
Set Height 600

Env TERM xterm-256color
Env COLORTERM truecolor

Type "go run ./cmd/calculator"
Enter
Sleep 2000ms

# Complete calculation using only arrow keys and Enter/Space
# Navigate to "7" button (row 1, col 0)
Down
Sleep 500ms
Enter
Sleep 500ms

# Navigate to "+" button (row 3, col 3)
Down
Sleep 300ms
Down
Sleep 300ms
Right
Sleep 300ms
Right
Sleep 300ms
Right
Sleep 300ms
Enter
Sleep 500ms

# Navigate to "3" button (row 2, col 2)
Up
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
Enter
Sleep 800ms

# Show result in display, then quit
Type "q"
