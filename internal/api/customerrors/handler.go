package customerrors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var entityNotFoundErr *EntityNotFoundError
	var validationErr *ValidationError

	switch {
	case errors.As(err, &entityNotFoundErr):
		c.JSON(http.StatusNotFound, gin.H{"error": entityNotFoundErr.Error()})
	case errors.As(err, &validationErr):
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
