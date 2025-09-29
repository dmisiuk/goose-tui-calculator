Output calc-demo.gif
Set FontSize 18
Set Width 1000
Set Height 600

# Build binary for stable run speed
Type "go build -o goosecalc ./cmd/calculator"
Enter
Sleep 1600ms

# Launch calculator TUI
Type "./goosecalc"
Enter
Sleep 2000ms  # Allow Bubble Tea to init alternate screen

# 7 + 3 HONK -> 10
Type "7"
Sleep 300ms
Type "+"
Sleep 300ms
Type "3"
Sleep 300ms
Type "H"  # HONK button (alias for equals)
Sleep 800ms  # Show result

# Quit application
Type "q"
Sleep 600ms