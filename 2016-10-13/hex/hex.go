package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type HexWriter struct {
	w io.Writer
}

func NewHexWriter(w io.Writer) *HexWriter {
	return &HexWriter{w}
}

func (h *HexWriter) Write(p []byte) (int, error) {
	return h.w.Write(hex(p))
}

func hex(p []byte) []byte {
	s := make([]byte, 0, 2*len(p))
	for i := range p {
		s = append(s, []byte(fmt.Sprintf("%02x", i))...)
	}
	return s
}

func ToHex(w io.Writer, r io.Reader) {
	b := make([]byte, 1)
	if _, err := r.Read(b); err == io.EOF {
		return
	}
	in, _ := ioutil.ReadAll(r)
	for _, i := range append(b, in...) {
		w.Write([]byte(fmt.Sprintf("%02x", i)))
	}
	w.Write([]byte{'\n'})
}

func main() {
	ToHex(os.Stdout, os.Stdin)
}
