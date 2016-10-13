package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestHexWriter(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"Hello, world!", "48656c6c6f2c20776f726c6421\n"},
		{"Hello", "48656c6c6f\n"},
		{"Hello\n", "48656c6c6f0a\n"},
		// not ready to do this yet...
		// {"Ahoj, světe!", ".."},
		// {"大家好！", "e5 a4 a7 e5 ae b6 e5 a5 bd ef bc 81\n"},
	}
	for _, tt := range tests {
		var b bytes.Buffer
		h := NewHexWriter(&b)
		io.Copy(h, strings.NewReader(tt.in))
		if got := b.String(); got != tt.want {
			t.Errorf("HexWriter.Write(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestHex(t *testing.T) {
	tests := []struct {
		in   []byte
		want []byte
	}{
		{},
		{[]byte("H"), []byte("48")},
		{[]byte("Hello"), []byte("48656c6c6f")},
		{[]byte(","), []byte("20")},
		{[]byte("!"), []byte("21")},
		{[]byte("\n"), []byte("0a")},
	}
	for _, tt := range tests {
		got := hex(tt.in)
		if got := b.String(); got != tt.want {
			t.Errorf("ToHex(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestToHex(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"Hello, world!", "48656c6c6f2c20776f726c6421\n"},
		{"Hello", "48656c6c6f\n"},
		{"Hello\n", "48656c6c6f0a\n"},
		// not ready to do this yet...
		// {"Ahoj, světe!", ".."},
		// {"大家好！", "e5 a4 a7 e5 ae b6 e5 a5 bd ef bc 81\n"},
	}
	for _, tt := range tests {
		var b bytes.Buffer
		ToHex(&b, strings.NewReader(tt.in))
		if got := b.String(); got != tt.want {
			t.Errorf("ToHex(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
