// integration test
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"bytes"
	"encoding/json"
	"gomentoring/entities"
	"gomentoring/repository"
	"gomentoring/handlers"
	"gomentoring/usecase"
)

func TestCreateProduct(t *testing.T) {
	r := setupRouter()

	product := entities.Product{
		Id: 3,
		Name: "Laptop",
		Description: "Gaming Laptop",
		BarCode: "123456789",
		Price: 1299,
	}

	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	repo := repository.NewInMemoryProductRepo()
	uc := usecase.NewProductUsecase(repo)
	handlers.NewProductHandler(r, uc)
	return r
}