Output .tapes/assets/calculator-alias.gif
Set Theme "Catppuccin Mocha"
Set FontSize 18
Set Width 1000
Set Height 600

Env TERM xterm-256color
Env COLORTERM truecolor
Env CLICOLOR_FORCE "1"

# This tape is an alias of calculator-basic.tape
# It ensures CI demo step won't fail if .tapes/calculator.tape is expected

Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

Type "./calc"
Enter
Sleep 1500ms

Down
Sleep 400ms
Down
Sleep 400ms
Down
Sleep 400ms
Right
Sleep 400ms
Space
Sleep 1500ms

Right
Sleep 400ms
Right
Sleep 400ms
Right
Sleep 400ms
Space
Sleep 1500ms

Left
Sleep 400ms
Left
Sleep 400ms
Space
Sleep 1500ms

Right
Sleep 400ms
Right
Sleep 400ms
Space
Sleep 1500ms

Type "q"
Sleep 1000ms