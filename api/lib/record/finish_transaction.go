package record

import (
	"github.com/jinzhu/gorm"
)

// FinishTransaction : commit or rollback transaction
func FinishTransaction(db *gorm.DB, err error) {
	if err != nil {
		db.Rollback()
	} else {
		db.Commit()
	}
}
