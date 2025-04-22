package main

import (
	"testing"
	"gomentoring/entities"
	"gomentoring/usecase"
	"gomentoring/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductUseCase(t *testing.T) {
	repo := repository.NewInMemoryProductRepo()
	uc := usecase.NewProductUsecase(repo)

	product := &entities.Product{
		Id: 1,
		Name: "Test Product",
		Description: "A product for testing",
		BarCode: "12345",
		Price: 100,
	}

	err := uc.Create(product)

	assert.Nil(t, err)

	savedProduct, err := repo.GetByBarCode("12345")
	assert.Nil(t, err)
	assert.Equal(t, "Test Product", savedProduct.Name)
}