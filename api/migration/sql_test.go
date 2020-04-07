package migration

import (
	"testing"

	"github.com/Bhinneka/alpha/api/config"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestMigrateSQL(t *testing.T) {
	sql, _, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", sql)

	t.Run("POSITIVE_MIGRATE", func(t *testing.T) {
		config.PostgresWrite = gormDB
		MigrateSQL()
	})

	t.Run("NEGATIVE_SQL_MIGRATION", func(t *testing.T) {

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		config.PostgresWrite = nil
		MigrateSQL()
	})
}
