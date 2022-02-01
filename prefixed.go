package idf

import (
	"fmt"
)

// WithPrefix normalizes the input string to lowercase and trim spaces.
func WithPrefix(prefix string) func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.formatters = append(f.formatters, &PrefixFormater{
			prefix: prefix,
		})
	}
}

type PrefixFormater struct {
	prefix string
}

func (f *PrefixFormater) Format(v string) (string, error) {
	return fmt.Sprintf("%s%s", f.prefix, v), nil
}
