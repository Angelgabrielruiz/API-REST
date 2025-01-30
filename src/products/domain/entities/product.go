package entities

type product struct {
	id int32
	name string
	price float32
}

func NewProduct(name string, price float32) *product{
	return &product{id: 1, name: name, price: price}
}

func (p *product) GetName() string {
	return p.name
}

func (p *product) SetName(name string) {
	p.name = name
}