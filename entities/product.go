package entities

type Product struct {
	Id int
	Name string
	Description string
	BarCode string
	Price int
}

type ProductRepository interface {
	GetByBarCode(barcode string) (*Product, error)
	Save(product *Product) error
	GetAllProducts() []*Product
}
