package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	t.Run("ROOT_PATH", func(t *testing.T) {
		assert.NotNil(t, Root())
	})
}
