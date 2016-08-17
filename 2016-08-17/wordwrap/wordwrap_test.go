package wordwrap

import "testing"

func TestWordWrap(t *testing.T) {
	tests := []struct {
		text    string
		columns int
		want    string
	}{
		{},
		{
			text:    "a",
			columns: 2,
			want:    "a",
		},
		{
			text:    "aa",
			columns: 1,
			want:    "a\na",
		},
		{
			text:    "aaa",
			columns: 1,
			want:    "a\na\na",
		},
		{
			text:    "aaaa",
			columns: 1,
			want:    "a\na\na\na",
		},
		{
			text:    "aaaa",
			columns: 2,
			want:    "aa\naa",
		},
		{
			text:    "aaaa",
			columns: 3,
			want:    "aaa\na",
		},
		{
			text:    "aaaa",
			columns: 4,
			want:    "aaaa",
		},
		{
			text:    "aaa",
			columns: 4,
			want:    "aaa",
		},
		{
			text:    "bbb",
			columns: 2,
			want:    "bb\nb",
		},
		{
			text:    "ccccccc",
			columns: 2,
			want:    "cc\ncc\ncc\nc",
		},
		{
			text:    "ccccc",
			columns: 2,
			want:    "cc\ncc\nc",
		},
	}
	for _, tt := range tests {
		if got := WordWrap(tt.text, tt.columns); got != tt.want {
			t.Errorf("WordWrap(%q, %d) = %q, want %q", tt.text, tt.columns, got, tt.want)
		}
	}
}
