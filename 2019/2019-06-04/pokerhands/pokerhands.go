package pokerhands

type Hand [5]struct {
	Value  CardValue
	Colour CardColour
}

type CardValue int

const (
	Two CardValue = 2 + iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type CardColour int

const (
	Clubs CardColour = 1 + iota
	Diamonds
	Hearts
	Spades
)

func (c CardColour) String() string {
	switch c {
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	}
	return "N/A"
}

type Result int

const (
	AWon Result = 1 + iota
	BWon
	Tie
)

func (r Result) String() string {
	switch r {
	case AWon:
		return "AWon"
	case BWon:
		return "BWon"
	case Tie:
		return "Tie"
	}
	return "N/A"
}

func Compare(a, b Hand) Result {
	a.sort()
	b.sort()
	aIsOnePair, aOnePairValue := isOnePairAndPairValue(a)
	bIsOnePair, bOnePairValue := isOnePairAndPairValue(b)
	if aIsOnePair && bIsOnePair {
		if aOnePairValue > bOnePairValue {
			return AWon
		}
		if aOnePairValue < bOnePairValue {
			return BWon
		}
		return compareHighCard(a, b)
	}
	if isOnePair(a) {
		return AWon
	}
	if isOnePair(b) {
		return BWon
	}
	return compareHighCard(a, b)
}

func isOnePairAndPairValue(h Hand) (bool, int) {
	numberOfPairs := 0
	pairValue := 0
	for i := 0; i < len(h)-1; i++ {
		if h[i].Value == h[i+1].Value {
			numberOfPairs++
			pairValue = int(h[i].Value) + int(h[i+1].Value)
		}
	}
	if numberOfPairs == 1 {
		return true, pairValue
	}
	return false, 0
}

func isOnePair(h Hand) bool {
	numberOfPairs := 0
	for i := 0; i < len(h)-1; i++ {
		if h[i].Value == h[i+1].Value {
			numberOfPairs++
		}
	}
	if numberOfPairs == 1 {
		return true
	}
	return false
}

func compareHighCard(a, b Hand) Result {
	for i := range a {
		if a[i].Value > b[i].Value {
			return AWon
		}
		if a[i].Value < b[i].Value {
			return BWon
		}
	}
	return Tie
}

func (h *Hand) sort() {
	for i := range h {
		for j := i + 1; j < len(h); j++ {
			if h[i].Value < h[j].Value {
				tempCard := h[j]
				h[j] = h[i]
				h[i] = tempCard
			}
		}
	}
}
