package idf

import (
	"fmt"
)

// WithPrefix prepends the prefix to the input string.
func WithSuffix(suffix string) func(*IDFormater) {
	return func(f *IDFormater) {
		f.formaters = append(f.formaters, &SuffixFormater{
			suffix: suffix,
		})
	}
}

type SuffixFormater struct {
	suffix string
}

func (f *SuffixFormater) Format(v string) (string, error) {
	return fmt.Sprintf("%s%s", v, f.suffix), nil
}
