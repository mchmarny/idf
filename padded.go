package cidme

import (
	"fmt"
	"strings"
)

const (
	PadCharDefault = "0"
)

// NewPaddedIDProvider creates a validated padded ID provider.
func NewPaddedIDProvider(normalized bool, length int, padChar string) (*PaddedIDProvider, error) {
	if padChar == "" {
		return nil, fmt.Errorf("padding character must be specified")
	}
	if len(padChar) != 1 {
		return nil, fmt.Errorf("padding character must be a single character")
	}
	if length < 1 {
		return nil, fmt.Errorf("length must be greater than 0")
	}

	return &PaddedIDProvider{
		Normalized: normalized,
		Length:     length,
		PadChar:    padChar,
	}, nil
}

// PaddedIDProvider creates IDs based on configured values.
type PaddedIDProvider struct {
	Length     int
	PadChar    string
	Normalized bool
}

// ToID returns a a formatted string ID.
func (p *PaddedIDProvider) ToID(v string) (string, error) {
	if p.Normalized {
		v = normalize(v)
	}

	if len(v) > p.Length {
		return "", fmt.Errorf("length of '%s' is greater than configured length: %d", v, p.Length)
	}
	return fmt.Sprintf("%s%s", v, strings.Repeat(p.PadChar, p.Length-len(v))), nil
}
