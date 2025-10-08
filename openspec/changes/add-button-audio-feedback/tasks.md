# Implementation Tasks

## 1. Audio Infrastructure
- [ ] 1.1 Research and select audio library or approach (system beep, Go audio library, or platform-specific calls)
- [ ] 1.2 Create `internal/audio/audio.go` module with `PlayNumberSound()`, `PlayFunctionalSound()`, and `PlaySpecialActionSound()` functions
- [ ] 1.3 Implement cross-platform audio detection and graceful degradation when audio unavailable
- [ ] 1.4 Add necessary dependencies to `go.mod` (if using external audio library)

## 2. Sound Generation
- [ ] 2.1 Define or generate number button sound (frequency/tone parameters or audio file)
- [ ] 2.2 Define or generate standard functional button sound (different frequency/tone or audio file)
- [ ] 2.3 Define or generate special action button sound for AC and = (distinct frequency/tone or audio file)
- [ ] 2.4 Ensure sounds are brief (<100ms) to avoid disrupting calculator flow
- [ ] 2.5 Test sound playback on macOS, Linux, and Windows

## 3. Calculator Integration
- [ ] 3.1 Import audio module in `internal/calculator/calculator.go`
- [ ] 3.2 Modify `handleButtonPress()` to call appropriate audio function based on button type (number, functional, or special action)
- [ ] 3.3 Ensure audio plays asynchronously and doesn't block UI updates
- [ ] 3.4 Verify existing visual feedback timing (300ms) doesn't conflict with audio

## 4. Testing
- [ ] 4.1 Add unit tests for audio button categorization (number, standard functional, special action)
- [ ] 4.2 Add integration tests verifying correct audio calls are made for each button type
- [ ] 4.3 Test manual playback on macOS, Linux, and Windows
- [ ] 4.4 Verify AC and = buttons produce distinct sounds from other buttons
- [ ] 4.5 Verify graceful degradation when audio system unavailable
- [ ] 4.6 Run `go test -race ./...` to ensure no concurrency issues

## 5. Documentation and Build
- [ ] 5.1 Update README or documentation mentioning audio feedback feature
- [ ] 5.2 Ensure `go mod tidy` produces no diff
- [ ] 5.3 Verify single binary build still works across platforms
- [ ] 5.4 Consider optional VHS demo showing audio feature (note: VHS may not capture audio)
