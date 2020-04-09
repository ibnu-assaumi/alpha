package record

import (
	"testing"
	"time"

	"github.com/Bhinneka/alpha/api/lib/constant"

	"github.com/stretchr/testify/assert"
)

func TestValidateID(t *testing.T) {
	t.Parallel()
	t.Run("NIL_VALIDATE_ID", func(t *testing.T) {
		id := uint64(0)
		assert.Nil(t, ValidateID(id))
	})

	t.Run("NOT_NIL_VALIDATE_ID", func(t *testing.T) {
		id := uint64(1)
		assert.NotNil(t, ValidateID(id))
	})
}

func TestValidateString(t *testing.T) {
	t.Parallel()
	t.Run("NIL_VALIDATE_STRING", func(t *testing.T) {
		str := ""
		assert.Nil(t, ValidateString(str))
	})

	t.Run("NOT_NIL_VALIDATE_STRING", func(t *testing.T) {
		str := "test"
		assert.NotNil(t, ValidateString(str))
	})
}

func TestValidateDateString(t *testing.T) {
	t.Parallel()
	t.Run("NIL_VALIDATE_DATE", func(t *testing.T) {
		dt := time.Time{}
		assert.Nil(t, ValidateDateString(dt, constant.DateLayout))
	})

	t.Run("NOT_NIL_VALIDATE_DATE", func(t *testing.T) {
		now := time.Now()
		assert.NotNil(t, ValidateDateString(now, constant.DateLayout))
	})
}
