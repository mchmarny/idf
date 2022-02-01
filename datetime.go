package idf

import (
	"fmt"
	"time"
)

// WithPrefix prepends the prefix to the input string.
func WithDatetime(format, seperator string, v time.Time) func(*IDFormater) {
	return func(f *IDFormater) {
		f.formaters = append(f.formaters, &DatetimeFormater{
			format:    format,
			seperator: seperator,
			v:         v,
		})
	}
}

type DatetimeFormater struct {
	seperator string
	format    string
	v         time.Time
}

func (f *DatetimeFormater) Format(v string) (string, error) {
	return fmt.Sprintf("%s%s%s", v, f.seperator, f.v.Format(f.format)), nil
}
