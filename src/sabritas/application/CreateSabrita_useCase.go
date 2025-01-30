package application

import "demo/src/sabritas/domain"

type CreateSabrita struct {
	db domain.ISabritaRepository
}

func NewCreateSabrita(db domain.ISabritaRepository) *CreateSabrita {
	return &CreateSabrita{db: db}
}

// Recibir name y price como parámetros
func (cp *CreateSabrita) Execute(name string, price float32) {
	cp.db.Save(name, price)
}