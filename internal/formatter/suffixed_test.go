package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuffixedFormatter(t *testing.T) {
	t.Run("test empty padded Formatter", func(t *testing.T) {
		f := &SuffixFormatter{}
		assert.NotNil(t, f)
		id, err := f.Format("id12345")
		assert.NoError(t, err)
		assert.Equal(t, "id12345", id)
	})
	t.Run("test empty padded Formatter", func(t *testing.T) {
		f := &SuffixFormatter{
			Suffix: "-test",
		}
		assert.NotNil(t, f)
		id, err := f.Format("id12345")
		assert.NoError(t, err)
		assert.Equal(t, "id12345-test", id)
	})
}
