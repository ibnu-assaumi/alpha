package charlie

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Bhinneka/alpha/api/lib/constant"
	"github.com/Bhinneka/alpha/api/lib/record"
)

// TableName : table name of charlie
const TableName string = "api_charlie"

// Domain : charlie model
type Domain struct {
	CharlieID   uint64 `json:"charlieID" bson:"charlie_id" gorm:"column_name:charlie_id;PRIMARY_KEY;AUTO_INCREMENT;type:SERIAL;BIGINT;UNSIGNED;"`
	CharlieName string `json:"charlieName" bson:"charlie_name"`
	record.EmbeddedStatus
}

// TableName : implement gorm override table name
func (impl *Domain) TableName() string {
	return TableName
}

// MarshalJSON : implement json marshaller to return nil values instead of empty as a response
func (impl *Domain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		CharlieID    *uint64 `json:"charlieID"`
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
		DateIn:       record.ValidateDateString(impl.DateIn, constant.DateTimeLayout),
		DateUp:       record.ValidateDateString(impl.DateUp, constant.DateTimeLayout),
		StatusRecord: record.ValidateString(impl.StatusRecord),
	})
}

// MarshalBSONValue : implement bson marshaller to return nil values instead of empty as a response
func (impl *Domain) MarshalBSONValue() ([]byte, error) {
	return bson.Marshal(struct {
		CharlieID    *uint64 `json:"charlieID" bson:"charlie_id"`
		CharlieName  *string `json:"charlieName" bson:"charlie_name"`
		UserIn       *uint64 `json:"userIn" bson:"user_in"`
		UserUp       *uint64 `json:"userUp" bson:"user_up"`
		DateIn       *string `json:"dateIn" bson:"date_in"`
		DateUp       *string `json:"dateUp" bson:"date_up"`
		StatusRecord *string `bson:"statusRecord"`
	}{
		CharlieID:    record.ValidateID(impl.CharlieID),
		CharlieName:  record.ValidateString(impl.CharlieName),
		UserIn:       record.ValidateID(impl.UserIn),
		UserUp:       record.ValidateID(impl.UserUp),
		DateIn:       record.ValidateDateString(impl.DateIn, constant.DateTimeLayout),
		DateUp:       record.ValidateDateString(impl.DateUp, constant.DateTimeLayout),
		StatusRecord: record.ValidateString(impl.StatusRecord),
	})
}

// MarshalElasticValue : json marshaller to value for elastic
func (impl *Domain) MarshalElasticValue() ([]byte, error) {
	return json.Marshal(struct {
		CharlieID    *uint64    `json:"charlieID"`
		CharlieName  *string    `json:"charlieName"`
		UserIn       *uint64    `json:"userIn"`
		UserUp       *uint64    `json:"userUp"`
		DateIn       *time.Time `json:"dateIn"`
		DateUp       *time.Time `json:"dateUp"`
		StatusRecord *string    `json:"statusRecord"`
	}{
		CharlieID:    record.ValidateID(impl.CharlieID),
		CharlieName:  record.ValidateString(impl.CharlieName),
		UserIn:       record.ValidateID(impl.UserIn),
		UserUp:       record.ValidateID(impl.UserUp),
		DateIn:       record.ValidateDateTime(impl.DateIn.UTC()),
		DateUp:       record.ValidateDateTime(impl.DateUp.UTC()),
		StatusRecord: record.ValidateString(impl.StatusRecord),
	})
}
