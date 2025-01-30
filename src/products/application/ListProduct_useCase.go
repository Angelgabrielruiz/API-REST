package application

import "demo/src/products/domain"

type ListProduct struct {
	db domain.IProduct
}

func NewListProduct(db domain.IProduct) *ListProduct {
	return &ListProduct{db: db}
}

func (lp *ListProduct) Execute() {
	lp.db.GetAll()
}