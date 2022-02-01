package cidme

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaddedID(t *testing.T) {
	t.Run("value to invalid ID length", func(t *testing.T) {
		p, err := NewPaddedIDProvider(true, 3, PadCharDefault)
		assert.NoError(t, err)
		assert.NotNil(t, p)
		_, err = p.ToID("id12345 ")
		assert.Error(t, err)
	})
	t.Run("value to ID length with default char", func(t *testing.T) {
		id1 := "12345"
		idLen := 10
		p, err := NewPaddedIDProvider(true, idLen, PadCharDefault)
		assert.NoError(t, err)
		assert.NotNil(t, p)
		id2, err := p.ToID(id1)
		assert.NoError(t, err)
		assert.Len(t, id2, idLen)
		assert.Equal(t, id2, fmt.Sprintf("%s%s", id1, strings.Repeat(PadCharDefault, idLen-len(id1))))
	})
	t.Run("value to ID length with custom char", func(t *testing.T) {
		id1 := "my-word-is-my-id"
		idLen := 36
		padChar := "*"
		p, err := NewPaddedIDProvider(true, idLen, padChar)
		assert.NoError(t, err)
		assert.NotNil(t, p)
		id2, err := p.ToID(id1)
		assert.NoError(t, err)
		assert.NoError(t, err)
		assert.Len(t, id2, idLen)
		assert.Equal(t, id2, fmt.Sprintf("%s%s", id1, strings.Repeat(padChar, idLen-len(id1))))
	})
}
