package services_test

import (
	"main-admin-api/internal/models"
	"main-admin-api/internal/repository/mocks"
	"main-admin-api/internal/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := services.NewProductService(mockRepo)

	mockRepo.EXPECT().Create(gomock.Any()).Return(nil)
	duration := "5 days"
	maxQun := 100000

	product := &models.Product{
		CloudLabID:           11191119,
		Name:                 "stella_1",
		Code:                 "stella_1",
		Type:                 "stella_1",
		MinimumQuantity:      0,
		MaximumQuantity:      &maxQun,
		PackingUnit:          1,
		EnableCustomQuantity: true,
		EnableCustomFormat:   false,
		TimeToProduce:        &duration,
		//RenamingRules: []map[string]string{
		//	{"refinement": "Numbering"},
		//	{"book_binding": "Binding"},
		//	{"colors": "Ink"},
		//	{"finishing": "Packaging"},
		//},
		//OrderRules: [["paper", "format", "pages", "colors"]],
		DefaultQuantity: 0,
		//QuantitiesSelection: [],
		PriceCalculationType: ""}

	ctx, _ := gin.CreateTestContext(nil)

	err := productService.CreateProduct(product, ctx)

	assert.NoError(t, err)
	mockRepo.EXPECT().Create(product)
}

func TestGetProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := services.NewProductService(mockRepo)

	expectedProduct := &models.Product{
		CloudLabID:      11191119,
		Name:            "stella_1",
		Code:            "stella_1",
		Type:            "stella_1",
		MinimumQuantity: 0,
	}

	mockRepo.EXPECT().GetByID(uint(1)).Return(expectedProduct, nil)

	product, err := productService.GetProductByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)

	mockRepo.EXPECT().GetByID(uint(1))
}
