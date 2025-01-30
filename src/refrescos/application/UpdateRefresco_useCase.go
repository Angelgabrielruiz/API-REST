package application

import "demo/src/refrescos/domain"

type UpdateRefresco struct {
	db domain.IRefrescoRepository
}

func NewUpdateRefresco(db domain.IRefrescoRepository) *UpdateRefresco {
	return &UpdateRefresco{db: db}
}

func (up *UpdateRefresco) Execute(id int, name string, price float32) {
	up.db.Update(id, name, price)
}