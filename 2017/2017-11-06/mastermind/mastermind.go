package mastermind

const (
	White PegColour = iota
	Yellow
	Pink
	Blue
	Red
	Black
	Brown
	Gray
)

type PegColour int

var PegColourToString = map[PegColour]string{
	White:  "White",
	Yellow: "Yellow",
	Pink:   "Pink",
	Blue:   "Blue",
	Red:    "Red",
	Black:  "Black",
	Brown:  "Brown",
	Gray:   "Gray",
}

func (p PegColour) String() string {
	if v, ok := PegColourToString[p]; ok {
		return v
	}
	return "Unknown"
}

const (
	numberOfPegs = 5
)

type Pegs [numberOfPegs]PegColour

type colourCounters struct {
	secret int
	guess  int
}

func MasterMind(secrets Pegs, guesses Pegs) (int, int) {
	potentialMisplaced := make(map[PegColour]colourCounters)
	wellPlaced := 0
	for i, secret := range secrets {
		if secret == guesses[i] {
			wellPlaced++
		} else {
			potentialMisplaced = colourCountersInc(potentialMisplaced, secret, "secret")
			potentialMisplaced = colourCountersInc(potentialMisplaced, guesses[i], "guess")
		}
	}
	misPlaced := 0
	for _, pm := range potentialMisplaced {
		misPlaced += minimum(pm.secret, pm.guess)
	}
	return wellPlaced, misPlaced
}

func colourCountersInc(current map[PegColour]colourCounters, colour PegColour, field string) map[PegColour]colourCounters {
	incremented := current
	if field == "secret" {
		incremented[colour] = colourCounters{secret: current[colour].secret + 1, guess: current[colour].guess}
	} else {
		incremented[colour] = colourCounters{secret: current[colour].secret, guess: current[colour].guess + 1}
	}
	return incremented
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
