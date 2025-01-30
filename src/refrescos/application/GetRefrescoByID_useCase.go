package application

import "demo/src/refrescos/domain"

type GetRefrescoByID struct {
	db domain.IRefrescoRepository
}

func NewGetRefrescoByID(db domain.IRefrescoRepository) *GetRefrescoByID {
	return &GetRefrescoByID{db: db}
}

func (gp *GetRefrescoByID) Execute(id int) (map[string]interface{}, error) {
	return gp.db.GetByID(id)
}