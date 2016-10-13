package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

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
