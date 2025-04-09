package handlers

import (
	"gomentoring/entities"
	"gomentoring/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Usecase *usecase.ProductUsecase
}

func NewProductHandler(r *gin.Engine, uc *usecase.ProductUsecase) {
	handler := &ProductHandler{Usecase: uc}
	r.GET("/products/all", handler.GetAllProducts)
	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:barcode", handler.GetProductByBarCode)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	h.Usecase.Repo.Save(&product)
	c.JSON(http.StatusCreated, gin.H{"message": "product created", "body": product})
}

func (h *ProductHandler) GetProductByBarCode(c *gin.Context) {
	barcode := c.Param("barcode")
	product, stock, err := h.Usecase.GetProductWithStock(barcode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "product found",
		"body": gin.H{
			"product": product,
			"stock":   stock,
		},
	})
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products := h.Usecase.GetAllProducts()
	c.JSON(http.StatusOK, gin.H{
		"message": "products found",
		"body": gin.H{
			"products": products,
		},
	})
}