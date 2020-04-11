package sql

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/Bhinneka/alpha/api/service/domain/v1/betacharlie"
)

// MigrateBetaCharlie : migrate / auto create beta charlie table if does not exists
func MigrateBetaCharlie(db *gorm.DB) {
	domainBetaCharlie := new(betacharlie.Domain)
	if !db.HasTable(domainBetaCharlie) {
		if err := db.AutoMigrate(domainBetaCharlie).Error; err != nil {
			fmt.Println(fmt.Sprintf("error migrate domain betacharlie : %s", err.Error()))
		}
	}
}
