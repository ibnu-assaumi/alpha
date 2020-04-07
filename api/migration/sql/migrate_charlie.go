package sql

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
)

// MigrateCharlie : migrate / auto create charlie table if does not exists
func MigrateCharlie(db *gorm.DB) {
	domainCharlie := new(charlie.Domain)
	if !db.HasTable(domainCharlie) {
		if err := db.AutoMigrate(domainCharlie).Error; err != nil {
			fmt.Println(fmt.Sprintf("error migrate domain charlie : %s", err.Error()))
		}
	}

	domainCharlieHistory := new(charlie.DomainHistory)
	if !db.HasTable(domainCharlieHistory) {
		if err := db.AutoMigrate(domainCharlieHistory).Error; err != nil {
			fmt.Println(fmt.Sprintf("error migrate domain charlie history : %s", err.Error()))
		}
	}
}
