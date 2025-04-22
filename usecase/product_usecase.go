package usecase

import (
	"gomentoring/entities"
	"time"
	"math/rand"
)

type ProductUsecase struct {
	Repo entities.ProductRepository
}

func (uc *ProductUsecase) GetAllProducts() []*entities.Product {
	products := uc.Repo.GetAllProducts()
	return products
}

func NewProductUsecase(repo entities.ProductRepository) *ProductUsecase {
	return &ProductUsecase{Repo: repo}
}

func (uc *ProductUsecase) Create(p *entities.Product) error {
	return uc.Repo.Save(p)
}

func (uc *ProductUsecase) GetProductWithStock(barcode string) (*entities.Product, int, error) {
	productChan := make(chan *entities.Product)
	stockChan := make(chan int)
	errChan := make(chan error)

	go func() {
		product, err := uc.Repo.GetByBarCode(barcode)
		if err != nil {
			errChan <- err
			return
		}
		productChan <- product
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		stockChan <- rand.Intn(100)
	}()

	var (
		product *entities.Product
		stock   int
	)

	for i := 0; i < 2; i++ {
		select {
		case p := <-productChan:
			product = p
		case s := <-stockChan:
			stock = s
		case err := <-errChan:
			return nil, 0, err
		}
	}

	return product, stock, nil
}