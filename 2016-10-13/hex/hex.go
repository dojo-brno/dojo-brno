package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

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
