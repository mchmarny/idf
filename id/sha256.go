package id

import (
	"crypto/sha256"
	"fmt"
)

// WithSHA256Encoding sha256 encodes the input string.
func WithSHA256Encoding() func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.formatters = append(f.formatters, &SHA256EncodingFormatter{})
	}
}

type SHA256EncodingFormatter struct{}

func (f *SHA256EncodingFormatter) Format(v string) (string, error) {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(v))), nil
}