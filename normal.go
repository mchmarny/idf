package idf

import "strings"

// WithNormalized will normalize the input string to lowercase and trim spaces.
func WithNormalized() func(*IDFormater) {
	return func(f *IDFormater) {
		f.formaters = append(f.formaters, &NormalizedFormater{
			normalized: true,
		})
	}
}

type NormalizedFormater struct {
	normalized bool
}

func (f *NormalizedFormater) Format(v string) (string, error) {
	if f.normalized {
		v = strings.TrimSpace(strings.ToLower(v))
	}
	return v, nil
}
