# Audio Feedback for Button Presses

## Why
Enhance the calculator's tactile experience by providing auditory feedback when buttons are pressed, making the TUI calculator feel more authentic and responsive to user input. Different button types will have distinct sounds to reinforce their purpose, with special emphasis on the critical AC (clear) and = (equals) buttons.

## What Changes
- Add audio feedback system that plays sounds on button press
- Implement three distinct sound types:
  - Number keys (0-9): One sound signature
  - Standard functional/operator keys (+/-, %, +, -, x, /, .): Different sound signature
  - Special action keys (AC, =): Unique sound signature to emphasize their special functionality
- Integrate audio playback into existing button press handling (`handleButtonPress`)
- Support cross-platform audio playback (Linux, macOS, Windows)

## Impact
- Affected specs: `audio-feedback` (new capability)
- Affected code:
  - `internal/calculator/calculator.go` - Add audio playback calls in `handleButtonPress`
  - New audio module for sound generation/playback
  - May require external audio library or OS-level audio commands
- Dependencies: May need to add audio library to go.mod (e.g., beep, oto, or system calls)
- Build size: Minimal increase if using system audio commands; moderate if embedding audio library
