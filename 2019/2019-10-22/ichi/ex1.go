package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cat struct {
	Name string
}

func marshal() {
	c := &Cat{
		Name: "Lisa",
	}

	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}

func unmarshal() {
	b := []byte(`{"name":"Lisa"}`)
	c := &Cat{}

	if err := json.UnmarshalJSON(b, c); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(c)
}

func main() {
	marshal()
	unmarshal()
}
