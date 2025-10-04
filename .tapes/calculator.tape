# VHS tape for generating a demo of the TUI calculator.
# This tape demonstrates basic arithmetic operations and showcases the UI.
#
# For more information on VHS, see: https://github.com/charmbracelet/vhs

Output .tapes/assets/calculator-demo.gif
Set Theme "Catppuccin Mocha"
Set FontSize 18
Set Width 1000
Set Height 600

# Set environment variables to ensure colors are rendered correctly in CI.
Env TERM "xterm-256color"
Env COLORTERM "truecolor"
Env CLICOLOR_FORCE "1"

# Pre-build the binary for stability, as cold compiles can be slow on CI.
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# First session: 2 + 3 = 5
Type "./calc"
Enter
Sleep 2000ms
Type "2"
Sleep 300ms
Type "+"
Sleep 300ms
Type "3"
Sleep 300ms
Type "="
Sleep 800ms
Type "q"
Sleep 600ms

# Second session: 9 x 4 = 36
Type "./calc"
Enter
Sleep 2000ms
Type "9"
Sleep 300ms
Type "x"
Sleep 300ms
Type "4"
Sleep 300ms
Type "="
Sleep 800ms
Type "q"
Sleep 1000ms