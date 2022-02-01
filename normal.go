package idf

import "strings"

// WithNormalizing normalizes the input string to lowercase and trim spaces.
func WithNormalizing() func(*IDFormater) {
	return func(f *IDFormater) {
		f.formaters = append(f.formaters, &NormalizingFormater{
			normalized: true,
		})
	}
}

type NormalizingFormater struct {
	normalized bool
}

func (f *NormalizingFormater) Format(v string) (string, error) {
	if f.normalized {
		v = strings.TrimSpace(strings.ToLower(v))
	}
	return v, nil
}
