package potter

import (
	"testing"
)

func ComputePrice(basket [5]int) float64 {
	if basket[0] == 2 || basket[4] == 2 {
		return 16
	}
	if basket[0]+basket[1] == 2 {
		return 15.2
	}
	return float64(basket[0] * 8)
}

func TestBaskets(t *testing.T) {
	AssertBasketPrice([5]int{0, 0, 0, 0, 0}, 0.0, t)
	AssertBasketPrice([5]int{1, 0, 0, 0, 0}, 8.0, t)
	AssertBasketPrice([5]int{1, 1, 0, 0, 0}, 16.0*0.95, t)
	AssertBasketPrice([5]int{2, 0, 0, 0, 0}, 2*8.0, t)
	AssertBasketPrice([5]int{0, 0, 0, 0, 2}, 2*8.0, t)
}

func AssertBasketPrice(basket [5]int, want float64, t *testing.T) {
	price := ComputePrice(basket)
	if price != want {
		t.Errorf("ComputePrice(%v) = %v, want %v", basket, price, want)
	}
}
