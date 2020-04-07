package errorutil

import (
	"fmt"
)

// Error : custom error
type Error struct {
	key     string
	Message map[string][]string `json:"error"`
}

// NewError : new custom error helper
func NewError(key string) *Error {
	msg := make(map[string][]string)
	msg[key] = []string{}
	return &Error{
		key:     key,
		Message: msg,
	}
}

// Error : implement error
func (e *Error) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

// AppendMessage : append error message
func (e *Error) AppendMessage(value string) *Error {
	e.Message[e.key] = append(e.Message[e.key], value)
	return e
}

// IsSystemError : check if error is custom error
func IsSystemError(err error) (message []string, system bool) {
	switch err.(type) {
	// custom error type
	case *Error:
		ce := err.(*Error)
		return ce.GetErrorMessage(), false
	// system error type
	default:
		return nil, true
	}
}

// HasError : check if custom error has some errors
func (e *Error) HasError() bool {
	if len(e.Message[e.key]) > 0 {
		return true
	}
	return false
}

// GetErrorMessage : get appended error messages
func (e *Error) GetErrorMessage() []string {
	return e.Message[e.key]
}
