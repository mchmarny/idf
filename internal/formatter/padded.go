package formatter

import (
	"fmt"
	"strings"
)

const (
	defaultPadChar = "0"
)

type PaddingFormatter struct {
	Length  int
	PadChar string
}

func (f *PaddingFormatter) Format(v string) (string, error) {
	if f.Length == 0 {
		return v, nil
	}

	if f.Length > 0 && len(v) > f.Length {
		return "", fmt.Errorf("input string is longer than the target length")
	}

	if f.PadChar == "" {
		f.PadChar = defaultPadChar
	}

	return fmt.Sprintf("%s%s", v, strings.Repeat(f.PadChar, f.Length-len(v))), nil
}
