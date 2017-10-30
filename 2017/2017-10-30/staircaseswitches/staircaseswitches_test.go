package staircaseswitches

import (
	"fmt"
	"testing"
)

func TestGetLightBulb(t *testing.T) {
	tests := []struct {
		currentLightBulbState bool
		switchState           [2]bool
		switchIsPressed       [2]bool
		nextLightBulbState    bool
	}{
		{
			currentLightBulbState: false,
			switchState:           [2]bool{false, false},
			switchIsPressed:       [2]bool{false, false},
			nextLightBulbState:    false,
		},
		{
			currentLightBulbState: true,
			switchState:           [2]bool{false, false},
			switchIsPressed:       [2]bool{false, false},
			nextLightBulbState:    true,
		},
		{
			currentLightBulbState: true,
			switchState:           [2]bool{false, false},
			switchIsPressed:       [2]bool{true, false},
			nextLightBulbState:    false,
		},
		{
			currentLightBulbState: true,
			switchState:           [2]bool{false, false},
			switchIsPressed:       [2]bool{false, true},
			nextLightBulbState:    false,
		},
		{
			currentLightBulbState: true,
			switchState:           [2]bool{false, false},
			switchIsPressed:       [2]bool{true, true},
			nextLightBulbState:    true,
		},
	}
	for _, tt := range tests {
		lightBulb := CreateLightBulb(tt.currentLightBulbState, tt.switchState)
		origLightBulb := fmt.Sprintf("lightBulb(%v)", lightBulb)
		if got := lightBulb.NextLightBulbState(tt.switchIsPressed); got != tt.nextLightBulbState {
			t.Errorf("%v.NextLightBulbState(%v) = %v, WANT: %v", origLightBulb, tt.switchIsPressed, got, tt.nextLightBulbState)
		}
	}
}

func TestGetLightBulbSequence(t *testing.T) {
	lightBulb := CreateLightBulb(false, [2]bool{false, false})
	tests := []struct {
		switchIsPressed    [2]bool
		nextLightBulbState bool
	}{
		{
			switchIsPressed:    [2]bool{false, false},
			nextLightBulbState: false,
		},
		{
			switchIsPressed:    [2]bool{true, false},
			nextLightBulbState: true,
		},
		{
			switchIsPressed:    [2]bool{false, true},
			nextLightBulbState: false,
		},
	}
	for i, tt := range tests {
		origLightBulb := fmt.Sprintf("lightBulb(%v)", lightBulb)
		if got := lightBulb.NextLightBulbState(tt.switchIsPressed); got != tt.nextLightBulbState {
			t.Errorf("%v: %v.NextLightBulbState(%v) = %v, WANT: %v", i, origLightBulb, tt.switchIsPressed, got, tt.nextLightBulbState)
		}
	}
}

func TestNewLightBulb(t *testing.T) {
	var lightBulb LightBulb
	if true {
		switchState := [2]bool{true, true}
		lightBulb = CreateLightBulb(true, switchState)
	}
	if want := true; lightBulb.lightBulbState != want {
		t.Errorf("lightBulb.lightBulbState = %v, WANT %v", lightBulb.lightBulbState, want)
	}
	if want := true; lightBulb.switchState[0] != want {
		t.Errorf("lightBulb.switchState[0] = %v, WANT %v", lightBulb.switchState[0], want)
	}
	if want := true; lightBulb.switchState[1] != want {
		t.Errorf("lightBulb.switchState[1] = %v, WANT %v", lightBulb.switchState[1], want)
	}
}

func TestGetSwitchState(t *testing.T) {
	tests := []struct {
		lightBulbState     bool
		currentSwitchState [2]bool
		switchIsPressed    [2]bool
		nextSwitchState    [2]bool
	}{
		{
			lightBulbState:     false,
			currentSwitchState: [2]bool{false, false},
			switchIsPressed:    [2]bool{false, false},
			nextSwitchState:    [2]bool{false, false},
		},
		{
			lightBulbState:     false,
			currentSwitchState: [2]bool{false, false},
			switchIsPressed:    [2]bool{true, true},
			nextSwitchState:    [2]bool{true, true},
		},
		{
			lightBulbState:     false,
			currentSwitchState: [2]bool{true, true},
			switchIsPressed:    [2]bool{true, true},
			nextSwitchState:    [2]bool{false, false},
		},
	}
	for _, tt := range tests {
		lightBulb := CreateLightBulb(tt.lightBulbState, tt.currentSwitchState)
		origLightBulb := fmt.Sprintf("lightBulb(%v)", lightBulb)
		lightBulb.NextLightBulbState(tt.switchIsPressed)
		if got := lightBulb.GetSwitchStates(); got != tt.nextSwitchState {
			t.Errorf("tt.switchIsPressed: (%v), %v.NextSwitchState() = %v, WANT: %v", tt.switchIsPressed, origLightBulb, got, tt.nextSwitchState)
		}
	}
}
