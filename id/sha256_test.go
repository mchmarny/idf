package id

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256EncodingFormatter(t *testing.T) {
	t.Run("test resulting id length", func(t *testing.T) {
		f := &SHA256EncodingFormatter{}
		assert.NotNil(t, f)
		id1, err := f.Format("!@#$&$(*=")
		assert.NoError(t, err)
		assert.Len(t, id1, 64)
		id2, err := f.Format("!@#$&$(*=")
		assert.NoError(t, err)
		assert.Equal(t, id1, id2)
		id3, err := f.Format(`!@#dssdfs345646dgfghj7897546
		                      4564232dfddffdgdf78123455545
							  646457567567567567tyuytutfhg
							  fhfgdfeterwtyjtjul;ukjlu;ktr
							  asdasdsdfsd53453sfrsfsdshjsh`)
		assert.NoError(t, err)
		assert.Len(t, id3, 64)
	})
}
