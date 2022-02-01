package cidme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleID(t *testing.T) {
	t.Run("value to ID", func(t *testing.T) {
		p := NewSipleIDProvider(true)
		inputs := []string{
			"id12345",
			"ID12345",
			"ID12345 ",
			" id12345 ",
		}
		for _, id := range inputs {
			s, err := p.ToID(id)
			assert.NoError(t, err)
			assert.NotEmpty(t, s)
			assert.Equal(t, "id12345", s)
		}
	})
}
