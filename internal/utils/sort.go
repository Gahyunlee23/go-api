package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

func isValidSortDirection(direction string) bool {
	sortBy := []string{"asc", "desc"}
	return funk.Contains(sortBy, direction)
}

func Sort(c *gin.Context, allowedColumns []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		sort := c.QueryMap("sort")

		for sortingCol, direction := range sort {
			sortingCol := strings.ToLower(sortingCol)
			direction := strings.ToLower(direction)

			if !funk.Contains(allowedColumns, sortingCol) {
				db.Error = fmt.Errorf("invalid sort column: %s", sortingCol)
				return db
			}

			if !isValidSortDirection(direction) {
				db.Error = fmt.Errorf("invalid sort direction: %s", direction)
				return db
			}

			db = db.Order(fmt.Sprintf("%s %s", sortingCol, direction))
		}
		return db
	}
}
