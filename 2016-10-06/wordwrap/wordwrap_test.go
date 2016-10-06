package wordwrap

import "testing"

func TestWordwrap(t *testing.T) {
	tests := []struct {
		columns int
		text    string
		want    string
	}{
		{0, "", ""},
		{1, "aa", "a\na"},
		{2, "aa", "aa"},
		{2, "aaaaa", "aa\naa\na"},
		{1, "bb", "b\nb"},
		{2, "bb", "bb"},
		{1, "cc", "c\nc"},
		{3, "aa", "aa"},
		{2, "aaa", "aa\na"},
	}
	for _, tt := range tests {
		columns := tt.columns
		text := tt.text
		want := tt.want
		wrappedText := wordwrap(columns, text)
		if want != wrappedText {
			t.Errorf("Columns: %v, text: %v, result: %q, want: %q", columns, text, wrappedText, want)
		}
	}
}
