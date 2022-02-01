package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaddedFormatter(t *testing.T) {
	t.Run("test empty padded Formatter", func(t *testing.T) {
		f := &PaddingFormatter{}
		assert.NotNil(t, f)
		_, err := f.Format("id12345")
		assert.NoError(t, err)
	})
	t.Run("test long padded Formatter", func(t *testing.T) {
		f := &PaddingFormatter{
			Length: 10,
		}
		assert.NotNil(t, f)
		id, err := f.Format("id12345")
		assert.NoError(t, err)
		assert.Len(t, id, 10)
		assert.Equal(t, "id12345000", id)
	})
	t.Run("test length padded Formatter", func(t *testing.T) {
		f := &PaddingFormatter{
			PadChar: "0",
			Length:  10,
		}
		assert.NotNil(t, f)
		id, err := f.Format("id12345")
		assert.NoError(t, err)
		assert.Len(t, id, 10)
	})
	t.Run("test too long padded Formatter", func(t *testing.T) {
		f := &PaddingFormatter{
			PadChar: "0",
			Length:  5,
		}
		assert.NotNil(t, f)
		_, err := f.Format("id12345234324")
		assert.Error(t, err)
	})
}
