package search

import "testing"

func TestSearchEmptyList(t *testing.T) {
	var list []int
	want := false
	for n := -100; n < 100; n++ {
		if got := SearchSorted(n, list); got != want {
			t.Errorf("Search(%v, %v) = %v, want %v", n, list, got, want)
		}
	}
}
func TestSearchNumberListHead(t *testing.T) {
	want := true
	for n := -100; n < 100; n++ {
		list := []int{n}
		if got := SearchSorted(n, list); got != want {
			t.Errorf("Search(%v, %v) = %v, want %v", n, list, got, want)
		}
	}
}

func TestNotEmptyListNumberNotFound(t *testing.T) {
	want := false
	for n := -100; n < 100; n++ {
		list := []int{n + 1}
		if got := SearchSorted(n, list); got != want {
			t.Errorf("Search(%v, %v) = %v, want %v", n, list, got, want)
		}
	}
}

func TestReturnTrueWhenNumberIsInListButNotAtListHead(t *testing.T) {
	want := true
	for n := -100; n < 100; n++ {
		list := []int{n - 1, n}
		if got := SearchSorted(n, list); got != want {
			t.Errorf("Search(%v, %v) = %v, want %v", n, list, got, want)
		}
	}
}

func TestFindTheLastOne(t *testing.T) {
	want := true
	list := []int{1, 2, 3, 4, 5}
	n := 4
	if got := SearchSorted(n, list); got != want {
		t.Errorf("Search(%v, %v) = %v, want %v", n, list, got, want)
	}

}
