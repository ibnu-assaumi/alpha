package postgresql

import (
	"os"
	"testing"

	"github.com/Bhinneka/alpha/api/config/internal/postgresql/postgresqltest"
)

func Test_setDBStats(t *testing.T) {
	t.Parallel()
	t.Run("POSITIVE_SET_DB_STATS", func(t *testing.T) {
		postgresqltest.LoadEnv()
		db := postgresqltest.GetFakeDB()
		setDBStats(db, dbTypeRead)
	})

	t.Run("PANIC_PING_SET_DB_STATS", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		db := postgresqltest.GetFakeDB()
		db.Close() // close the database to make it panic
		setDBStats(db, dbTypeRead)
	})

	t.Run("PANIC_TYPE_SET_DB_STATS", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		db := postgresqltest.GetFakeDB()
		setDBStats(db, "")
	})

	t.Run("NEGATIVE_SET_LOG_MODE", func(t *testing.T) {
		db := postgresqltest.GetFakeDB()
		defer db.Close()
		os.Setenv("READ_POSTGRES_DB_LOG_MODE", "x")
		setDBStats(db, dbTypeRead)
	})

	t.Run("NEGATIVE_SET_MAX_IDLE_CONN", func(t *testing.T) {
		db := postgresqltest.GetFakeDB()
		defer db.Close()
		os.Setenv("READ_POSTGRES_DB_LOG_MODE", "true")
		os.Setenv("READ_POSTGRES_DB_MAX_IDLE_CONN", "x")
		setDBStats(db, dbTypeRead)
	})

	t.Run("NEGATIVE_SET_MAX_OPEN_CONN", func(t *testing.T) {
		db := postgresqltest.GetFakeDB()
		defer db.Close()
		os.Setenv("READ_POSTGRES_DB_LOG_MODE", "true")
		os.Setenv("READ_POSTGRES_DB_MAX_IDLE_CONN", "10")
		os.Setenv("READ_POSTGRES_DB_MAX_OPEN_CONN", "x")
		setDBStats(db, dbTypeRead)
	})

	t.Run("NEGATIVE_SET_MAX_LIFE_TIME", func(t *testing.T) {
		db := postgresqltest.GetFakeDB()
		defer db.Close()
		os.Setenv("READ_POSTGRES_DB_LOG_MODE", "true")
		os.Setenv("READ_POSTGRES_DB_MAX_IDLE_CONN", "10")
		os.Setenv("READ_POSTGRES_DB_MAX_OPEN_CONN", "10")
		os.Setenv("READ_POSTGRES_DB_MAX_LIFETIME", "x")
		setDBStats(db, dbTypeRead)
	})
}
