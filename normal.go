package idf

import "strings"

// WithNormalizing normalizes the input string to lowercase and trim spaces.
func WithNormalizing() func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.Formatters = append(f.Formatters, &NormalizingFormatter{
			normalized: true,
		})
	}
}

type NormalizingFormatter struct {
	normalized bool
}

func (f *NormalizingFormatter) Format(v string) (string, error) {
	if f.normalized {
		v = strings.TrimSpace(strings.ToLower(v))
	}
	return v, nil
}
