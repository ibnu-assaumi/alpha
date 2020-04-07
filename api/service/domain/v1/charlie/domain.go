package charlie

import (
	"encoding/json"

	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/record"
)

// TableName : table name of charlie
const TableName string = "api_charlie"

// Domain : charlie model
type Domain struct {
	CharlieID   uint64 `json:"charlieID" gorm:"column_name:charlie_id;PRIMARY_KEY;AUTO_INCREMENT;type:SERIAL;BIGINT;UNSIGNED;"`
	CharlieName string `json:"charlieName"`
	record.EmbeddedStatus
}

// TableName : implement gorm override table name
func (impl *Domain) TableName() string {
	return TableName
}

// MarshalJSON : implement json marshaller to return nil values instead of empty as a response
func (impl *Domain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		CharlieID   *uint64 `json:"charlieID"`
		CharlieName *string `json:"charlieName"`
		UserIn      *uint64 `json:"userIn"`
		UserUp      *uint64 `json:"userUp"`
		DateIn      *string `json:"dateIn"`
		DateUp      *string `json:"dateUp"`
	}{
		CharlieID:   record.ValidateID(impl.CharlieID),
		CharlieName: record.ValidateString(impl.CharlieName),
		UserIn:      record.ValidateID(impl.UserIn),
		UserUp:      record.ValidateID(impl.UserUp),
		DateIn:      record.ValidateDate(impl.DateIn, constant.DateTimeLayout),
		DateUp:      record.ValidateDate(impl.DateUp, constant.DateTimeLayout),
	})
}
