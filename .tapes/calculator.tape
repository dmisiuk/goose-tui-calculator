Output .tapes/assets/calculator-alias.gif
Set FontSize 18
Set Width 1000
Set Height 600

# This tape is an alias of calculator-basic.tape
# It ensures CI demo step won't fail if .tapes/calculator.tape is expected

Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

Type "./calc"
Enter
Sleep 2000ms

Type "{{down}}"
Sleep 400ms
Type "{{down}}"
Sleep 400ms
Type "{{down}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type " "
Sleep 1500ms

Type "{{right}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type " "
Sleep 1500ms

Type "{{left}}"
Sleep 400ms
Type "{{left}}"
Sleep 400ms
Type " "
Sleep 1500ms

Type "{{right}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type " "
Sleep 1500ms

Type "q"
Sleep 1000ms
