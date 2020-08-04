package Models

type Customer struct {
	ID      int64
	Name    string
	Address string
	Phone   string
}

type Item struct {
	Name  string
	Price float32
}

type Order struct {
	ID       int64
	C_ID     int64
	R_ID     int64
	ItemLine []Item
	Price    float32
	Discount int64
}

type Restaurant struct {
	ID       int64
	Name     string
	Online   bool
	Menu     []Item
	Rating   float32
	Category string
}
