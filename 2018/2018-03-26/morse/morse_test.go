package morse

import "testing"

func TestDecrypt(t *testing.T) {
	tests := []struct {
		morseCode string
		want      string
	}{
		{morseCode: "", want: ""},
		{morseCode: ".", want: "E"},
		{morseCode: "-", want: "T"},
		{morseCode: ".-", want: "A"},
		{morseCode: "-...", want: "B"},
		{morseCode: "....", want: "H"},
		{morseCode: ".-/..../.-", want: "AHA"},
		{morseCode: "---", want: "O"},
		{morseCode: ".---", want: "J"},
		{morseCode: ".-/..../---/.---", want: "AHOJ"},
	}

	for _, tt := range tests {
		if got := Decrypt(tt.morseCode); got != tt.want {
			t.Errorf("Decrypt(%q) = %q WANT %q", tt.morseCode, got, tt.want)
		}
	}
}
