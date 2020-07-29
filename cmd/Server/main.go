package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"Team2CaseStudy1/pkg/AvgPrice"
	"Team2CaseStudy1/pkg/Models"
	"Team2CaseStudy1/pkg/TopRestauBuyers"
	"Team2CaseStudy1/pkg/order/orderProto"

	"google.golang.org/grpc"
)

var data []Models.Order
var byteValue []byte
var fname string

type server struct{}

func (*server) CreateOrder(ctx context.Context, req *orderProto.OrderRequest) (*orderProto.OrderResponse, error) {
	fmt.Println("Function called... ")
	var allOrders []*orderProto.OrderStruct
	for i := range data {

		var items []*orderProto.OrderStruct_Item
		for j := range data[i].ItemLine {
			items = append(items, &orderProto.OrderStruct_Item{Name: data[i].ItemLine[j].Name, Price: data[i].ItemLine[j].Price, Quantity: data[i].ItemLine[j].Quantity})
		}
		allOrders = append(allOrders, &orderProto.OrderStruct{
			OrderID:    data[i].OrderID,
			CustomerID: data[i].CustomerID,
			Restaurant: data[i].Restaurant,
			ItemLine:   items,
			Price:      data[i].Price,
			Quantity:   data[i].Quantity,
			Discount:   data[i].Discount,
			Date:       data[i].Date,
		})
	}

	res := &orderProto.OrderResponse{OrdRes: allOrders}
	return res, nil

}

func (*server) GetAvgPricesOrders(ctx context.Context, req *orderProto.AvgPriceInfoRequest) (*orderProto.AvgPriceInfoResponse, error) {
	fmt.Println("AvgPrice Function called... ")
	//var obj *orderProto.OrderRequest
	avgPrices := AvgPrice.INIT(strings.TrimSpace(fname))
	var allPrices []*orderProto.AvgPriceInfo
	for i := range avgPrices {
		allPrices = append(allPrices, &orderProto.AvgPriceInfo{
			CustomerID: avgPrices[i].CustomerID,
			AvgPrice:   float32(avgPrices[i].AvgPrice),
			AvgOrders:  avgPrices[i].AvgOrders,
		})
	}

	res := &orderProto.AvgPriceInfoResponse{Res: allPrices}
	return res, nil
}

func (*server) GetTopCustomers(ctx context.Context, req *orderProto.TopCustomersRequest) (*orderProto.TopCustomersResponse, error) {
	fmt.Println("TopCustomer Function called... ")
	numberOfBuyers := req.GetNum()
	topCustomersList := TopRestauBuyers.FindTopBuyers(byteValue, numberOfBuyers)
	//fmt.Println(topCustomersList)
	var allCust []*orderProto.TopCustomer
	for i := range topCustomersList {
		allCust = append(allCust, &orderProto.TopCustomer{
			CustomerID:  topCustomersList[i].CustomerID,
			Expenditure: float32(topCustomersList[i].Expenditure),
		})
	}
	res := &orderProto.TopCustomersResponse{Res: allCust}
	return res, nil
}

func (*server) GetTopRest(ctx context.Context, req *orderProto.TopRestaurantsRequest) (*orderProto.TopRestaurantsResponse, error) {
	fmt.Println("TopRest Function called... ")
	numberOfRestaurants := req.GetNum()
	topRestaurantsList := TopRestauBuyers.FindTopRestaurants(byteValue, numberOfRestaurants)
	var allRest []*orderProto.TopRest
	for i := range topRestaurantsList {
		allRest = append(allRest, &orderProto.TopRest{
			Restaurant: topRestaurantsList[i].Restaurant,
			Revenue:    float32(topRestaurantsList[i].Revenue),
		})
	}
	res := &orderProto.TopRestaurantsResponse{Res: allRest}
	return res, nil
}

func (*server) PostOrder(ctx context.Context, req *orderProto.PostRequest) (*orderProto.PostResponse, error) {
	fmt.Println("PostOrder Function called... ")
	orderID := req.Res.GetOrderID()
	customerID := req.Res.GetCustomerID()
	restaurant := req.Res.GetRestaurant()
	itemLine := req.Res.GetItemLine()
	price := req.Res.GetPrice()
	quantity := req.Res.GetQuantity()
	discount := req.Res.GetDiscount()
	date := req.Res.GetDate()

	res := &orderProto.PostResponse{Res: &orderProto.OrderStruct{
		OrderID:    orderID,
		CustomerID: customerID,
		Restaurant: restaurant,
		ItemLine:   itemLine,
		Price:      price,
		Quantity:   quantity,
		Discount:   discount,
		Date:       date,
	},
	}
	var items []Models.Item
	for i := range itemLine {
		items = append(items, Models.Item{
			Name:     itemLine[i].GetName(),
			Price:    itemLine[i].GetPrice(),
			Quantity: itemLine[i].GetQuantity(),
		})
	}

	NewOrder := Models.Order{
		OrderID:    orderID,
		CustomerID: customerID,
		Restaurant: restaurant,
		ItemLine:   items,
		Price:      price,
		Quantity:   quantity,
		Discount:   discount,
		Date:       date,
	}
	data = append(data, NewOrder)
	toJSON()

	fmt.Println("New Entry Added in the output file")

	return res, nil

}

//To update the json file  and the byteValue slice
func toJSON() {
	byteValue, _ = json.MarshalIndent(data, "", "	  ")
	err := ioutil.WriteFile(fname, byteValue, 0644)
	if err != nil {
		fmt.Println("Error in writing the file")
	}
	fmt.Println("Output file is stored as: " + fname)
}

func main() {

	fmt.Println("Looking for .json file to INIT DB...")
	args := os.Args[1:]
	if len(args) == 1 {
		fname = args[0]
		byteValue, _ = ioutil.ReadFile(fname)
		err := json.Unmarshal(byteValue, &data)
		fmt.Println("DB is initialised as: ", fname)
		lis, err := net.Listen("tcp", "0.0.0.0:5051")
		if err != nil {
			log.Fatalf("Sorry failed to load server :() %v:", err)
		}

		s := grpc.NewServer()
		fmt.Println("Hello from grpc server :)")

		orderProto.RegisterOrderServer(s, &server{})

		if s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve :( %v", err)
		}

	} else {
		println("Please Give File name as an Argument!")
	}
}
