package idf

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
}
