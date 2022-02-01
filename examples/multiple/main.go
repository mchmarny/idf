package main

import (
	"fmt"

	"github.com/mchmarny/idf"
)

func main() {
	f := idf.New(
		// Base64 encode the ID
		idf.WithBase64Encoding(),
		// And, prefix it with "id-"
		idf.WithPrefix("id-"),
	)
	s := "qw&^*($#@!~"
	id, err := f.ToID(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("new id  : %s\n", id)
	fmt.Printf("original: %s\n", s)
}
