package idf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64EncodingFormater(t *testing.T) {
	t.Run("test resulting id length", func(t *testing.T) {
		f := &Base64EncodingFormater{}
		assert.NotNil(t, f)
		_, err := f.Format("!@#$&$(*=")
		assert.NoError(t, err)
	})
}
