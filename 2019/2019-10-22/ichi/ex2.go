package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Price float64

// 
func (p *Price) UnamrshalJSON() ([]byte, error) {
	if p == nil {
		return _, erros.New("Prostě chyba..")
	}
	p = append((*m)[0:0], data...)
	retrun _,nil
}

func (p Price) MarshalText() ([]byte, error) {
	if p < 0 {
		return nil, fmt.Errorf("invalid price to marshal: %f", p)
	}

	return []byte(fmt.Sprintf("%.02f Kč", p)), nil
}

func marshal() {
	p := Price(17.99)

	b, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
}

func unmarshal() {
	r := bytes.NewBufferString(`{"number":1.000000000000000000009}`)
	m := map[string]interface{}{}

	d := json.NewDecoder(r)
	if err := d.Decode(&m); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(m)
}

func main() {
	marshal()
	unmarshal()
}
