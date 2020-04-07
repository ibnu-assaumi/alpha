package record

import (
	"strings"
	"time"
)

// ValidateID : validate empty id
func ValidateID(id uint64) *uint64 {
	if id == 0 {
		return nil
	}
	return &id
}

// ValidateString : validate empty string
func ValidateString(s string) *string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	return &s
}

// ValidateDate : validate empty date
func ValidateDate(date time.Time, layout string) *string {
	if date.IsZero() {
		return nil
	}
	dateString := date.Format(layout)
	return &dateString
}
