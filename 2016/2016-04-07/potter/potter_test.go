package potter

import "testing"

// t.Logf(string)
// t.FailNow() - marks as fail
// t.Fatal() - marks as fail and exits
// t.Fatalf() -

func PriceOf(books string) float32 {
	totalPrice := float32(bookPrice * len(books))
	if len(books) >= 2 {
		firstBook := books[0]
		secondBook := books[1]
		if firstBook != secondBook {
			// Compute discount
			totalPrice -= 0.8
		}
	}
	if len(books) == 3 && books[2]
	return totalPrice
}

const bookPrice = 8

const (
	discount2 = 0.95 // 5%
	discount3 = 0.90 // 10%
)

func TestPotter(t *testing.T) {
	tests := []struct {
		books string
		want  float32
	}{
		{
			books: "",
			want:  0,
		},
		{
			books: "A",
			want:  bookPrice,
		},
		{
			books: "B",
			want:  bookPrice,
		},
		{
			books: "AA",
			want:  2 * bookPrice,
		},
		{
			books: "AB",
			want:  2 * bookPrice * discount2,
		},
		{
			books: "AAA",
			want:  3 * bookPrice,
		},
		{
			books: "AAAA",
			want:  4 * bookPrice,
		},
		{
			books: "AAAAA",
			want:  5 * bookPrice,
		},
		{
			books: "ABA",
			want:  2*bookPrice*discount2 + bookPrice,
		},
		{
			books: "ABB",
			want:  2*bookPrice*discount2 + bookPrice,
		},
		{
			books: "ABC",
			want:  3 * bookPrice * discount3,
		},
	}
	for _, tt := range tests {
		if got := PriceOf(tt.books); got != tt.want {
			t.Errorf("PriceOf(%q) = %.2f, want %.2f", tt.books, got, tt.want)
		}
	}
}
