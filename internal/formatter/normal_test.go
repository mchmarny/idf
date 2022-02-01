package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizingFormatter(t *testing.T) {
	t.Run("test normalizing Formatter", func(t *testing.T) {
		f := &NormalizingFormatter{}
		assert.NotNil(t, f)
		id, err := f.Format("Id12345 ")
		assert.NoError(t, err)
		assert.Equal(t, "id12345", id)
	})
}
