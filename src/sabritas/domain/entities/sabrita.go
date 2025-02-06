package entities

// Sabrita representa una entidad de sabritas (snacks)
type Sabrita struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}