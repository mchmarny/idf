package idf

import "encoding/base64"

// WithBase64Encoding base64 with URL encodes the input string.
func WithBase64Encoding() func(*IDFormater) {
	return func(f *IDFormater) {
		f.formaters = append(f.formaters, &Base64EncodingFormater{})
	}
}

type Base64EncodingFormater struct{}

func (f *Base64EncodingFormater) Format(v string) (string, error) {
	return base64.URLEncoding.EncodeToString([]byte(v)), nil
}
