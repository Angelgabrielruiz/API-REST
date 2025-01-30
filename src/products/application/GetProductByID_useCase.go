package application

import "demo/src/products/domain"

type GetProductByID struct {
	db domain.IProduct
}

func NewGetProductByID(db domain.IProduct) *GetProductByID {
	return &GetProductByID{db: db}
}

func (gp *GetProductByID) Execute(id int) {
	gp.db.GetByID(id)
}