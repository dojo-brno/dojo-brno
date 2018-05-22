package poker

type CardValue int

const (
	Two CardValue = iota
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

type CardSuit int

const (
	Hearts CardSuit = iota
	Spades
	Diamonds
	Clubs
)

type Card struct {
	Value CardValue
	Suit  CardSuit
}

type Hand [5]Card

func IsStrainghtFlush(hand Hand) bool {

	for i := 0; i < len(hand)-1; i++ {
		if (int(hand[i+1].Value) != (int(hand[i].Value) + 1)) || (hand[i+1].Suit != hand[i].Suit) {
			return false
		}
	}
	return true
}
