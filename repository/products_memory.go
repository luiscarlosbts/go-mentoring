package repository
import (
	"gomentoring/entities"
	"errors"
)

type InMemoryProductRepo struct {
	data map[string]*entities.Product
}

func NewInMemoryProductRepo() *InMemoryProductRepo {
	return &InMemoryProductRepo{data: make(map[string]*entities.Product)}
}

func (r *InMemoryProductRepo) Save(product *entities.Product) error {
	r.data[product.BarCode] = product
	return nil
}

func (r *InMemoryProductRepo) GetByBarCode(barcode string) (*entities.Product, error) {
	product, ok := r.data[barcode]
	if !ok {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (r *InMemoryProductRepo) GetAllProducts() []*entities.Product {
	products := []*entities.Product{}
	for _, p := range r.data {
		products = append(products, p)
	}
	return products
}