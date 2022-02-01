package idf

import "github.com/mchmarny/idf/internal/formatter"

type IDFormatter struct {
	formatters []formatter.Formatter
}

// ToID converts the given string to an ID using the configured options.
func (p *IDFormatter) ToID(v string) (string, error) {
	var err error
	for _, f := range p.formatters {
		v, err = f.Format(v)
		if err != nil {
			return "", err
		}
	}
	return v, nil
}

// New creates a new IDFormatter with given options.
func New(opts ...func(*IDFormatter)) *IDFormatter {
	f := &IDFormatter{
		formatters: []formatter.Formatter{},
	}
	for _, o := range opts {
		o(f)
	}
	return f
}
