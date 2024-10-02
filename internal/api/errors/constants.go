package errors

import "errors"

var (
	ErrValidationFailed  = errors.New("validation failed")
	ErrNotFound          = errors.New("entity not found")
	ErrDatabaseOperation = errors.New("database operation failed")
	ErrInvalidInput      = errors.New("invalid input")
)
