package customerrors

import "fmt"

// EntityNotFoundError represents an error when an entity is not found
type EntityNotFoundError struct {
	EntityType string
	ID         uint
}

func (e *EntityNotFoundError) Error() string {
	return fmt.Sprintf("Entity '%s' with ID '%d' not found", e.EntityType, e.ID)
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field '%s': %s", e.Field, e.Message)
}

// IDMismatchError represents an error when IDs don't match
type IDMismatchError struct {
	URLId      uint
	ResourceId uint
}

func (e *IDMismatchError) Error() string {
	return fmt.Sprintf("ID '%d' does not match resource ID '%d'", e.URLId, e.ResourceId)
}

// CheckIDExists represents an error when the resource ID is not included in the request body
type CheckIDExists struct {
	ResourceId uint
}

func (e *CheckIDExists) Error() string {
	return fmt.Sprintf("ID '%d' is not valid in the request body", e.ResourceId)
}
