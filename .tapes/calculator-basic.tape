Output .tapes/assets/calculator-basic.gif
Set FontSize 18
Set Width 1000
Set Height 600

# Pre-build binary for stability (cold compile can be slow on CI)
Type "go build -o calc ./cmd/calculator\n"
Sleep 1600ms

# First run session
Type "./calc\n"
Sleep 450ms
Type "2"
Sleep 120ms
Type "+"
Sleep 120ms
Type "3"
Sleep 120ms
Type "="
Sleep 500ms
Type "q"
Sleep 350ms

# Second run session (fresh state)
Type "./calc\n"
Sleep 450ms
Type "9"
Sleep 120ms
Type "x"
Sleep 120ms
Type "4"
Sleep 120ms
Type "="
Sleep 500ms
Type "q"
