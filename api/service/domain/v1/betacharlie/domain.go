package betacharlie

const TableName string = "pref_beta_charlie"

type Domain struct {
	BetaCharlieID uint64 `json:"betaCharlieID" bson:"betaCharlieID" gorm:"PRIMARY_KEY;AUTO_INCREMENT;type:SERIAL;BIGINT;UNSIGNED;"`
	BetaID        uint64 `json:"betaID" bson:"betaID" gorm:"column_name:index;type:SERIAL;BIGINT;UNSIGNED;"`
	CharlieID     uint64 `json:"charlieID" bson:"charlieID" gorm:"column_name:index;type:SERIAL;BIGINT;UNSIGNED;"`
}

// TableName : implement gorm override table name
func (impl *Domain) TableName() string {
	return TableName
}
