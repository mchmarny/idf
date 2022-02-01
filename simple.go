package cidme

import (
	"strings"
)

// NewSipleIDProvider creates a simple ID provider.
func NewSipleIDProvider(normalized bool) *SimpleIDProvider {
	return &SimpleIDProvider{
		Normalized: normalized,
	}
}

// SimpleIDProvider is a simple ID provider.
type SimpleIDProvider struct {
	Normalized bool
}

// ToID returns a a formatted string ID.
func (p *SimpleIDProvider) ToID(v string) (string, error) {
	if p.Normalized {
		return normalize(v), nil
	}
	return v, nil
}

func normalize(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}
