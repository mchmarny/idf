package idf

import (
	"fmt"
)

// WithPrefix prepends the prefix to the input string.
func WithSuffix(suffix string) func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.Formatters = append(f.Formatters, &SuffixFormatter{
			suffix: suffix,
		})
	}
}

type SuffixFormatter struct {
	suffix string
}

func (f *SuffixFormatter) Format(v string) (string, error) {
	return fmt.Sprintf("%s%s", v, f.suffix), nil
}
