package codecracker

import "testing"

func TestDecrypt(t *testing.T) {
	tests := []struct {
		msg  string
		want string
	}{
		{"", ""},
		{"!", "a"},
		{")", "b"},
		{"!&d<", "ahoj"},
		{"!)\"(£*%&><@abcdefghijklmno", "abcdefghijklmnopqrstuvwxyz"},
		{" ", ""},
		{"z", ""},
	}

	for _, tt := range tests {
		if got := Decrypt(tt.msg); got != tt.want {
			t.Errorf("Decrypt(%q) = %q WANT %q", tt.msg, got, tt.want)
		}
	}
}

func TestEncrypt(t *testing.T) {
	tests := []struct {
		msg  string
		want string
	}{
		{"", ""},
		{"a", "!"},
		{"ahoj", "!&d<"},
		{"abcdefghijklmnopqrstuvwxyz", "!)\"(£*%&><@abcdefghijklmno"},
		{"a!h o%j", "!&d<"},
	}

	for _, tt := range tests {
		if got := Encrypt(tt.msg); got != tt.want {
			t.Errorf("Encrypt(%q) = %q WANT %q", tt.msg, got, tt.want)
		}
	}
}
