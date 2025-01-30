package domain

type IRefrescoRepository interface {
	Save(name string, price float32)
	GetAll() ([]map[string]interface{}, error)
	GetByID(id int) (map[string]interface{}, error)
	Update(id int, name string, price float32)
	Delete(id int)
}