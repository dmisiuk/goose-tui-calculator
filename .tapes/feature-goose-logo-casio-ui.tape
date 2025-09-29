Output .tapes/assets/feature-goose-logo-casio-ui.gif
Set FontSize 18
Set Width 1200
Set Height 800

# Pre-build binary for stability (cold compile can be slow on CI)
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# Demonstrate the new Casio-style calculator with Goose logo
Type "./calc"
Enter
Sleep 2500ms

# Show basic calculation with direct keyboard input (blue style)
Type "4"
Sleep 500ms
Type "-"
Sleep 800ms # Pause to show "4 -" on green LCD display
Type "5"
Sleep 800ms # Pause to show "5" with "4 -" above

# Calculate to show result
Type "="
Sleep 1500ms # Pause to show the result with Goose logo and battery indicator

# Show keyboard navigation (orange style) - navigate to AC button
Down
Sleep 400ms
Down
Sleep 400ms
Up
Sleep 400ms
Left
Sleep 400ms
Up
Sleep 400ms
Up
Sleep 400ms
# Now on AC button
Enter # Press AC
Sleep 800ms

# Quick calculation
Type "9"
Sleep 400ms
Type "x"
Sleep 400ms
Type "6"
Sleep 400ms
Type "="
Sleep 1500ms

# Exit
Type "q"
Sleep 600ms
