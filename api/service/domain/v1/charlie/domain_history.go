package charlie

import "github.com/Bhinneka/alpha/api/lib/record"

// TableNameHistory : table name of charlie history model
const TableNameHistory string = "history_api_charlie"

// DomainHistory : charlie history model
type DomainHistory struct {
	CharlieIDHistory uint64 `json:"charlieIDHistory" gorm:"PRIMARY_KEY;AUTO_INCREMENT;type:SERIAL;BIGINT;UNSIGNED;"`
	CharlieID        uint64 `json:"charlieID" gorm:"index;type:SERIAL;BIGINT;UNSIGNED;"`
	CharlieName      string `json:"charlieName"`
	record.EmbeddedStatus
}

// TableName : implement gorm override table name
func (impl *DomainHistory) TableName() string {
	return TableNameHistory
}
