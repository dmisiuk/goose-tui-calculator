package audio

import (
	"sync"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/generators"
	"github.com/gopxl/beep/v2/speaker"
)

var (
	// Speaker initialization
	speakerInitialized bool
	speakerMutex       sync.Mutex
	audioEnabled       bool = true
)

const (
	sampleRate = beep.SampleRate(48000)
	duration   = time.Millisecond * 50 // Brief sound: 50ms
)

// ButtonType represents different categories of calculator buttons
type ButtonType int

const (
	ButtonTypeNumber ButtonType = iota
	ButtonTypeFunctional
	ButtonTypeSpecialAction
)

// init initializes the audio speaker once
func init() {
	speakerMutex.Lock()
	defer speakerMutex.Unlock()

	if !speakerInitialized {
		err := speaker.Init(sampleRate, sampleRate.N(time.Millisecond*100))
		if err != nil {
			// If speaker initialization fails, disable audio
			audioEnabled = false
		} else {
			speakerInitialized = true
			audioEnabled = true
		}
	}
}

// playTone plays a tone with the specified frequency
func playTone(frequency float64) {
	if !audioEnabled {
		return
	}

	// Generate a sine wave tone
	tone, err := generators.SineTone(sampleRate, frequency)
	if err != nil {
		return
	}

	// Limit the tone duration
	limited := beep.Take(sampleRate.N(duration), tone)

	// Play asynchronously to avoid blocking
	speaker.Play(limited)
}

// PlayNumberSound plays the sound for number buttons (0-9)
// Uses a mid-range frequency (600 Hz)
func PlayNumberSound() {
	playTone(600)
}

// PlayFunctionalSound plays the sound for functional/operator buttons (+, -, *, /, %, +/-, .)
// Uses a higher frequency (800 Hz)
func PlayFunctionalSound() {
	playTone(800)
}

// PlaySpecialActionSound plays the sound for special action buttons (AC, =)
// Uses a distinct lower frequency (400 Hz) with slightly longer duration for emphasis
func PlaySpecialActionSound() {
	if !audioEnabled {
		return
	}

	// Generate a sine wave tone
	tone, err := generators.SineTone(sampleRate, 400)
	if err != nil {
		return
	}

	// Slightly longer duration for special actions (80ms)
	limited := beep.Take(sampleRate.N(time.Millisecond*80), tone)

	// Play asynchronously to avoid blocking
	speaker.Play(limited)
}

// GetButtonType determines the button type based on the button label
func GetButtonType(button string) ButtonType {
	switch button {
	case "AC", "=":
		return ButtonTypeSpecialAction
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return ButtonTypeNumber
	default:
		// +, -, x, /, %, +/-, .
		return ButtonTypeFunctional
	}
}

// PlayButtonSound plays the appropriate sound based on the button type
func PlayButtonSound(button string) {
	buttonType := GetButtonType(button)
	switch buttonType {
	case ButtonTypeNumber:
		PlayNumberSound()
	case ButtonTypeFunctional:
		PlayFunctionalSound()
	case ButtonTypeSpecialAction:
		PlaySpecialActionSound()
	}
}

// IsEnabled returns whether audio is currently enabled
func IsEnabled() bool {
	return audioEnabled
}
