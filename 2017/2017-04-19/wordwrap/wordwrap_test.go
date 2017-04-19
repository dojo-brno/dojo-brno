package wordwrap

import "testing"

func TestWordWrap(t *testing.T) {
	tests := []struct {
		input  string
		column int
		want   string
	}{
		{"", 5, ""},
		{"hokuspokus", 5, "hokus\npokus"},
		{"hokushokuspokus", 5, "hokus\nhokus\npokus"},
		{"ho kus ", 5, "ho\nkus "},
		{"sem tam sem", 5, "sem\ntam\nsem"},
		{"s e m t a m", 5, "s e\nm t\na m"},
		{"s e m t a m", 9, "s e m t\na m"},
		{"hokusy pokusy", 5, "hokus\ny\npokus\ny"},
		{"hokus pokus", 5, "hokus\n\npokus"},
		{"       ", 5, "    \n  "},
		{"h\tkus pokus", 5, "h\tkus\n\npokus"},
	}
	for _, tt := range tests {
		if got := WordWrap(tt.input, tt.column); got != tt.want {
			t.Errorf("WordWrap(%q, %v) = %q want %q", tt.input, tt.column, got, tt.want)
		}
	}
}
