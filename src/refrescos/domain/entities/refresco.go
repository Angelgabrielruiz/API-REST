package entities

type Refresco struct {
	ID    int32
	Name  string
	Price float32
}

func NewRefresco(name string, price float32) *Refresco {
	return &Refresco{Name: name, Price: price}
}