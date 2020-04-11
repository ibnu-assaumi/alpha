package charlie

import (
	"encoding/json"
	"time"

	"github.com/Bhinneka/alpha/api/lib/record"
	"go.mongodb.org/mongo-driver/bson"
)

// TableName : table name of charlie
const TableName string = "pref_charlie"

// Domain : charlie model
type Domain struct {
	CharlieID             uint64 `json:"charlieID" bson:"charlieID" gorm:"PRIMARY_KEY;AUTO_INCREMENT;type:SERIAL;BIGINT;UNSIGNED;"`
	CharlieName           string `json:"charlieName" bson:"charlieName"`
	record.EmbeddedStatus `bson:",inline"`
}

// TableName : implement gorm override table name
func (impl *Domain) TableName() string {
	return TableName
}

// MarshalJSON : implement json marshaller to return nil values instead of empty as a response
func (impl *Domain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		CharlieID    *uint64 `json:"charlieID" `
		CharlieName  *string `json:"charlieName"`
		UserIn       *uint64 `json:"userIn"`
		UserUp       *uint64 `json:"userUp"`
		DateIn       *string `json:"dateIn"`
		DateUp       *string `json:"dateUp"`
		StatusRecord *string `json:"statusRecord"`
	}{
		CharlieID:    record.ValidateID(impl.CharlieID),
		CharlieName:  record.ValidateString(impl.CharlieName),
		UserIn:       record.ValidateID(impl.UserIn),
		UserUp:       record.ValidateID(impl.UserUp),
		DateIn:       record.ValidateDateTimeString(impl.DateIn),
		DateUp:       record.ValidateDateTimeString(impl.DateUp),
		StatusRecord: record.ValidateString(impl.StatusRecord),
	})
}

// MarshalBSONValue : implement bson marshaller to return nil values instead of empty as a response
func (impl *Domain) MarshalBSONValue() ([]byte, error) {
	return bson.Marshal(struct {
		CharlieID    *uint64    `bson:"charlieID"`
		CharlieName  *string    `bson:"charlieName"`
		UserIn       *uint64    `bson:"userIn"`
		UserUp       *uint64    `bson:"userUp"`
		DateIn       *time.Time `bson:"dateIn"`
		DateUp       *time.Time `bson:"dateUp"`
		StatusRecord *string    `bson:"statusRecord"`
	}{
		CharlieID:    record.ValidateID(impl.CharlieID),
		CharlieName:  record.ValidateString(impl.CharlieName),
		UserIn:       record.ValidateID(impl.UserIn),
		UserUp:       record.ValidateID(impl.UserUp),
		DateIn:       record.ValidateDateTime(impl.DateIn),
		DateUp:       record.ValidateDateTime(impl.DateUp),
		StatusRecord: record.ValidateString(impl.StatusRecord),
	})
}
