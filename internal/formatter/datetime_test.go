package formatter

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatetimeFormatter(t *testing.T) {
	t.Run("test datetime Formatter", func(t *testing.T) {
		f := &DatetimeFormatter{
			BaseFormat: "2006-01-02",
			Separator:  ":",
			UTC:        true,
		}
		assert.NotNil(t, f)
		id, err := f.Format("id1234")
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("id1234:%s", time.Now().UTC().Format("2006-01-02")), id)
	})
}
