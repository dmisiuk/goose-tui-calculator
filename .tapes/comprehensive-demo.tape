Output .tapes/assets/comprehensive-demo.gif
Set FontSize 18
Set Width 1000
Set Height 600

# Pre-build calculator binary (improves CI speed)
Type "go build -o calc ./cmd/calculator"
Enter
Sleep 1600ms

# ────────────────────────────────────────────────
# Session 1 – Direct keyboard input (2 + 3 = 5)
# ────────────────────────────────────────────────
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
Sleep 1200ms

# Clean screen with AC ("c" is mapped to AC)
Type "c"
Sleep 1200ms

# ────────────────────────────────────────────────
# Session 2 – Arrow-key navigation (9 × 4 = 36)
# ────────────────────────────────────────────────
# Navigate to 9
Type "{{down}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type " "
Sleep 1500ms

# Navigate to ×
Type "{{right}}"
Sleep 400ms
Type " "
Sleep 1500ms

# Navigate to 4
Type "{{down}}"
Sleep 400ms
Type "{{left}}"
Sleep 400ms
Type "{{left}}"
Sleep 400ms
Type "{{left}}"
Sleep 400ms
Type " "
Sleep 1500ms

# Navigate to =
Type "{{down}}"
Sleep 400ms
Type "{{down}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type "{{right}}"
Sleep 400ms
Type " "
Sleep 1500ms

# Clean screen again
Type "c"
Sleep 1200ms

# ────────────────────────────────────────────────
# Session 3 – Error state (8 / 0 = Error)
# ────────────────────────────────────────────────
Type "8"
Sleep 300ms
Type "/"
Sleep 300ms
Type "0"
Sleep 300ms
Type "="
Sleep 1500ms

# Quit application
Type "q"
Sleep 1000ms
