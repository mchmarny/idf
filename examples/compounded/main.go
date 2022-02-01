package main

import (
	"fmt"

	"github.com/mchmarny/idf/id"
)

func main() {
	f := id.New(
		// Normalizes the ID by removing spaces and converting to lowercase
		id.WithNormalizing(),
		// Padded so that all IDs are of the same length
		id.WithPadding("^", 20),
		// SHA246 encode the ID resulting in 64 char lne string
		id.WithSHA256Encoding(),
		// And, prefix it with "id-"
		id.WithPrefix("id-"),
		// Endign with a date to allow for predictable sortability
		id.WithDatetime("2006-01-02", "-", true),
	)

	ids := []string{
		"usr-1",
		"usr-2",
		"usr-3",
		"usr-4",
		"usr-5",
	}

	for _, id := range ids {
		cid, err := f.ToID(id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s from %s\n", cid, id)
	}
}
