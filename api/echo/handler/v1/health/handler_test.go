package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	t.Run("POSITIVE_NEW_HANDLER", func(t *testing.T) {
		assert.NotNil(t, NewHandler())
	})
}
