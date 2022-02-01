package id

import (
	"encoding/base64"
)

// WithBase64Encoding base64 with URL encodes the input string.
func WithBase64Encoding() func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.formatters = append(f.formatters, &Base64EncodingFormatter{})
	}
}

type Base64EncodingFormatter struct{}

func (f *Base64EncodingFormatter) Format(v string) (string, error) {
	return base64.URLEncoding.EncodeToString([]byte(v)), nil
}
