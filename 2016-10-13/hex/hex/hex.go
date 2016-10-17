package hex

import (
	"fmt"
	"io"
)

// A Writer is an io.Writer.
// Writes to a Writer write the hexdecimal representation of each byte to w.
type Writer struct {
	w io.Writer
}

// NewWriter returns a new Writer.
func NewWriter(w io.Writer) *Writer {
	return &Writer{w}
}

// Write writes the hexdecimal representation of each byte in p to the
// underlying io.Writer.
func (h *Writer) Write(p []byte) (int, error) {
	s := hex(p)
	written, err := h.w.Write(s)
	if written > len(s) {
		panic("hex.Writer.Write: invalid Write count")
	}
	if written != len(s) && err == nil {
		err = io.ErrShortWrite
	}
	return len(p), err
}

func hex(p []byte) []byte {
	return []byte(fmt.Sprintf("%x", p))
}
