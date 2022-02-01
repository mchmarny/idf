package idf

type IDFormater struct {
	formaters []Formater
}

type Formater interface {
	Format(v string) (string, error)
}

// New creates a new IDFormater with given options.
func New(opts ...func(*IDFormater)) *IDFormater {
	f := &IDFormater{
		formaters: []Formater{},
	}
	for _, o := range opts {
		o(f)
	}
	return f
}

// ToID converts the given string to an ID using the configured options.
func (p *IDFormater) ToID(v string) (string, error) {
	var err error
	for _, f := range p.formaters {
		v, err = f.Format(v)
		if err != nil {
			return "", err
		}
	}
	return v, nil
}
