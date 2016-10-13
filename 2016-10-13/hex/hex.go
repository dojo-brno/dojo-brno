package main

import "fmt"

func main() {
	s := "Hello, world!"
	for _, i := range s {
		fmt.Printf("%x", i)
	}
	fmt.Println()
}
