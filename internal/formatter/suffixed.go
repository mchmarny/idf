package formatter

import (
	"fmt"
)

type SuffixFormatter struct {
	Suffix string
}

func (f *SuffixFormatter) Format(v string) (string, error) {
	return fmt.Sprintf("%s%s", v, f.Suffix), nil
}
