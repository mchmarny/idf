package idf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDFormater(t *testing.T) {
	t.Run("with no options", func(t *testing.T) {
		p := New()
		assert.NotNil(t, p)
		id, err := p.ToID("Test1234 ")
		assert.NoError(t, err)
		assert.Equal(t, "Test1234 ", id)
	})
	t.Run("with normalized option", func(t *testing.T) {
		p := New(WithNormalized())
		assert.NotNil(t, p)
		id, err := p.ToID("TesT1234")
		assert.NoError(t, err)
		assert.Equal(t, "test1234", id)
	})
	t.Run("with padded option", func(t *testing.T) {
		p := New(WithPadding("0", 10))
		assert.NotNil(t, p)
		id, err := p.ToID("test1234")
		assert.NoError(t, err)
		assert.Equal(t, "test123400", id)
	})
	t.Run("with padded option short", func(t *testing.T) {
		p := New(WithPadding("0", 5))
		assert.NotNil(t, p)
		_, err := p.ToID("Test1234")
		assert.Error(t, err)
	})
	t.Run("with multiple option", func(t *testing.T) {
		p := New(WithNormalized(), WithPadding("1", 10))
		assert.NotNil(t, p)
		id, err := p.ToID("TesT1234")
		assert.NoError(t, err)
		assert.Equal(t, "test123411", id)
	})
	t.Run("with multiple option short", func(t *testing.T) {
		p := New(WithNormalized(), WithPadding("1", 5))
		assert.NotNil(t, p)
		_, err := p.ToID("TesT1234")
		assert.Error(t, err)
	})
}