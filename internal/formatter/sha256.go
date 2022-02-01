package formatter

import (
	"crypto/sha256"
	"fmt"
)

type SHA256EncodingFormatter struct{}

func (f *SHA256EncodingFormatter) Format(v string) (string, error) {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(v))), nil
}
