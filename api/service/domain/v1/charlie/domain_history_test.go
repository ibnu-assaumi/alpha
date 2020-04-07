package charlie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainHistory_TableName(t *testing.T) {
	t.Run("POSITIVE_TABLE_NAME", func(t *testing.T) {
		obj := &DomainHistory{}
		assert.Equal(t, TableNameHistory, obj.TableName())
	})
}
