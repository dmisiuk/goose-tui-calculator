Output .tapes/assets/feature-visual-feedback.gif

Set FontSize 18
Set Width 1000
Set Height 600

Type "go run ./cmd/calculator"
Enter
Sleep 2000ms

# Navigate to show button highlighting
Right
Sleep 300ms
Right
Sleep 300ms
Down
Sleep 300ms
Left
Sleep 300ms

# Show button press feedback
Enter
Sleep 400ms

# Simple calculation with visual feedback
Type "7"
Sleep 300ms
Type "+"
Sleep 300ms
Type "3"
Sleep 300ms
Type "="
Sleep 500ms

Type "q"
