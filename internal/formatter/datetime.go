package formatter

import (
	"fmt"
	"time"
)

type DatetimeFormatter struct {
	Separator  string
	BaseFormat string
	UTC        bool
}

func (f *DatetimeFormatter) Format(v string) (string, error) {
	now := time.Now()
	if f.UTC {
		now = now.UTC()
	}
	return fmt.Sprintf("%s%s%s", v, f.Separator, now.Format(f.BaseFormat)), nil
}
