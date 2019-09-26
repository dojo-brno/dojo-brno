package ant

import "testing"

func GetMagicNumber() int {
	return 333
}

func TestGetMagicNumber(t *testing.T) {
	if got := GetMagicNumber(); got != 333 {
		t.Errorf("GetMagicNumber():%v, WANT: 333", got)
	}
}
