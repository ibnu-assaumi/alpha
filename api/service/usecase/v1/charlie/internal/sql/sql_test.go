package sql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSQL(t *testing.T) {
	t.Run("POSITIVE_NEW_SQL", func(t *testing.T) {
		assert.NotNil(t, NewSQL())
	})
}
