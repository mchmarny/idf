package main

import (
	"fmt"

	"github.com/mchmarny/idf/id"
)

func main() {
	f := id.New(
		// SHA246 encode the ID resulting in 64 char lne string
		id.WithSHA256Encoding(),
		// And, prefix it with "id-"
		id.WithPrefix("id-"),
	)
	s1 := "123-^!("
	id1, err := f.ToID(s1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("'%s' from '%s'\n", id1, s1)

	s2 := "1234567891011121314151617181920212223242526-^!("
	id2, err := f.ToID(s2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("'%s' from '%s'\n", id2, s2)
}
