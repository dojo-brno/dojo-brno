package hex

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestWriter(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"Hello, world!", "48656c6c6f2c20776f726c6421"},
		{"Hello", "48656c6c6f"},
		{"Hello\n", "48656c6c6f0a"},
		// not ready to do this yet...
		// {"Ahoj, světe!", ".."},
		// {"大家好！", "e5 a4 a7 e5 ae b6 e5 a5 bd ef bc 81\n"},
	}
	for _, tt := range tests {
		var b bytes.Buffer
		h := NewWriter(&b)
		written, err := io.Copy(h, strings.NewReader(tt.in))
		if err != nil {
			t.Errorf("io.Copy(h, strings.NewReader(%q)) error: %v", tt.in, err)
		}
		if written != int64(len(tt.in)) {
			t.Errorf("io.Copy(h, strings.NewReader(%q)) wrote %d bytes, want %d", tt.in, written, len(tt.in))
		}
		if got := b.String(); got != tt.want {
			t.Errorf("Writer.Write(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

// failWriter is an io.Writer that always returns errors on write.
type failWriter struct{}

func (failWriter) Write(p []byte) (int, error) {
	return 0, errors.New("failWriter error")
}

func TestFailWriter(t *testing.T) {
	var w failWriter
	h := NewWriter(w)
	_, err := h.Write([]byte("hello"))
	if err == nil {
		t.Errorf("err = %v, want %v", err, nil)
	}
}

// shortWriter is an io.Writer that never returns errors, but writes 0 bytes.
type shortWriter struct{}

func (shortWriter) Write(p []byte) (int, error) {
	return 0, nil
}

func TestShortWriter(t *testing.T) {
	var w shortWriter
	h := NewWriter(w)
	_, err := h.Write([]byte("hello"))
	if err != io.ErrShortWrite {
		t.Errorf("err = %v, want %v", err, io.ErrShortWrite)
	}
}

// invalidWriter is an io.Writer that never returns errors, but writes more
// bytes than given.
type invalidWriter struct{}

func (invalidWriter) Write(p []byte) (int, error) {
	return len(p) + 1, nil
}

func TestInvalidWriter(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("h.Write did not panic!")
		}
	}()
	var w invalidWriter
	h := NewWriter(w)
	h.Write([]byte("hello"))
}

func TestHex(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{}, // intentionally empty
		{"H", "48"},
		{",", "2c"},
		{"!", "21"},
		{"\n", "0a"},
		{"Hello", "48656c6c6f"},
	}
	for _, tt := range tests {
		got := string(hex([]byte(tt.in)))
		if got != tt.want {
			t.Errorf("hex(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
