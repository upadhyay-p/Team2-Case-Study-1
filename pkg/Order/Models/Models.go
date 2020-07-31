package Models

type Item struct {
	Name  string
	Price float32
}

type Order struct {
	OrderId      int64
	CustomerId   int64
	RestaurantId int64
	ItemLine     []Item
	Price        float32
	Discount     int64
}
