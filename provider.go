package cidme

type IDProvider interface {
	ToID(v string) (string, error)
}
