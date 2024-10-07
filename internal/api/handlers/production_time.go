package handler

import (
	"main-admin-api/internal/api/customerrors"
	"main-admin-api/internal/models"
	services "main-admin-api/internal/services/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductionTimeHandler struct {
	productionTimeService services.ProductionTimeService
}

func NewProductionTimeHandler(service services.ProductionTimeService) *ProductionTimeHandler {
	return &ProductionTimeHandler{productionTimeService: service}
}

// CreateProductionTime godoc
// @Summary Create a new productionTime
// @Tags ProductionTime
// @Accept json
// @Produce  json
// @Param   productionTime  body  models.ProductionTime  true  "Production Time data"
// @Success 200 {object} models.ProductionTime
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /product-ionTimes/ [post]
func (c *ProductionTimeHandler) CreateProductionTime(ctx *gin.Context) {
	var productionTime models.ProductionTime

	if err := ctx.ShouldBindJSON(&productionTime); err != nil {
		customerrors.HandleError(ctx, &customerrors.ValidationError{Field: "body", Message: err.Error()})
		return
	}

	if err := c.productionTimeService.CreateProductionTime(&productionTime); err != nil {
		customerrors.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, productionTime)
}
