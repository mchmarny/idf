package formatter

import "strings"

type NormalizingFormatter struct{}

func (f *NormalizingFormatter) Format(v string) (string, error) {
	return strings.TrimSpace(strings.ToLower(v)), nil
}
