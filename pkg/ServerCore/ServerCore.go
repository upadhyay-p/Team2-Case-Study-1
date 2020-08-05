package ServerCore

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/shashijangra22/Team2-Case-Study-1/pkg/Models"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/customer"
	CustomerServices "github.com/shashijangra22/Team2-Case-Study-1/pkg/customer/services"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/order"
	OrderServices "github.com/shashijangra22/Team2-Case-Study-1/pkg/order/services"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/restaurant"
	RestaurantServices "github.com/shashijangra22/Team2-Case-Study-1/pkg/restaurant/services"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Server struct{}

var DB *dynamodb.DynamoDB

// AWS STRUCT to save credentials
type AWS_STRUCT struct {
	AWS_KEY_ID     string
	AWS_SECRET_KEY string
	REGION         string
}

var secret AWS_STRUCT

// initialises the DB with credentials stored in filename
func CreateDBSession(filename string) *dynamodb.DynamoDB {
	secretsFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening secrets.json! from path: ", filename)
		os.Exit(1)
	}
	defer secretsFile.Close()
	byteValue, _ := ioutil.ReadAll(secretsFile)
	json.Unmarshal(byteValue, &secret)
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(secret.REGION),
		// Endpoint:    aws.String("http://localhost:8000"), // [for local should be http] remove this line to connect to cloud dynamodDB with creds in secrets.json file
		Credentials: credentials.NewStaticCredentials(secret.AWS_KEY_ID, secret.AWS_SECRET_KEY, ""),
	}))
	db := dynamodb.New(sess)
	return db
}

func init() {
	path, _ := os.Getwd()
	DB = CreateDBSession(path + "/secrets.json")
}

// gRPC Core method implementing the AddCustomer interface
func (*Server) AddCustomer(ctx context.Context, req *customer.Customer) (*customer.Customer, error) {
	fmt.Println("AddCustomer Function called... ")
	id := req.GetID()
	name := req.GetName()
	address := req.GetAddress()
	phone := req.GetPhone()
	cst := Models.Customer{ID: id, Name: name, Address: address, Phone: phone}
	CustomerServices.AddOne(DB, cst)
	return req, nil
}

// gRPC Core method implementing the GetCustomer interface
func (*Server) GetCustomer(ctx context.Context, req *customer.IDRequest) (*customer.Customer, error) {
	fmt.Println("GetCustomer is called... ")
	id := req.GetID()
	res := CustomerServices.GetOne(DB, id)
	return res, nil
}

// gRPC Core method implementing the GetCustomers interface
func (*Server) GetCustomers(ctx context.Context, req *customer.NoParamRequest) (*customer.Customers, error) {
	fmt.Println("GetCustomers is called... ")
	allCustomers := CustomerServices.GetAll(DB)
	res := &customer.Customers{Customers: allCustomers}
	return res, nil
}

// gRPC Core method implementing the AddOrder interface
func (*Server) AddOrder(ctx context.Context, req *order.Order) (*order.Order, error) {
	fmt.Println("AddOrder Function called... ")
	itemlist := req.GetItemLine()
	var items []Models.Item
	for i := range itemlist {
		items = append(items, Models.Item{
			Name:  itemlist[i].GetName(),
			Price: itemlist[i].GetPrice(),
		})
	}
	orderDetails := Models.Order{ID: req.GetID(), C_ID: req.GetC_ID(), ItemLine: items, Price: req.GetPrice(), Discount: req.GetDiscount()}
	OrderServices.Add(DB, orderDetails)
	return req, nil
}

// gRPC Core method implementing the GetOrder interface
func (*Server) GetOrder(ctx context.Context, req *order.IDRequest) (*order.Order, error) {
	fmt.Println("GetOrder Function called... ")
	id := req.GetID()
	res := OrderServices.GetOne(DB, id)
	return res, nil
}

// gRPC Core method implementing the GetOrders interface
func (*Server) GetOrders(ctx context.Context, req *order.NoParamRequest) (*order.Orders, error) {
	fmt.Println("GetOrders is called...")
	allOrders := OrderServices.GetAll(DB)
	res := &order.Orders{Orders: allOrders}
	return res, nil
}

// gRPC Core method implementing the GetRestaurant interface
func (*Server) GetRestaurant(ctx context.Context, req *restaurant.IDRequest) (*restaurant.Restaurant, error) {
	fmt.Println("GetRestaurant Function called... ")
	id := req.GetID()
	res := RestaurantServices.GetOne(DB, id)
	return res, nil
}

// gRPC Core method implementing the GetRestaurants interface
func (*Server) GetRestaurants(ctx context.Context, req *restaurant.NoParamRequest) (*restaurant.Restaurants, error) {
	fmt.Println("GetRestaurants is called...")
	allRestaurants := RestaurantServices.GetAll(DB)
	res := &restaurant.Restaurants{Restaurants: allRestaurants}
	return res, nil
}
