package main

import (
	"fmt"

	"github.com/mchmarny/idf/id"
)

func main() {
	f := id.New(
		// Normalizes the ID by removing spaces and converting to lowercase
		id.WithNormalizing(),
	)
	s := "TesT1234"
	id, err := f.ToID(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("new id  : %s\n", id)
	fmt.Printf("original: %s\n", s)
}
