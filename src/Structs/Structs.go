package Structs


type Item struct {
	Name     string  `json:"Name"`
	Price    float64 `json:"Price"`
	Quantity int64   `json:"Quantity"`
}

type Order struct {
	OrderID    int64 `json: "orderId"`
	CustomerID int64 `json:"CustomerID"`
	Restaurant string `json:"Restaurant"`
	ItemLine   []Item `json:"ItemLine"`
	Price      float64 `json:"Price"`
	Quantity   int64 `json:"Quantity"`
	Discount   int64 `json:"Discount"`
	Date       string `json:"Date"`
}

type AvgPriceInfo struct{
	CustomerID int64
	AvgPrice float64
	AvgOrders int64
}