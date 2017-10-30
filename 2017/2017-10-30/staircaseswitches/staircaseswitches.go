package staircaseswitches

type LightBulb struct {
	lightBulbState bool
	switchState    [2]bool
}

func CreateLightBulb(lightBulbState bool, switchState [2]bool) LightBulb {
	var l LightBulb
	l.lightBulbState = lightBulbState
	l.switchState = switchState
	return l
}

func (l *LightBulb) NextLightBulbState(switchIsPressed [2]bool) bool {
	if (switchIsPressed[0] || switchIsPressed[1]) && switchIsPressed[0] != switchIsPressed[1] {
		l.lightBulbState = !l.lightBulbState
	}
	for i, _ := range l.switchState {
		if switchIsPressed[i] {
			l.switchState[i] = !l.switchState[i]
		}
	}
	return l.lightBulbState
}

func (l *LightBulb) GetSwitchStates() [2]bool {
	return l.switchState
}
