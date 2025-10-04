Output .tapes/assets/feature-color-verification.gif
Set FontSize 18
Set Width 1000
Set Height 600

Env TERM xterm-256color
Env COLORTERM truecolor
Env CLICOLOR_FORCE 1
Env FORCE_COLOR 1

Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

Type "./calc"
Enter
Sleep 1500ms

Right
Sleep 400ms
Right
Sleep 400ms
Space
Sleep 1200ms

Type "q"
Sleep 800ms
