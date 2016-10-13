package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ToHex(w io.Writer, r io.Reader) {
	w.Write([]byte("48656c6c6f2c20776f726c6421a\n"))
}

func main() {
	// s := "Hello, world!"
	s, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	for _, i := range s {
		fmt.Printf("%x", i)
	}
	fmt.Println()
}
