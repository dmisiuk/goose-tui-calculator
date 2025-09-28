# VHS Best Practices for Terminal App Demos

## Critical Fix: Command Execution

### ❌ Wrong - Types literal newline
```tape
Type "go run ./cmd/calculator\n"
```

### ✅ Correct - Executes command
```tape
Type "go run ./cmd/calculator"
Enter
```

## Timing Requirements for TUI Applications

### Bubble Tea App Launch Timing
- **After `Enter` for pre-built binary**: 2000ms minimum
- **After `Enter` for `go run` (cold compile)**: 3000-5000ms
- **Between TUI keystrokes**: 300ms minimum
- **After displaying results**: 800ms for visibility

### Why Timing Matters
TUI applications need time to:
1. Initialize alternate screen buffer
2. Set up raw terminal input mode
3. Take control of keyboard input from shell

Without adequate timing, keystrokes go to the shell instead of the app.

## Standard Pattern Template

```tape
Output app-demo.gif
Set FontSize 18
Set Width 1000
Set Height 600

# Pre-build for stability
Type "go build -o app ./cmd/app"
Enter
Sleep 1600ms

# Launch with adequate timing
Type "./app"
Enter
Sleep 2000ms  # Critical: TUI initialization time

# Interact with adequate pauses
Type "1"
Sleep 300ms
Type "+"
Sleep 300ms
Type "="
Sleep 800ms  # Show result
Type "q"
Sleep 600ms  # Clean exit
```

## Development Workflow

### Creating New Tapes
1. Copy baseline tape as template
2. Test locally: `vhs your-tape.tape`
3. Verify keystrokes go to TUI, not shell
4. Commit both `.tape` and `.gif`

### Debugging Checklist
- [ ] Commands execute (not typed as text)
- [ ] TUI captures all input keystrokes
- [ ] Results display clearly
- [ ] GIF file size reasonable

## Quick Reference

### Most Common Fix
```diff
- Type "./calc\n"
+ Type "./calc"
+ Enter
+ Sleep 2000ms
```

### Standard Timings
```tape
Enter
Sleep 2000ms    # After app launch
Type "key"
Sleep 300ms     # Between keystrokes
Type "="
Sleep 800ms     # Show results
Type "q"
Sleep 600ms     # Clean exit
```

## Common Issues

### Issue 1: Characters appear in shell instead of TUI
**Cause**: TUI hasn't captured input yet
**Fix**: Increase Sleep after Enter (try 2000ms+)

### Issue 2: Commands don't execute
**Cause**: Using `\n` in Type command
**Fix**: Replace with separate Enter command

### Issue 3: Demo too fast to follow
**Cause**: Short delays between actions
**Fix**: Use 300ms+ between keystrokes, 800ms for results

## Testing
Always test locally before committing:
```bash
vhs .tapes/your-tape.tape
# Verify GIF shows correct behavior
```
