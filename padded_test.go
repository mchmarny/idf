package idf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaddedFormater(t *testing.T) {
	t.Run("test empty padded formater", func(t *testing.T) {
		f := &PaddingFormater{}
		assert.NotNil(t, f)
		_, err := f.Format("id12345")
		assert.NoError(t, err)
	})
}
