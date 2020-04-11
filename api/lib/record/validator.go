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

// ValidateDateString : validate empty date return as string
func ValidateDateString(date time.Time, layout string) *string {
	if date.IsZero() {
		return nil
	}
	dateString := date.Format(layout)
	return &dateString
}

// ValidateDateTimeString : validate date time string
func ValidateDateTimeString(date time.Time) *string {
	if date.IsZero() {
		return nil
	}
	dateString := date.Format(time.RFC3339)
	return &dateString
}

// ValidateDateTime : validate empty date return as date
func ValidateDateTime(date time.Time) *time.Time {
	if date.IsZero() {
		return nil
	}
	return &date
}
