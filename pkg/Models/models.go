package Models

// Data model of Customer
type Customer struct {
	ID      int64
	Name    string
	Address string
	Phone   string
}

// Data model of Item
type Item struct {
	Name  string
	Price float32
}

// Data model of Order
type Order struct {
	ID       int64
	C_ID     int64
	R_ID     int64
	ItemLine []Item
	Price    float32
	Discount int64
}

// Data model of Restaurant
type Restaurant struct {
	ID       int64
	Name     string
	Online   bool
	Menu     []Item
	Rating   float32
	Category string
}
