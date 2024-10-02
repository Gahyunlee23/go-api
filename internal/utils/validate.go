package utils

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type EntityNotFoundError struct {
	EntityType string
	ID         uint
}

func (u *EntityNotFoundError) Error() string {
	return fmt.Sprintf("Entity '%s' with ID '%d' not found", u.EntityType, u.ID)
}

func ValidateAndFetchEntity[T any](repo interface{ GetByID(id uint) (*T, error) }, id uint, entityType string) (*T, error) {
	entity, err := repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &EntityNotFoundError{EntityType: entityType, ID: id}
		}
		return nil, fmt.Errorf("error fetching %s: %w", entityType, err)
	}
	return entity, nil
}
