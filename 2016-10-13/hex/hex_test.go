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
		{"Hello, world!", "48656c6c6f2c20776f726c6421a"},
	}
	for _, tt := range tests {
		var b bytes.Buffer
		ToHex(&b, strings.NewReader(tt.in))
		if got := b.String(); got != tt.want {
			t.Errorf("ToHex(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}

}
