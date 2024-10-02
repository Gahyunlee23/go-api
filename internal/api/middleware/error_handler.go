package middleware

import (
	"main-admin-api/internal/api/errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errors.HandleError(c, err.(error))
			}
		}()
		c.Next()
	}
}
