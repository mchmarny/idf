package main

import (
	"fmt"

	"github.com/mchmarny/idf/id"
)

func main() {
	f := id.New(
		// Base64 encode the ID
		id.WithBase64Encoding(),
		// And, prefix it with "id-"
		id.WithPrefix("id-"),
	)
	s := "qw&^*($#@!~"
	id, err := f.ToID(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("new id  : %s\n", id)
	fmt.Printf("original: %s\n", s)
}
