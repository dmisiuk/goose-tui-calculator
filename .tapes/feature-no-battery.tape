Output .tapes/assets/feature-no-battery.gif
Set FontSize 18
Set Width 1200
Set Height 800

# Pre-build binary for stability
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Launch calculator - now without battery indicator
Type "./calc"
Enter
Sleep 2500ms

# Show basic calculation with direct keyboard input (blue style)
Type "7"
Sleep 500ms
Type "+"
Sleep 800ms
Type "3"
Sleep 800ms
Type "="
Sleep 1500ms

# Navigate with arrows (orange style)
Down
Sleep 400ms
Down
Sleep 400ms
Left
Sleep 400ms
Up
Sleep 400ms
Up
Sleep 400ms
# Press AC
Enter
Sleep 800ms

# Another calculation
Type "5"
Sleep 400ms
Type "x"
Sleep 400ms
Type "8"
Sleep 400ms
Type "="
Sleep 1500ms

# Exit
Type "q"
Sleep 600ms
