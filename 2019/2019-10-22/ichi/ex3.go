package main

import (
	"fmt"
)

type StringSet struct {
	items map[string]int
}

func NewStringSet(items ...string) *StringSet {
	return make(map[string]StringSet)
}

func (s *StringSet) Contains(item string) bool {
	if val, ok := map[item]; ok {
		return true
	}
	return false
}

func main() {
	s := NewStringSet("cat", "dog", "cow")
	fmt.Println("contains cat:", s.Contains("cat"))
	fmt.Println("contains car:", s.Contains("car"))
}
