package record

import "time"

// EmbeddedStatus : embedded status record in every structs / domain
type EmbeddedStatus struct {
	UserIn       uint64    `json:"userIn" bson:"user_in" gorm:"type:BIGINT;UNSIGNED;NOT_NULL"`
	UserUp       uint64    `json:"userUp" bson:"user_up" gorm:"type:BIGINT;UNSIGNED;DEFAULT:NULL"`
	DateIn       time.Time `json:"dateIn" bson:"date_in" gorm:"type:TIMESTAMP;NOT_NULL"`
	DateUp       time.Time `json:"dateUp" bson:"date_up" gorm:"type:TIMESTAMP;DEFAULT:NULL"`
	StatusRecord string    `json:"statusRecord" gorm:"NOT_NULL"`
}
