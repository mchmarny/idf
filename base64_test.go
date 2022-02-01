package idf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64EncodingFormatter(t *testing.T) {
	t.Run("test resulting id length", func(t *testing.T) {
		f := &Base64EncodingFormatter{}
		assert.NotNil(t, f)
		_, err := f.Format("!@#$&$(*=")
		assert.NoError(t, err)
	})
}
