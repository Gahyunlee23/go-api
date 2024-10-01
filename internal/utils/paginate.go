package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"page is not valid": err.Error()})
		}
		if page == 0 {
			page = 1
		}

		pageSize, err := strconv.Atoi(c.Query("page_size"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"page size is not valid": err.Error()})
		}
		if pageSize == 0 {
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
