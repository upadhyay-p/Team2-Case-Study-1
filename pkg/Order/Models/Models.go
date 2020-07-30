package Models

type Item struct {
	Name     string
	Price    float32
	Quantity int64
}

type Order struct {
	OrderID    int64
	CustomerID int64
	Restaurant string
	ItemLine   []Item
	Price      float32
	Quantity   int64
	Discount   int64
	Date       string
}
