package usecase

import (
	"gomentoring/entities"
)

type ProductUsecase interface {
	Create(p *entities.Product) error
	GetAllProducts() []*entities.Product
	GetProductWithStock(barcode string) (*entities.Product, int, error)
}