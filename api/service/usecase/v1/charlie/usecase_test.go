package charlie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUseCase(t *testing.T) {
	t.Run("POSITIVE_NEW_USECASE", func(t *testing.T) {
		assert.NotNil(t, NewUseCase())
	})
}
