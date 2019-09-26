package coder

import "testing"

func TestEncrypt(t *testing.T) {
	// a b c d e f g h i j k l m n o p q r s t u v w x y z
	// ! ) " ( £ * % & > < @ a b c d e f g h i j k l m n o
	type testCase struct {
		text           string
		encrypted_text string
	}

	testCases := []testCase{
		testCase{"brno", ")gcd"},
		testCase{"ahoj", "!&d<"},
		testCase{"a", "!"},
		testCase{"b", ")"},
		testCase{"c", "\""},
		testCase{"d", "("},
		testCase{"e", "£"},
		testCase{"ab", "!)"},
	}

	for _, tt := range testCases {
		if got := Encrypt(tt.text); got != tt.encrypted_text {
			t.Errorf("Encrypt(%q): %q, WANT: %q", tt.text, got, tt.encrypted_text)
		}
	}
}

// func TestEncryptOriginal(t *testing.T) {
// 	// a b c d e f g h i j k l m n o p q r s t u v w x y z
// 	// ! ) " ( £ * % & > < @ a b c d e f g h i j k l m n o
// 	var text string
// 	var encrypted_text string
//
// 	text = "brno"
// 	encrypted_text = ")gcd"
// 	if got := Encrypt(text); got != encrypted_text {
// 		t.Errorf("Encrypt(%q): %q, WANT: %q", text, got, encrypted_text)
// 	}
//
// 	text = "ahoj"
// 	encrypted_text = "!&d<"
// 	if got := Encrypt(text); got != encrypted_text {
// 		t.Errorf("Encrypt(%q): %q, WANT: %q", text, got, encrypted_text)
// 	}
//
// 	text = "a"
// 	encrypted_text = "!"
// 	if got := Encrypt(text); got != encrypted_text {
// 		t.Errorf("Encrypt(%q): %q, WANT: %q", text, got, encrypted_text)
// 	}
// }
