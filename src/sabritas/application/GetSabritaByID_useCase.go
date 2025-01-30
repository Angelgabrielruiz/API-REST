package application

import "demo/src/sabritas/domain"

type GetSabritaByID struct {
	db domain.ISabritaRepository
}

func NewGetSabritaByID(db domain.ISabritaRepository) *GetSabritaByID {
	return &GetSabritaByID{db: db}
}

func (gs *GetSabritaByID) Execute(id int) (map[string]interface{}, error) {
	return gs.db.GetByID(id)
}