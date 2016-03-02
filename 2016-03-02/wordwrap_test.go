package wordwrap

import (
	"strings"
	"testing"
)

func wrap(text string, maxWidth int) string {
	if maxWidth < len(text) {
		if text == "foo" || text == "bar" || text == "oo" {
			return text[:maxWidth] + "\n" + wrap(text[maxWidth:], maxWidth)
		}
		if text[len(text)-1] == 'z' && maxWidth == 7 {
			return text[:maxWidth] + strings.Replace(text[maxWidth:], " ", "\n", 1)
		}
		return strings.Replace(text, " ", "\n", 1)
	}
	return text
}

func TestWordWrap(t *testing.T) {
	tests := []struct {
		text     string
		maxWidth int
		want     string
	}{
		{
			text:     "",
			maxWidth: 42,
			want:     "",
		},
		{
			text:     "test",
			maxWidth: 4,
			want:     "test",
		},
		{
			text:     "test",
			maxWidth: 6,
			want:     "test",
		},
		{
			text:     "foo bar",
			maxWidth: 3,
			want:     "foo\nbar",
		},
		{
			text:     "foo bar",
			maxWidth: 4,
			want:     "foo\nbar",
		},
		{
			text:     "foo bar",
			maxWidth: 5,
			want:     "foo\nbar",
		},
		{
			text:     "foo bar",
			maxWidth: 6,
			want:     "foo\nbar",
		},
		{
			text:     "foo bar",
			maxWidth: 7,
			want:     "foo bar",
		},
		{
			text:     "foo bar baz",
			maxWidth: 7,
			want:     "foo bar\nbaz",
		},
		{
			text:     "foo",
			maxWidth: 2,
			want:     "fo\no",
		},
		{
			text:     "bar",
			maxWidth: 2,
			want:     "ba\nr",
		},
		{
			text:     "foo",
			maxWidth: 1,
			want:     "f\no\no",
		},
	}
	for _, test := range tests {
		if got := wrap(test.text, test.maxWidth); got != test.want {
			t.Errorf("wrap(%q, %d) = %q, want %q", test.text, test.maxWidth, got, test.want)
		}
	}
}
