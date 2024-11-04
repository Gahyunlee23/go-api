package utils

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

func isValidSortDirection(direction string) bool {
	sortBy := []string{"asc", "desc"}
	return funk.Contains(sortBy, direction)
}

func Sort(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		sort := c.QueryMap("sort")

		for sortingCol, direction := range sort {
			sortingCol := strings.ToLower(sortingCol)
			direction := strings.ToLower(direction)

			if !isValidSortDirection(direction) {
				log.Printf("invalid sort direction: %s", direction)
			}

			db = db.Order(fmt.Sprintf("%s %s", sortingCol, direction))
		}
		return db
	}
}
