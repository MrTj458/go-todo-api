package model

import "errors"

var (
	ErrConflict       = errors.New("conflict")
	ErrInternal       = errors.New("internal")
	ErrInvalid        = errors.New("invalid")
	ErrNotFound       = errors.New("not_found")
	ErrNotImplemented = errors.New("not_implemented")
	ErrUnauthorized   = errors.New("unauthorized")
)

// Error represents a JSON error response
type Error struct {
	Code   int           `json:"code"`
	Detail string        `json:"detail"`
	Fields []*ErrorField `json:"fields"`
}

// Error prints the error detail to implement the error interface
func (e *Error) Error() string {
	return e.Detail
}

// ErrorField represents an error on a JSON field in a request
type ErrorField struct {
	Location string `json:"location"`
	Type     string `json:"type"`
	Detail   string `json:"detail"`
}

// NewError is a shortcut to create a new Error
func NewError(code int, detail string) *Error {
	return &Error{
		Code:   code,
		Detail: detail,
		Fields: []*ErrorField{},
	}
}

// NewErrorWithFields is a shortcut to create a new Error with a slice of
// ErrorFields
func NewErrorWithFields(code int, detail string, fields []*ErrorField) *Error {
	return &Error{
		Code:   code,
		Detail: detail,
		Fields: fields,
	}
}
