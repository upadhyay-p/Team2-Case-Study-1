package Models

type Item struct {
	Name  string
	Price string
}

type Order struct {
	OrderId      string
	CustomerId   string
	RestaurantId string
	ItemLine     []Item
	Price        string
	Discount     string
}
