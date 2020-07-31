package Models

type Item struct {
	Name     string
	Price    float64
}


type Rest struct {
	ID    int64
	Name string
	Availability bool
	Items   []Item
	Rating   float64
	Category  string
}
