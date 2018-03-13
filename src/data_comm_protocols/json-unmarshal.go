package main

import (
	"encoding/json"
	"fmt"
)

// Animal represents a type of animal and the order it belongs to.
type Animal struct {
	Name  string
	Order string
}

func main() {
	// example json
	var raw = []byte(`[
		{"Name":"Platypus","Order": "Monotremata"},
		{"Name":"Quoll","Order":"Dasyuromorphia"}
		]`)

	var animals []Animal
	err := json.Unmarshal(raw, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, a := range animals {
		fmt.Printf("%#v\n", a)
	}
}
