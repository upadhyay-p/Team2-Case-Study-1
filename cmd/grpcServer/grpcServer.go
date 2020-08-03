package grpcServer

import (
	CustomerModels "Team2CaseStudy1/pkg/Customer/Models"
	CustomerServices "Team2CaseStudy1/pkg/Customer/Services"
	OrderModels "Team2CaseStudy1/pkg/Order/Models"
	OrderServices "Team2CaseStudy1/pkg/Order/Services"
	"Team2CaseStudy1/pkg/OrderProto/orderpb"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"google.golang.org/grpc"
)

// var customerTable []CustomerModels.Customer
var db *dynamodb.DynamoDB
var allCustomers []*orderpb.Customer
var allOrders []*orderpb.Order

func InitializeDB() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("us-east-1"),
	}))
	db = dynamodb.New(sess)
}

type Server struct{}

// fetch orders from db and give it as response to client
func (*Server) GetOrders(ctx context.Context, req *orderpb.NoParamRequest) (*orderpb.OrderResponse, error) {
	fmt.Println("GetOrders Function called... ")
	res := &orderpb.OrderResponse{Res: allOrders}
	return res, nil
}

// fetch customers from db and give it as response to client
func (*Server) GetCustomers(ctx context.Context, req *orderpb.NoParamRequest) (*orderpb.CustomerAllResponse, error) {
	fmt.Println("GetCustomers Function called... ")
	res := &orderpb.CustomerAllResponse{Res: allCustomers}
	return res, nil
}

// fetch restaurants from db and give it as response to client
func (*Server) GetRestaurants(ctx context.Context, req *orderpb.NoParamRequest) (*orderpb.RestaurantResponse, error) {
	fmt.Println("GetRestaurants Function called... ")
	res := &orderpb.RestaurantResponse{DummyRes: "Hi this is a test call"}
	return res, nil
}

func (*Server) GetACustomer(ctx context.Context, req *orderpb.SpecificCustomerRequest) (*orderpb.SpecificCustomerResponse, error) {

	fmt.Println("GetACustomer Function called... ")

	customerid := req.GetCustId()

	customerDetails := CustomerServices.GetSpecificCustomerDetails(db, customerid)

	res := &orderpb.SpecificCustomerResponse{Res: customerDetails}

	return res, nil

}

func (*Server) GetAnOrder(ctx context.Context, req *orderpb.SpecificOrderRequest) (*orderpb.SpecificOrderResponse, error) {

	fmt.Println("GetACustomer Function called... ")

	orderid := req.GetOrderId()
	orderDetails := OrderServices.GetSpecificOrderDetails(db, orderid)

	res := &orderpb.SpecificOrderResponse{Res: orderDetails}

	return res, nil

}

// add order to db
func (*Server) AddOrder(ctx context.Context, req *orderpb.OrderRequest) (*orderpb.OrderPostResponse, error) {
	fmt.Println("AddOrders Function called... ")

	orderid := req.Ord.GetOrderId()
	customerid := req.Ord.GetCustomerId()
	restaurantid := req.Ord.GetRestaurantId()
	itemlist := req.Ord.ItemLine
	price := req.Ord.GetPrice()
	discount := req.Ord.GetDiscount()

	var itemline []*orderpb.Item
	var itemlinestruct []OrderModels.Item

	for i := range itemlist {
		itemname := itemlist[i].GetName()
		itemprice := itemlist[i].GetPrice()
		itemline = append(itemline, &orderpb.Item{
			Name:  itemname,
			Price: itemprice,
		})
		itemlinestruct = append(itemlinestruct, OrderModels.Item{
			Name:  itemname,
			Price: itemprice,
		})
	}

	orderDetails := OrderModels.Order{
		OrderId:      orderid,
		CustomerId:   customerid,
		RestaurantId: restaurantid,
		ItemLine:     itemlinestruct,
		Price:        price,
		Discount:     discount,
	}

	allOrders = append(allOrders, &orderpb.Order{
		OrderId:      orderid,
		CustomerId:   customerid,
		RestaurantId: restaurantid,
		ItemLine:     itemline,
		Price:        price,
		Discount:     discount,
	})

	res := &orderpb.OrderPostResponse{Res: &orderpb.Order{
		OrderId:      orderid,
		CustomerId:   customerid,
		RestaurantId: restaurantid,
		ItemLine:     itemline,
		Price:        price,
		Discount:     discount,
	}}

	OrderServices.AddOrderDetails(db, orderDetails)

	return res, nil
}

// add customer to db
func (*Server) AddCustomer(ctx context.Context, req *orderpb.CustomerRequest) (*orderpb.CustomerPostResponse, error) {
	fmt.Println("AddCustomer Function called... ")

	customerid := req.Cust.GetCustomerId()
	name := req.Cust.GetName()
	address := req.Cust.GetAddress()
	phone := req.Cust.GetPhone()

	res := &orderpb.CustomerPostResponse{Res: &orderpb.Customer{
		CustomerId: customerid,
		Name:       name,
		Address:    address,
		Phone:      phone,
	},
	}

	customerItem := CustomerModels.Customer{
		CustomerId: customerid,
		Name:       name,
		Address:    address,
		Phone:      phone,
	}
	allCustomers = append(allCustomers, &orderpb.Customer{
		CustomerId: customerid,
		Name:       name,
		Address:    address,
		Phone:      phone,
	})

	CustomerServices.AddCustomerDetails(db, customerItem)

	return res, nil
}

// add restaurant to db
func (*Server) AddRestaurant(ctx context.Context, req *orderpb.RestaurantRequest) (*orderpb.PostResponse, error) {
	fmt.Println("AddRestaurant Function called... ")
	res := &orderpb.PostResponse{Res: "Hi this is a test call"}
	return res, nil
}

func main() {
	fmt.Println("Hello from grpc server.")

	InitializeDB()

	allCustomers = CustomerServices.FetchCustomerTable(db)
	allOrders = OrderServices.FetchOrderTable(db)

	lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		log.Fatalf("Sorry failed to load server %v:", err)
	}

	s := grpc.NewServer()

	orderpb.RegisterQueryServiceServer(s, &Server{})

	if s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
