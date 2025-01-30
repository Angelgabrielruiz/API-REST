package application

import "demo/src/sabritas/domain"

type ListSabrita struct {
	db domain.ISabritaRepository
}

func NewListSabrita(db domain.ISabritaRepository) *ListSabrita {
	return &ListSabrita{db: db}
}

func (ls *ListSabrita) Execute() ([]map[string]interface{}, error) {
	return ls.db.GetAll()
}