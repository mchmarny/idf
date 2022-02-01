package idf

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIDFormatter(t *testing.T) {
	t.Run("test with no options", func(t *testing.T) {
		f := New()
		assert.NotNil(t, f)
		id, err := f.ToID("Test1234 ")
		assert.NoError(t, err)
		assert.Equal(t, "Test1234 ", id)
	})
	t.Run("test with normalized option", func(t *testing.T) {
		f := New(WithNormalizing())
		assert.NotNil(t, f)
		id, err := f.ToID("TesT1234")
		assert.NoError(t, err)
		assert.Equal(t, "test1234", id)
	})
	t.Run("test with padded option", func(t *testing.T) {
		f := New(WithPadding("0", 10))
		assert.NotNil(t, f)
		id, err := f.ToID("test1234")
		assert.NoError(t, err)
		assert.Equal(t, "test123400", id)
	})
	t.Run("test with padded option short", func(t *testing.T) {
		f := New(WithPadding("0", 5))
		assert.NotNil(t, f)
		_, err := f.ToID("Test1234")
		assert.Error(t, err)
	})
	t.Run("test with multiple option", func(t *testing.T) {
		f := New(WithNormalizing(), WithPadding("1", 10))
		assert.NotNil(t, f)
		id, err := f.ToID("TesT1234")
		assert.NoError(t, err)
		assert.Equal(t, "test123411", id)
	})
	t.Run("test with sequenced option", func(t *testing.T) {
		f := New(WithNormalizing(),
			WithPadding("*", 36),
			WithBase64Encoding(),
			WithPrefix("id-"))
		assert.NotNil(t, f)
		id, err := f.ToID("TesT1234")
		assert.NoError(t, err)
		assert.Contains(t, id, "id-")
	})
	t.Run("test for reproducibility", func(t *testing.T) {
		f1 := New(WithNormalizing(),
			WithPadding("*", 10),
			WithBase64Encoding(),
			WithPrefix("id-"))
		assert.NotNil(t, f1)
		id1, err := f1.ToID("TesT1234")
		assert.NoError(t, err)
		id2, err := f1.ToID("TesT1234")
		assert.NoError(t, err)
		assert.Equal(t, id1, id2)
		f2 := New(WithNormalizing(),
			WithPadding("*", 10),
			WithBase64Encoding(),
			WithPrefix("id-"))
		assert.NotNil(t, f2)
		id3, err := f2.ToID("TesT1234")
		assert.NoError(t, err)
		assert.Equal(t, id2, id3)
	})
	t.Run("test with sequenced and predictability option", func(t *testing.T) {
		f := New(WithPadding("^", 20),
			WithSHA256Encoding(),
			WithPrefix("id-"),
			WithSuffix("-test"))
		assert.NotNil(t, f)
		id, err := f.ToID("user1234")
		assert.NoError(t, err)
		assert.Contains(t, id, "id-")
		assert.Contains(t, id, "-test")
		assert.Len(t, id, 72) // 64 from sha and 3 from the prefix and 5 from suffix
	})
	t.Run("test with sequenced and datetime option", func(t *testing.T) {
		f := New(WithPadding("^", 20),
			WithSHA256Encoding(),
			WithPrefix("id-"),
			WithDatetime("2006-01-02", "-", time.Now().UTC()))
		assert.NotNil(t, f)
		id, err := f.ToID("user1234")
		assert.NoError(t, err)
		assert.Contains(t, id, "id-")
		assert.Contains(t, id, time.Now().UTC().Format("2006-01-02"))
		assert.Len(t, id, 78) // 64 from sha and 3 from the prefix and 11 from suffix
	})
}
