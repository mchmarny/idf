package idf

import (
	"fmt"
	"strings"
)

// WithPadding will pad the input string with the given character to the given length.
func WithPadding(char string, length int) func(*IDFormater) {
	return func(f *IDFormater) {
		f.formaters = append(f.formaters, &PaddingFormater{
			padChar: char,
			length:  length,
		})
	}
}

type PaddingFormater struct {
	length  int
	padChar string
}

func (f *PaddingFormater) Format(v string) (string, error) {
	if f.length == 0 {
		return v, nil
	}

	if f.length > 0 && len(v) > f.length {
		return "", fmt.Errorf("input string is longer than the target length")
	}

	return fmt.Sprintf("%s%s", v, strings.Repeat(f.padChar, f.length-len(v))), nil
}
