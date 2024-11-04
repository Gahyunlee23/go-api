package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Search(c *gin.Context, fields []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, field := range fields {
			if value := c.Query(field); value != "" {
				db = db.Where(fmt.Sprintf("%s LIKE ?", field), "%"+value+"%")
			}
		}

		if searchTerm := c.Query("search"); searchTerm != "" {
			words := strings.Fields(searchTerm)
			for _, word := range words {
				subQuery := db.Where("")
				for _, field := range fields {
					subQuery = subQuery.Or(fmt.Sprintf("%s LIKE ?", field), "%"+word+"%")
				}
				db = db.Where(subQuery)
			}
		}

		return db
	}
}
