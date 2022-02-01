package idf

import (
	"fmt"
	"time"
)

// WithDatetime formats the time value and appends it to the id.
func WithDatetime(format, separator string, v time.Time) func(*IDFormatter) {
	return func(f *IDFormatter) {
		f.Formatters = append(f.Formatters, &DatetimeFormatter{
			format:    format,
			separator: separator,
			v:         v,
		})
	}
}

type DatetimeFormatter struct {
	separator string
	format    string
	v         time.Time
}

func (f *DatetimeFormatter) Format(v string) (string, error) {
	return fmt.Sprintf("%s%s%s", v, f.separator, f.v.Format(f.format)), nil
}
