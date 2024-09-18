package controllers

import (
	"github.com/gin-gonic/gin"
	"main-admin-api/models"
	"main-admin-api/services"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{productService: service}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.productService.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	ctx.JSON(http.StatusCreated, product)
}

func (c *ProductController) GetProductByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := c.productService.GetProductByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.productService.UpdateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.productService.DeleteProduct(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
