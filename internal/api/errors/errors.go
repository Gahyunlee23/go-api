package errors

import "fmt"

// EntityNotFoundError NotFoundError represents an error when an entity is not found
type EntityNotFoundError struct {
	EntityType string
	ID         uint
}

func (e *EntityNotFoundError) Error() string {
	return fmt.Sprintf("Entity '%s' with ID '%d' not found", e.EntityType, e.ID)
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field '%s': %s", e.Field, e.Message)
}
