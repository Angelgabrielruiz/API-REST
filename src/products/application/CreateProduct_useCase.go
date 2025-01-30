package application

import "demo/src/products/domain"

type CreateProduct struct {
	db domain.IProduct
}

func NewCreateProduct(db domain.IProduct) *CreateProduct {
	return &CreateProduct{db: db}
}

func (cp *CreateProduct) Execute() {
	cp.db.Save()
}