package cidme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDProvider(t *testing.T) {
	t.Run("with no options", func(t *testing.T) {
		p := New()
		assert.NotNil(t, p)
		assert.Equal(t, 0, p.length)
		assert.Equal(t, "", p.padChar)
		assert.Equal(t, false, p.normalized)
		id, err := p.ToID("Test1234 ")
		assert.NoError(t, err)
		assert.Equal(t, "Test1234 ", id)
	})
	t.Run("with normalized option", func(t *testing.T) {
		p := New(WithNormalized())
		assert.NotNil(t, p)
		assert.Equal(t, 0, p.length)
		assert.Equal(t, "", p.padChar)
		assert.Equal(t, true, p.normalized)
		id, err := p.ToID("TesT1234")
		assert.NoError(t, err)
		assert.Equal(t, "test1234", id)
	})
	t.Run("with padded option", func(t *testing.T) {
		p := New(WithPadding("0", 10))
		assert.NotNil(t, p)
		assert.Equal(t, 10, p.length)
		assert.Equal(t, "0", p.padChar)
		assert.Equal(t, false, p.normalized)
		id, err := p.ToID("test1234")
		assert.NoError(t, err)
		assert.Equal(t, "test123400", id)
	})
	t.Run("with padded option short", func(t *testing.T) {
		p := New(WithPadding("0", 5))
		assert.NotNil(t, p)
		assert.Equal(t, 5, p.length)
		assert.Equal(t, "0", p.padChar)
		assert.Equal(t, false, p.normalized)
		id, err := p.ToID("Test1234")
		assert.NoError(t, err)
		assert.Equal(t, "Test1", id)
	})
	t.Run("with multiple option", func(t *testing.T) {
		p := New(WithNormalized(), WithPadding("1", 10))
		assert.NotNil(t, p)
		assert.Equal(t, 10, p.length)
		assert.Equal(t, "1", p.padChar)
		assert.Equal(t, true, p.normalized)
		id, err := p.ToID("TesT1234")
		assert.NoError(t, err)
		assert.Equal(t, "test123411", id)
	})
	t.Run("with multiple option short", func(t *testing.T) {
		p := New(WithNormalized(), WithPadding("1", 5))
		assert.NotNil(t, p)
		assert.Equal(t, 5, p.length)
		assert.Equal(t, "1", p.padChar)
		assert.Equal(t, true, p.normalized)
		id, err := p.ToID("TesT1234")
		assert.NoError(t, err)
		assert.Equal(t, "test1", id)
	})
}
