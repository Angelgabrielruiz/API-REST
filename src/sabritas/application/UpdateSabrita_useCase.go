package application

import "demo/src/sabritas/domain"

type UpdateSabrita struct {
	db domain.ISabritaRepository
}

func NewUpdateSabrita(db domain.ISabritaRepository) *UpdateSabrita {
	return &UpdateSabrita{db: db}
}

func (up *UpdateSabrita) Execute(id int, name string, price float32) {
	up.db.Update(id, name, price)
}