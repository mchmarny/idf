package idf

import (
	"fmt"
	"time"
)

// WithPrefix prepends the prefix to the input string.
func WithDatetime(format, seperator string, utc bool) func(*IDFormater) {
	return func(f *IDFormater) {
		f.formaters = append(f.formaters, &DatetimeFormater{
			format:    format,
			seperator: seperator,
			utc:       utc,
		})
	}
}

type DatetimeFormater struct {
	seperator string
	format    string
	utc       bool
}

func (f *DatetimeFormater) Format(v string) (string, error) {
	now := time.Now()
	if f.utc {
		now = now.UTC()
	}

	return fmt.Sprintf("%s%s%s", v, f.seperator, now.Format(f.format)), nil
}
