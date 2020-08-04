package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"log"
	"context"
	"net"
	"io"
	"strings"

	"Team2CaseStudy1/pkg/ServerHelper"
	"Team2CaseStudy1/pkg/OrderProto/orderpb"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/test/bufconn"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const bufSize = 1024*1024
var lis *bufconn.Listener
var router *gin.Engine

func init(){

	//setup mock grpc server
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	orderpb.RegisterQueryServiceServer(s, &ServerHelper.Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	//start db session at server
	ServerHelper.InitDB()

	//setup api routes
	router = SetupRoutes()
}


func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func performRequest(r http.Handler, body io.Reader, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func GetAnOrderItem() io.Reader{
	Ord := "{\"OrderId\":9000,\"CustomerId\":437,\"RestaurantId\":28,\"ItemLine\":[{\"Name\":\"Muskox - French Rack\",\"Price\":500},{\"Name\":\"Sugar - Invert\",\"Price\":500}],\"Price\":1000,\"Discount\":60}"
	r := strings.NewReader(Ord)
	return r
}
func GetACustomerItem() io.Reader{
	Cust := "{\"CustomerId\":9001,\"Name\":\"Constantino\",\"Phone\":\"864-515-2235\",\"Address\":\"44 Mayfield Avenue\"}"
	r := strings.NewReader(Cust)
	return r
}

func TestGetAllOrders(t *testing.T) {
	// Perform a GET request with that handler.
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	queryServiceClient = orderpb.NewQueryServiceClient(conn)
	w := performRequest(router, nil,"GET", "/api/orders")

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestGetSpecificOrderPass(t *testing.T) {
	// Perform a GET request with that handler.
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	queryServiceClient = orderpb.NewQueryServiceClient(conn)
	w := performRequest(router,nil, "GET", "/api/order/1")

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestPostOrderPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	var order io.Reader = GetAnOrderItem()
	queryServiceClient = orderpb.NewQueryServiceClient(conn)
	w := performRequest(router, order,"POST", "/api/new-order")
	assert.Equal(t, http.StatusOK, w.Code)

}
func TestPostOrderFail(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	var order io.Reader = strings.NewReader("")
	queryServiceClient = orderpb.NewQueryServiceClient(conn)
	w := performRequest(router, order,"POST", "/api/new-order")
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetAllCustomers(t *testing.T) {
	// Perform a GET request with that handler.
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	queryServiceClient = orderpb.NewQueryServiceClient(conn)
	w := performRequest(router, nil,"GET", "/api/customers")

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestGetSpecificCustomerPass(t *testing.T) {
	// Perform a GET request with that handler.
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	queryServiceClient = orderpb.NewQueryServiceClient(conn)
	w := performRequest(router, nil,"GET", "/api/customer/1")

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestPostCustomer(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	customer := GetACustomerItem()
	queryServiceClient = orderpb.NewQueryServiceClient(conn)
	w := performRequest(router, customer,"POST", "/api/new-customer")
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetAllRestaurants(t *testing.T) {
	// Perform a GET request with that handler.
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	queryServiceClient = orderpb.NewQueryServiceClient(conn)
	w := performRequest(router, nil,"GET", "/api/restaurants")

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestWrongAPIRoute(t *testing.T) {
	// Perform a GET request with that handler.
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	queryServiceClient = orderpb.NewQueryServiceClient(conn)
	w := performRequest(router,nil, "GET", "/api/customer1")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusNotFound, w.Code)
}
