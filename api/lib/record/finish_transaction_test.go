package record

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func getFakeDB() *gorm.DB {
	sql, _, _ := sqlmock.New()
	gormDB, _ := gorm.Open("postgres", sql)
	return gormDB
}

func TestFinishTransaction(t *testing.T) {
	t.Parallel()
	t.Run("COMMITED_TRANSACTION", func(t *testing.T) {
		db := getFakeDB()
		FinishTransaction(db, nil)
	})

	t.Run("ROLLED_BACK_TRANSACTION", func(t *testing.T) {
		db := getFakeDB()
		FinishTransaction(db, errors.New("error"))
	})
}
