package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixedFormatter(t *testing.T) {
	t.Run("test empty padded Formatter", func(t *testing.T) {
		f := &PrefixFormater{}
		assert.NotNil(t, f)
		id, err := f.Format("id12345")
		assert.NoError(t, err)
		assert.Equal(t, "id12345", id)
	})
	t.Run("test empty padded Formatter", func(t *testing.T) {
		f := &PrefixFormater{
			Prefix: "id-",
		}
		assert.NotNil(t, f)
		id, err := f.Format("id12345")
		assert.NoError(t, err)
		assert.Equal(t, "id-id12345", id)
	})
}
