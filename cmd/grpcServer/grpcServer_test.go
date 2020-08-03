package grpcServer_test

import (
	grpcServer "Team2CaseStudy1/cmd/grpcServer"
	"Team2CaseStudy1/pkg/OrderProto/orderpb"
	"context"
	"testing"
	"log"
	"net"
	"reflect"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"


)

const bufSize = 1024*1024

var lis *bufconn.Listener

func init(){
	grpcServer.InitializeDB()
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	orderpb.RegisterQueryServiceServer(s, &grpcServer.Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAddCustomerPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderpb.NewQueryServiceClient(conn)
	req := &orderpb.CustomerRequest{Cust: &orderpb.Customer{
		CustomerId: 9000,
		Name:       "Priya",
		Address:    "Address1",
		Phone:      "12345",
	}}
	res, err := oc.AddCustomer(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling AddCustomer : %v ", err)
	}
	res1 := &orderpb.CustomerPostResponse{Res: &orderpb.Customer{
		CustomerId: 9000,
		Name:       "Priya",
		Address:    "Address1",
		Phone:      "12345",
	},
	}
	if reflect.DeepEqual(res.Res,res1.Res)==false{
		t.Fatalf("Test failed, expected value- %v, actual value- %v",res1.Res, res.Res)
	}

}

func TestGetASpeceficCustomerPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderpb.NewQueryServiceClient(conn)
	var customerid int64 = 9000

	req := &orderpb.SpecificCustomerRequest{CustId: customerid}
	res, err := oc.GetACustomer(ctx, req)

	if err != nil {
		t.Fatalf("Error While calling GetACustomer : %v ", err)
	}
	res1 := &orderpb.SpecificCustomerResponse{Res: &orderpb.Customer{
		CustomerId: 9000,
		Name:       "Priya",
		Address:    "Address1",
		Phone:      "12345",
	},
	}
	if reflect.DeepEqual(res.Res,res1.Res)==false{
		t.Fatalf("Test failed, expected value- %v, actual value- %v",res1.Res, res.Res)
	}

}

//func TestAddCustomerFail(t *testing.T) {
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("failed to dial: %v", err)
//	}
//	defer conn.Close()
//	oc := orderpb.NewQueryServiceClient(conn)
//	req := &orderpb.CustomerRequest{Cust: &orderpb.Customer{
//		CustomerId: 9000,
//		Name:       "Priya",
//		Address:    "Address1",
//		Phone:      "12345",
//	}}
//	res, err := oc.AddCustomer(context.Background(), req)
//	if err != nil {
//		t.Fatalf("Error While calling AddCustomer : %v ", err)
//	}
//	res1 := &orderpb.CustomerPostResponse{Res: &orderpb.Customer{
//		CustomerId: 9001,
//		Name:       "Priya",
//		Address:    "Address1",
//		Phone:      "12345",
//	},
//	}
//	if reflect.DeepEqual(res.Res,res1.Res)==true{
//		t.Fatalf("Test failed, expected value- %v, actual value- %v",res1.Res, res.Res)
//	}
//
//}

func TestAddOrderPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderpb.NewQueryServiceClient(conn)
	req := &orderpb.OrderRequest{Ord: &orderpb.Order{
		OrderId:      9000,
		CustomerId:   9000,
		RestaurantId: 12,
		ItemLine:     []*orderpb.Item{&orderpb.Item{Name:"item1",Price:64},&orderpb.Item{Name:"item2",Price:94}, },
		Price:        168,
		Discount:     0,
	}}
	res, err := oc.AddOrder(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling AddCustomer : %v ", err)
	}
	res1 := &orderpb.OrderPostResponse{Res: &orderpb.Order{
		OrderId:      9000,
		CustomerId:   9000,
		RestaurantId: 12,
		ItemLine:     []*orderpb.Item{&orderpb.Item{Name:"item1",Price:64},&orderpb.Item{Name:"item2",Price:94}, },
		Price:        168,
		Discount:     0,
	}}
	if reflect.DeepEqual(res.Res,res1.Res)==false{
		t.Fatalf("Test failed, expected value- %v, actual value- %v",res1.Res, res.Res)
	}

}

func TestGetASpeceficOrderPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderpb.NewQueryServiceClient(conn)
	var orderid int64 = 9000

	req := &orderpb.SpecificOrderRequest{OrderId: orderid}
	res, err := oc.GetAnOrder(ctx, req)

	if err != nil {
		t.Fatalf("Error While calling GetACustomer : %v ", err)
	}
	res1 := &orderpb.SpecificOrderResponse{Res: &orderpb.Order{
		OrderId:      9000,
		CustomerId:   9000,
		RestaurantId: 12,
		ItemLine:     []*orderpb.Item{&orderpb.Item{Name:"item1",Price:64},&orderpb.Item{Name:"item2",Price:94}, },
		Price:        168,
		Discount:     0,
	},
	}
	if reflect.DeepEqual(res.Res,res1.Res)==false{
		t.Fatalf("Test failed, expected value- %v, actual value- %v",res1.Res, res.Res)
	}

}