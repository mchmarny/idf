package idf

import (
	"fmt"
)

// WithPrefix prepends the prefix to the input string.
func WithPrefix(prefix string) func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.Formatters = append(f.Formatters, &PrefixFormatter{
			prefix: prefix,
		})
	}
}

type PrefixFormatter struct {
	prefix string
}

func (f *PrefixFormatter) Format(v string) (string, error) {
	return fmt.Sprintf("%s%s", f.prefix, v), nil
}
