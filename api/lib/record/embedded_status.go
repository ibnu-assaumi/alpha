package record

import "time"

// EmbeddedStatus : embedded status record in every structs / domain
type EmbeddedStatus struct {
	UserIn       uint64    `json:"userIn" gorm:"type:BIGINT;UNSIGNED;NOT_NULL"`
	UserUp       uint64    `json:"userUp" gorm:"type:BIGINT;UNSIGNED;DEFAULT:NULL"`
	DateIn       time.Time `json:"dateIn" gorm:"type:TIMESTAMP;NOT_NULL"`
	DateUp       time.Time `json:"dateUp" gorm:"type:TIMESTAMP;DEFAULT:NULL"`
	StatusRecord string    `json:"statusRecord" gorm:"NOT_NULL"`
}
