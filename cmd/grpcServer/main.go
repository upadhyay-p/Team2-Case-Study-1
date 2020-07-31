package main

import (
	"Team2CaseStudy1/pkg/OrderProto/orderpb"
	"context"
	"fmt"

	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

// fetch orders from db and give it as response to client
func (*server) GetOrders(ctx context.Context, req *orderpb.NoParamRequest) (*orderpb.OrderResponse, error) {
	fmt.Println("GetOrders Function called... ")
	res := &orderpb.OrderResponse{DummyRes: "Hi this is a test call"}
	return res, nil
}

// fetch customers from db and give it as response to client
func (*server) GetCustomers(ctx context.Context, req *orderpb.NoParamRequest) (*orderpb.CustomerResponse, error) {
	fmt.Println("GetCustomers Function called... ")
	res := &orderpb.CustomerResponse{DummyRes: "Hi this is a test call"}
	return res, nil
}

// fetch restaurants from db and give it as response to client
func (*server) GetRestaurants(ctx context.Context, req *orderpb.NoParamRequest) (*orderpb.RestaurantResponse, error) {
	fmt.Println("GetRestaurants Function called... ")
	res := &orderpb.RestaurantResponse{DummyRes: "Hi this is a test call"}
	return res, nil
}

// add order to db
func (*server) AddOrder(ctx context.Context, req *orderpb.OrderRequest) (*orderpb.PostResponse, error) {
	fmt.Println("AddOrders Function called... ")
	res := &orderpb.PostResponse{DummyRes: "Hi this is a test call"}
	return res, nil
}

// add customer to db
func (*server) AddCustomer(ctx context.Context, req *orderpb.CustomerRequest) (*orderpb.PostResponse, error) {
	fmt.Println("AddCustomer Function called... ")
	res := &orderpb.PostResponse{DummyRes: "Hi this is a test call"}
	return res, nil
}

// add restaurant to db
func (*server) AddRestaurant(ctx context.Context, req *orderpb.RestaurantRequest) (*orderpb.PostResponse, error) {
	fmt.Println("AddRestaurant Function called... ")
	res := &orderpb.PostResponse{DummyRes: "Hi this is a test call"}
	return res, nil
}




func main() {
	fmt.Println("Hello from grpc server.")

	lis, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		log.Fatalf("Sorry failed to load server %v:", err)
	}

	s := grpc.NewServer()

	orderpb.RegisterQueryServiceServer(s, &server{})

	if s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}
