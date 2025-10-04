Output .tapes/assets/calculator-basic.gif
Output .tapes/assets/calculator-basic.txt

Set FontSize 18
Set Width 1200
Set Height 800

# Force colors for proper rendering
Env CLICOLOR_FORCE 1
Env FORCE_COLOR 1

# Pre-build binary for stability (cold compile can be slow on CI)
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
