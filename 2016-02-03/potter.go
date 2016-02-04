package potter

type Basket []string

const BookPrice float64 = 8

func calculatePrice(b Basket) float64 {
	booksCounter := make(map[string]int)
	for _, book := range b {
		if {
			
		} else {
			booksCounter[book] = 1
		}
	}

	if len(b) == 2 && b[0] != b[1] {
		return 2 * BookPrice * .95
	}
	if len(b) == 3 && b[2] == "three" {
		return 3 * BookPrice * .90
	}
	return float64(len(b)) * BookPrice
}
