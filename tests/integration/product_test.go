package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main-admin-api/internal/models"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func intPtr(i int) *int {
	return &i
}

func stringPtr(s string) *string {
	return &s
}

func TestCreateProduct(t *testing.T) {
	product := models.Product{
		CloudLabID:           11191119,
		Name:                 "stella_1",
		Code:                 "stella_1",
		Type:                 "stella_1",
		MinimumQuantity:      0,
		MaximumQuantity:      intPtr(100000),
		PackingUnit:          1,
		EnableCustomQuantity: true,
		EnableCustomFormat:   false,
		TimeToProduce:        stringPtr("5 days"),
		RenamingRules:        datatypes.JSON(`[{"refinement":"Numbering"},{"book_binding":"Binding"},{"colors":"Ink"},{"finishing":"Packaging"}]`),
		OrderRules:           datatypes.JSON(`[["paper", "format", "pages", "colors"]]`),
		DefaultQuantity:      0,
		QuantitiesSelection:  datatypes.JSON(`[]`),
		PriceCalculationType: "",
	}
	jsonProduct, _ := json.Marshal(product)

	resp, err := http.Post("http://localhost:8080/products", "application/json", bytes.NewBuffer(jsonProduct))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createdProduct models.Product
	err = json.NewDecoder(resp.Body).Decode(&createdProduct)
	assert.NoError(t, err)
	assert.NotZero(t, createdProduct.ID)
	assert.Equal(t, product.Name, createdProduct.Name)
	assert.Equal(t, product.Code, createdProduct.Code)
	assert.Equal(t, product.Type, createdProduct.Type)
}

func TestGetProductByID(t *testing.T) {
	// First creating a product
	product := models.Product{
		CloudLabID:           11191119,
		Name:                 "stella_1",
		Code:                 "stella_1",
		Type:                 "stella_1",
		MinimumQuantity:      0,
		MaximumQuantity:      intPtr(100000),
		PackingUnit:          1,
		EnableCustomQuantity: true,
		EnableCustomFormat:   false,
		TimeToProduce:        stringPtr("5 days"),
		RenamingRules:        datatypes.JSON(`[{"refinement":"Numbering"},{"book_binding":"Binding"},{"colors":"Ink"},{"finishing":"Packaging"}]`),
		OrderRules:           datatypes.JSON(`[["paper", "format", "pages", "colors"]]`),
		DefaultQuantity:      0,
		QuantitiesSelection:  datatypes.JSON(`[]`),
		PriceCalculationType: "",
	}

	jsonProduct, _ := json.Marshal(product)
	resp, err := http.Post("http://localhost:8080/products", "application/json", bytes.NewBuffer(jsonProduct))
	assert.NoError(t, err)

	var createdProduct models.Product
	err = json.NewDecoder(resp.Body).Decode(&createdProduct)
	assert.NoError(t, err)

	// Next getting a product that we created
	resp, err = http.Get(fmt.Sprintf("http://localhost:8080/products/%d", createdProduct.ID))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var fetchedProduct models.Product
	err = json.NewDecoder(resp.Body).Decode(&fetchedProduct)
	assert.NoError(t, err)
	assert.Equal(t, createdProduct.ID, fetchedProduct.ID)
	assert.Equal(t, product.Name, fetchedProduct.Name)
	assert.Equal(t, product.Code, fetchedProduct.Code)
	assert.Equal(t, product.Type, fetchedProduct.Type)
}
