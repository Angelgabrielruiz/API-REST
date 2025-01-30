package domain

type IProduct interface {
	Save()
	GetAll()
	GetByID(id int)
	Update(id int, name string, price float32)
	Delete(id int)
}