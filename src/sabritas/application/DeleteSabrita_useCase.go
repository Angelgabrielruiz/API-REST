package application

import "demo/src/sabritas/domain"

type Deletesabrita struct {
	db domain.ISabritaRepository
}

func NewDeleteSabrita(db domain.ISabritaRepository) *Deletesabrita {
	return &Deletesabrita{db: db}
}

func (ds *Deletesabrita) Execute(id int) {
	ds.db.Delete(id)
}