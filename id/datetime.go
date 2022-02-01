package id

import (
	"fmt"
	"time"
)

// WithDatetime formats the time value and appends it to the id.
func WithDatetime(format, separator string, utc bool) func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.formatters = append(f.formatters, &DatetimeFormatter{
			format:    format,
			separator: separator,
			utc:       utc,
		})
	}
}

type DatetimeFormatter struct {
	separator string
	format    string
	utc       bool
}

func (f *DatetimeFormatter) Format(v string) (string, error) {
	now := time.Now()
	if f.utc {
		now = now.UTC()
	}
	return fmt.Sprintf("%s%s%s", v, f.separator, now.Format(f.format)), nil
}
