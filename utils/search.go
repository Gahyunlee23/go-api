package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Search(c *gin.Context, fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		searchTerm := c.Query("search")
		if searchTerm == "" {
			return db
		}

		for i, field := range fields {
			if i == 0 {
				db = db.Where(fmt.Sprintf("%s LIKE ?", field), "%"+searchTerm+"%")
			} else {
				db = db.Or(fmt.Sprintf("%s LIKE ?", field), "%"+searchTerm+"%")
			}
		}
		return db
	}
}
