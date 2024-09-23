package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func MarshalAndAssignJSON(field interface{}, fieldName string, ctx *gin.Context) (datatypes.JSON, error) {
	jsonData, err := json.Marshal(field)
	if err != nil {
		log.Printf("Error marshalling %s: %v", fieldName, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process " + fieldName})
		return nil, err
	}
	return datatypes.JSON(jsonData), nil
}
