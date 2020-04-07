package sql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestMigrateCharlie(t *testing.T) {
	sql, _, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", sql)

	t.Run("NEGATIVE_MIGRATE_CHARLIE", func(t *testing.T) {
		MigrateCharlie(gormDB)
	})
}
