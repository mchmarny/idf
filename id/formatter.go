package id

type IDFormatter struct {
	formatters []Formatter
}

type Formatter interface {
	Format(v string) (string, error)
}

// New creates a new IDFormatter with given options.
func New(opts ...func(*IDFormatter)) *IDFormatter {
	f := &IDFormatter{
		formatters: []Formatter{},
	}
	for _, o := range opts {
		o(f)
	}
	return f
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
