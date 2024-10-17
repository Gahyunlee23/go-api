package utils

import (
	"errors"
	"fmt"
	"main-admin-api/internal/api/customerrors"

	"gorm.io/gorm"
)

// ValidateAndFetchEntity fetches an entity and validates its existence
func ValidateAndFetchEntity[T any](repo interface{ GetByID(id uint) (*T, error) }, id uint, entityType string) (*T, error) {
	entity, err := repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &customerrors.EntityNotFoundError{EntityType: entityType, ID: id}
		}
		return nil, fmt.Errorf("error fetching %s: %w", entityType, err)
	}
	return entity, nil
}

// ValidateID checks if the resource ID matches the URL ID and check existence of Resource ID
func ValidateID(urlID, resourceID uint) error {
	if resourceID == 0 {
		return &customerrors.CheckIDExists{ResourceId: resourceID}
	}

	if urlID != resourceID {
		return &customerrors.IDMismatchError{URLId: urlID, ResourceId: resourceID}
	}

	return nil
}
