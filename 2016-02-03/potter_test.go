package potter

import "testing"

func TestNoBooks(t *testing.T) {
	basket := Basket{}
	if calculatePrice(basket) != 0 {
		t.Fail()
	}
}

func TestBooksOne(t *testing.T) {
	basket := Basket{"one"}
	if got := calculatePrice(basket); got != BookPrice {
		t.Errorf("calculatePrice(%v) = %v, want %v", basket, got, BookPrice)
	}
}

func TestTwoEqualBooks(t *testing.T) {
	basket := Basket{"one", "one"}
	if got := calculatePrice(basket); got != 2*BookPrice {
		t.Errorf("calculatePrice(%v) = %v, want %v", basket, got, 2*BookPrice)
	}
}

func TestTwoDifferentBooks(t *testing.T) {
	basket := Basket{"one", "two"}
	want := 2 * BookPrice * .95
	if got := calculatePrice(basket); got != want {
		t.Errorf("calculatePrice(%v) = %v, want %v", basket, got, want)
	}
}

func TestThreeDifferentBooks(t *testing.T) {
	basket := Basket{"one", "two", "three"}
	want := 3 * BookPrice * .90
	if got := calculatePrice(basket); got != want {
		t.Errorf("calculatePrice(%v) = %v, want %v", basket, got, want)
	}
}

func TestThreeBooksAllEqual(t *testing.T) {
	basket := Basket{"one", "one", "one"}
	want := 3 * BookPrice
	if got := calculatePrice(basket); got != want {
		t.Errorf("calculatePrice(%v) = %v, want %v", basket, got, want)
	}
}
