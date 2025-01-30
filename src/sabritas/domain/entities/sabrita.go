package entities

type Sabrita struct {
	ID    int32
	Name  string
	Price float32
}

func NewSabrita(name string, price float32) *Sabrita {
	return &Sabrita{Name: name, Price: price}
}