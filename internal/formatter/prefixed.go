package formatter

import (
	"fmt"
)

type PrefixFormater struct {
	Prefix string
}

func (f *PrefixFormater) Format(v string) (string, error) {
	return fmt.Sprintf("%s%s", f.Prefix, v), nil
}
