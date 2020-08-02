package test_test

import (
	t "Team2CaseStudy1/cmd/grpcServer"
	"Team2CaseStudy1/pkg/OrderProto/orderpb"
	"context"
	"testing"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

)

const bufSize = 1024*1024

var lis *bufconn.Listener

func init(){
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	orderpb.RegisterQueryServiceServer(s, &t.Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGetOrdersPass(t *testing.T) {
	var allOrders []*orderpb.Order
	var itemline []*orderpb.Item

	itemline = append(itemline, &orderpb.Item{
		Name:  "item-1",
		Price: 30,
	})

	itemline = append(itemline, &orderpb.Item{
		Name:  "item-2",
		Price: 305,
	})
	allOrders = append(allOrders, &orderpb.Order{
		OrderId:      1,
		CustomerId:   1,
		RestaurantId: 1,
		ItemLine:     itemline,
		Price:        4,
		Discount:     0,
	})
	allOrders = append(allOrders, &orderpb.Order{
		OrderId:      2,
		CustomerId:   1,
		RestaurantId: 1,
		ItemLine:     itemline,
		Price:        4,
		Discount:     0,
	})
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderpb.NewQueryServiceClient(conn)
	req := &orderpb.NoParamRequest{}
	_, err = oc.GetOrders(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}

}