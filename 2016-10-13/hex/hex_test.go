package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestToHex(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"Hello, world!", "48656c6c6f2c20776f726c6421\n"},
		// not ready to do this yet...
		// {"大家好！", "...."},
		{"Hello", "48656c6c6f\n"},
		{"Hello\n", "48656c6c6f0a\n"},
	}
	for _, tt := range tests {
		var b bytes.Buffer
		ToHex(&b, strings.NewReader(tt.in))
		if got := b.String(); got != tt.want {
			t.Errorf("ToHex(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}

}
