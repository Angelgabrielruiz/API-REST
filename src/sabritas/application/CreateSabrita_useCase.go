package application

import "demo/src/sabritas/domain"

type CreateSabrita struct {
	db domain.ISabritaRepository
}

func NewCreateSabrita(db domain.ISabritaRepository) *CreateSabrita {
	return &CreateSabrita{db: db}
}


func (cp *CreateSabrita) Execute(name string, price float32) {
	cp.db.Save(name, price)
}