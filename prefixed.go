package idf

import (
	"fmt"
)

// WithPrefix prepends the prefix to the input string.
func WithPrefix(prefix string) func(*IDFormater) {
	return func(f *IDFormater) {
		f.formaters = append(f.formaters, &PrefixFormater{
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
