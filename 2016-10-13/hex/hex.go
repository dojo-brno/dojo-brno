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
	// s := "Hello, world!"
	s, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	for _, i := range s {
		fmt.Printf("%02x", i)
	}
	fmt.Println()
}
