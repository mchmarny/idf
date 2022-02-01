package formatter

type Formatter interface {
	Format(v string) (string, error)
}
