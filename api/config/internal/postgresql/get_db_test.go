package postgresql

import (
	"testing"

	"github.com/Bhinneka/alpha/api/config/internal/postgresql/postgresqltest"
	"github.com/stretchr/testify/assert"
)

func TestGetDBRead(t *testing.T) {
	t.Parallel()
	t.Run("POSITIVE_GET_DB_READ", func(t *testing.T) {
		postgresqltest.LoadEnv()
		db := GetDBRead()
		defer db.Close()
		assert.NotNil(t, db)
		assert.NoError(t, db.DB().Ping())
	})

	t.Run("PANIC_GET_DB_READ", func(t *testing.T) {

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		postgresqltest.SetFakeENV()
		GetDBRead()
	})
}

func TestGetDBWrite(t *testing.T) {
	t.Parallel()
	t.Run("POSITIVE_GET_DB_WRITE", func(t *testing.T) {
		postgresqltest.LoadEnv()
		db := GetDBWrite()
		defer db.Close()
		assert.NotNil(t, db)
		assert.NoError(t, db.DB().Ping())
	})

	t.Run("PANIC_GET_DB_WRITE", func(t *testing.T) {

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		postgresqltest.SetFakeENV()
		GetDBWrite()
	})
}
