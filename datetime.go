package idf

import (
	"fmt"
	"time"
)

// WithDatetime formats the time value and appends it to the id.
func WithDatetime(format, separator string, v time.Time) func(*IDFormater) {
	return func(f *IDFormater) {
		f.formaters = append(f.formaters, &DatetimeFormater{
			format:    format,
			separator: separator,
			v:         v,
		})
	}
}

type DatetimeFormater struct {
	separator string
	format    string
	v         time.Time
}

func (f *DatetimeFormater) Format(v string) (string, error) {
	return fmt.Sprintf("%s%s%s", v, f.separator, f.v.Format(f.format)), nil
}
