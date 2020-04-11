package sql

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/Bhinneka/alpha/api/service/domain/v1/beta"
)

// MigrateBeta : migrate / auto create beta table if does not exists
func MigrateBeta(db *gorm.DB) {
	domainBeta := new(beta.Domain)
	if !db.HasTable(domainBeta) {
		if err := db.AutoMigrate(domainBeta).Error; err != nil {
			fmt.Println(fmt.Sprintf("error migrate domain beta : %s", err.Error()))
		}
	}
}
