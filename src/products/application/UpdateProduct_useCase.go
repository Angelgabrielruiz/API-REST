package application

import "demo/src/products/domain"

type UpdateProduct struct {
	db domain.IProduct
}

func NewUpdateProduct(db domain.IProduct) *UpdateProduct {
	return &UpdateProduct{db: db}
}

func (up *UpdateProduct) Execute(id int, name string, price float32) {
	up.db.Update(id, name, price)
}