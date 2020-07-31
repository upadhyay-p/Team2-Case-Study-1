package main

import (
	"Team2CaseStudy1/pkg/OrderProto/orderpb"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var queryServiceClient orderpb.QueryServiceClient

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Team 2": "Hello from Aadithya, Abhishek, Priya, Shashi!",
	})
}

// To get all the orders
func GetAllOrders(c *gin.Context) {
	req := &orderpb.NoParamRequest{}
	res, err := queryServiceClient.GetOrders(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res.DummyRes})

}


// To get all the customers
func GetAllCustomers(c *gin.Context) {
	req := &orderpb.NoParamRequest{}
	res, err := queryServiceClient.GetCustomers(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res.DummyRes})

}

// To get all the restaurants
func GetAllRestaurants(c *gin.Context) {
	req := &orderpb.NoParamRequest{}
	res, err := queryServiceClient.GetRestaurants(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res.DummyRes})

}

// To place a new order
func PostOrder(c *gin.Context) {
	//body := c.Request.Body
	//byteContent, err := ioutil.ReadAll(body)
	//if err != nil {
	//	fmt.Println("Sorry no content found: ", err.Error())
	//}
	//var NewOrder Models.Order
	//_ = json.Unmarshal(byteContent, &NewOrder)
	//var items []*orderProto.OrderStruct_Item
	//for j := range NewOrder.ItemLine {
	//	items = append(items, &orderProto.OrderStruct_Item{Name: NewOrder.ItemLine[j].Name,
	//		Price:    NewOrder.ItemLine[j].Price,
	//		Quantity: NewOrder.ItemLine[j].Quantity})
	//}
	//req := &orderProto.PostRequest{Res: &orderProto.OrderStruct{
	//	OrderID:    NewOrder.OrderID,
	//	CustomerID: NewOrder.CustomerID,
	//	Restaurant: NewOrder.Restaurant,
	//	ItemLine:   items,
	//	Price:      NewOrder.Price,
	//	Quantity:   NewOrder.Quantity,
	//	Discount:   NewOrder.Discount,
	//	Date:       NewOrder.Date,
	//},
	//}
	req := &orderpb.OrderRequest{}

	res, err := queryServiceClient.AddOrder(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": res.DummyRes,
	})
}

// To add new customer
func PostCustomer(c *gin.Context) {

	req := &orderpb.CustomerRequest{}
	res, err := queryServiceClient.AddCustomer(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": res.DummyRes,
	})
}


// To place a new order
func PostRestaurant(c *gin.Context) {
	req := &orderpb.RestaurantRequest{}

	res, err := queryServiceClient.AddRestaurant(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": res.DummyRes,
	})
}


func main() {

	fmt.Println("hello from API INIT function")
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v", err)
	}

	defer conn.Close()

	queryServiceClient = orderpb.NewQueryServiceClient(conn)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiRouter := router.Group("/api")

	apiRouter.GET("/", GetIndex)
	apiRouter.GET("/orders", GetAllOrders)
	apiRouter.GET("/customers", GetAllCustomers)
	apiRouter.GET("/restaurants", GetAllRestaurants)

	apiRouter.POST("/new-order", PostOrder)
	apiRouter.POST("/new-customer", PostCustomer)
	apiRouter.POST("/new-restaurant", PostRestaurant)

	router.Run("localhost:9001")
}
