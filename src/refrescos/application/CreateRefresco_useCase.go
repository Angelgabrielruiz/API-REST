package application

import "demo/src/refrescos/domain"

type CreateRefresco struct {
	db domain.IRefrescoRepository
}

func NewCreateRefresco(db domain.IRefrescoRepository) *CreateRefresco {
	return &CreateRefresco{db: db}
}

// Recibir name y price como parámetros
func (cp *CreateRefresco) Execute(name string, price float32) {
	cp.db.Save(name, price)
}