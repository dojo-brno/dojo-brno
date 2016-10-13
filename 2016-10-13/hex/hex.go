package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func ToHex(w io.Writer, r io.Reader) {
	b := make([]byte, 1)
	if _, err := r.Read(b); err == io.EOF {
		return
	}
	in, _ := ioutil.ReadAll(r)
	for _, i := range append(b, in...) {
		w.Write([]byte(strconv.FormatInt(int64(i), 16)))
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
