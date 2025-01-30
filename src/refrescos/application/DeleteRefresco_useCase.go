package application

import "demo/src/refrescos/domain"

type DeleteRefresco struct {
	db domain.IRefrescoRepository
}

func NewDeleteRefresco(db domain.IRefrescoRepository) *DeleteRefresco {
	return &DeleteRefresco{db: db}
}

func (dr *DeleteRefresco) Execute(id int) {
	dr.db.Delete(id)
}