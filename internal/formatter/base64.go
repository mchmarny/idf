package formatter

import (
	"encoding/base64"
)

type Base64EncodingFormatter struct{}

func (f *Base64EncodingFormatter) Format(v string) (string, error) {
	return base64.URLEncoding.EncodeToString([]byte(v)), nil
}
