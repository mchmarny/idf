package idf

import (
	"crypto/sha256"
	"fmt"
)

// WithSHA256Encoding sha256 encodes the input string.
func WithSHA256Encoding() func(*IDFormater) {
	return func(f *IDFormater) {
		f.formaters = append(f.formaters, &SHA256EncodingFormater{})
	}
}

type SHA256EncodingFormater struct{}

func (f *SHA256EncodingFormater) Format(v string) (string, error) {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(v))), nil
}
