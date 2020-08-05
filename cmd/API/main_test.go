package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/ServerCore"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/customer"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/order"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/restaurant"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener
var router *gin.Engine

func init() {

	//setup mock grpc server
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	// orderpb.RegisterQueryServiceServer(s, &ServerHelper.Server{})
	customer.RegisterCustomerServiceServer(s, &ServerCore.Server{})
	order.RegisterOrderServiceServer(s, &ServerCore.Server{})
	restaurant.RegisterRestaurantServiceServer(s, &ServerCore.Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
	//setup api routes
	router = SetupRoutes(false) // called with flag = false to disable authentication
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
func GetAnOrderItem() io.Reader {
	Ord := "{\"OrderId\":9000,\"CustomerId\":437,\"RestaurantId\":28,\"ItemLine\":[{\"Name\":\"Muskox - French Rack\",\"Price\":500},{\"Name\":\"Sugar - Invert\",\"Price\":500}],\"Price\":1000,\"Discount\":60}"
	r := strings.NewReader(Ord)
	return r
}
func GetACustomerItem() io.Reader {
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
	order.OSC = order.NewOrderServiceClient(conn)
	w := performRequest(router, nil, "GET", "/api/orders")

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestGetOneOrderPass(t *testing.T) {
	// Perform a GET request with that handler.
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	order.OSC = order.NewOrderServiceClient(conn)
	w := performRequest(router, nil, "GET", "/api/order/1")

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestAddOrderPass(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	var ord io.Reader = GetAnOrderItem()
	order.OSC = order.NewOrderServiceClient(conn)
	w := performRequest(router, ord, "POST", "/api/order")
	assert.Equal(t, http.StatusCreated, w.Code)

}
func TestAddOrderFail(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	var ord io.Reader = strings.NewReader("")
	order.OSC = order.NewOrderServiceClient(conn)
	w := performRequest(router, ord, "POST", "/api/order")
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllCustomers(t *testing.T) {
	// Perform a GET request with that handler.
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	customer.CSC = customer.NewCustomerServiceClient(conn)
	w := performRequest(router, nil, "GET", "/api/customers")

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestGetACustomerPass(t *testing.T) {
	// Perform a GET request with that handler.
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	customer.CSC = customer.NewCustomerServiceClient(conn)
	w := performRequest(router, nil, "GET", "/api/customer/1")

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestAddCustomer(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cst := GetACustomerItem()
	customer.CSC = customer.NewCustomerServiceClient(conn)
	w := performRequest(router, cst, "POST", "/api/customer")
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllRestaurants(t *testing.T) {
	// Perform a GET request with that handler.
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	restaurant.RSC = restaurant.NewRestaurantServiceClient(conn)
	w := performRequest(router, nil, "GET", "/api/restaurants")
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
	customer.CSC = customer.NewCustomerServiceClient(conn)
	w := performRequest(router, nil, "GET", "/api/customer1")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusNotFound, w.Code)
}
