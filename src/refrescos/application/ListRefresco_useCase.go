package application

import "demo/src/refrescos/domain"

type ListRefresco struct {
	db domain.IRefrescoRepository
}

func NewListRefresco(db domain.IRefrescoRepository) *ListRefresco {
	return &ListRefresco{db: db}
}

func (lp *ListRefresco) Execute() ([]map[string]interface{}, error) {
	return lp.db.GetAll()
}