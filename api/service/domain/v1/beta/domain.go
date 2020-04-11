package beta

import (
	"encoding/json"
	"time"

	"github.com/Bhinneka/alpha/api/lib/record"
	domainCharlie "github.com/Bhinneka/alpha/api/service/domain/v1/charlie"
	"go.mongodb.org/mongo-driver/bson"
)

// TableName : table name of beta
const TableName string = "pref_beta"

// Domain : beta model
type Domain struct {
	BetaID                uint64                 `json:"betaID" bson:"betaID" gorm:"PRIMARY_KEY;AUTO_INCREMENT;type:SERIAL;BIGINT;UNSIGNED;"`
	BetaName              string                 `json:"betaName" bson:"betaName"`
	Charlie               []domainCharlie.Domain `json:"charlie" bson:"charlie,inline"`
	record.EmbeddedStatus `bson:",inline"`
}

// TableName : implement gorm override table name
func (impl *Domain) TableName() string {
	return TableName
}

// MarshalJSON : implement json marshaller to return nil values instead of empty as a response
func (impl *Domain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		BetaID       *uint64                `json:"betaID"`
		BetaName     *string                `json:"betaName"`
		Charlie      []domainCharlie.Domain `json:"charlie"`
		UserIn       *uint64                `json:"userIn"`
		UserUp       *uint64                `json:"userUp"`
		DateIn       *string                `json:"dateIn"`
		DateUp       *string                `json:"dateUp"`
		StatusRecord *string                `json:"statusRecord"`
	}{
		BetaID:       record.ValidateID(impl.BetaID),
		BetaName:     record.ValidateString(impl.BetaName),
		Charlie:      impl.Charlie,
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
		BetaID       *uint64                `bson:"betaID"`
		BetaName     *string                `bson:"betaName"`
		Charlie      []domainCharlie.Domain `bson:"charlie"`
		UserIn       *uint64                `bson:"userIn"`
		UserUp       *uint64                `bson:"userUp"`
		DateIn       *time.Time             `bson:"dateIn"`
		DateUp       *time.Time             `bson:"dateUp"`
		StatusRecord *string                `bson:"statusRecord"`
	}{
		BetaID:       record.ValidateID(impl.BetaID),
		BetaName:     record.ValidateString(impl.BetaName),
		Charlie:      impl.Charlie,
		UserIn:       record.ValidateID(impl.UserIn),
		UserUp:       record.ValidateID(impl.UserUp),
		DateIn:       record.ValidateDateTime(impl.DateIn),
		DateUp:       record.ValidateDateTime(impl.DateUp),
		StatusRecord: record.ValidateString(impl.StatusRecord),
	})
}
