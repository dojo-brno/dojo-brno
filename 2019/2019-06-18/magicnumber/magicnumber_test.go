package magicnumber

import "testing"

func TestGetMagicNumber(t *testing.T) {
	want := 123

	if got := GetMagicNumber(); got != want {
		t.Errorf("GetMagicNumber(): %v,	WANT: %v", got, want)
	}
}

func TestGetSuperMagicNumber(t *testing.T) {
	superMagicNumber := 13

	if got := GetSuperMagicNumber(); got != superMagicNumber {
		t.Errorf("GetSuperMagicNumber(): %v,		WANT: %v", got, superMagicNumber)
	}
}
