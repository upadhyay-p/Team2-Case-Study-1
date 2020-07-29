package Apisgo

import (
	"encoding/json"
	//"encoding/json"
	//"AvgPrice"
	"Structs"
	//"TopRestauBuyers"
	//"encoding/json"
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"strconv"
	//"strings"

	//"io"
	//"io/ioutil"
	"log"
	"net/http"
	//"strconv"
	//"strings"

	"github.com/gin-gonic/gin"
	"order/orderProto"
	//"Err"
	"google.golang.org/grpc"
)

var OrderClient orderProto.OrderClient

// HomePage for this webserver
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Team 2": "Hello from Aadithya, Abhishek, Priya, Shashi!",
	})
}

// To get all the orders
func GetAllOrders(c *gin.Context) {
	//c.JSON(http.StatusOK, &data)
	req := &orderProto.OrderRequest{}
	res, err := OrderClient.CreateOrder(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res.OrdRes})

}

//To get average order price and average number of orders per customer
func GetAvgPricesOrders(c *gin.Context) {
	req := &orderProto.AvgPriceInfoRequest{}

	res, err := OrderClient.GetAvgPricesOrders(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": res.Res,
	})

}

// To get "n" top-customers based on expenditure
func GetTopBuyers(c *gin.Context) {
	numberOfBuyers, _ := strconv.ParseInt(c.Param("numBuyers"), 10, 64)
	req := &orderProto.TopCustomersRequest{
		Num: numberOfBuyers,
	}
	res, err := OrderClient.GetTopCustomers(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": res.Res,
	})
	//c.JSON(http.StatusOK, &topCustomersList)
}

// To get "n" top-restaurants based on its revenue
func GetTopRestaurants(c *gin.Context) {
	numberOfRestaurants, _ := strconv.ParseInt(c.Param("numRestau"), 10, 64)
	req := &orderProto.TopRestaurantsRequest{
		Num: numberOfRestaurants,
	}
	res, err := OrderClient.GetTopRest(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": res.Res,
	})
	//c.JSON(http.StatusOK, &topRestaurantsList)
}

// To place a new order
func PostOrder(c *gin.Context) {
	body := c.Request.Body
	byteContent, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Sorry no content found: ", err.Error())
	}
	var NewOrder Structs.Order
	_ = json.Unmarshal(byteContent, &NewOrder)
	var items []*orderProto.OrderStruct_Item
	for j := range NewOrder.ItemLine {
		items = append(items, &orderProto.OrderStruct_Item{Name: NewOrder.ItemLine[j].Name,
			Price:    NewOrder.ItemLine[j].Price,
			Quantity: NewOrder.ItemLine[j].Quantity})
	}
	req := &orderProto.PostRequest{Res: &orderProto.OrderStruct{
		OrderID:    NewOrder.OrderID,
		CustomerID: NewOrder.CustomerID,
		Restaurant: NewOrder.Restaurant,
		ItemLine:   items,
		Price:      NewOrder.Price,
		Quantity:   NewOrder.Quantity,
		Discount:   NewOrder.Discount,
		Date:       NewOrder.Date,
	},
	}


	res, err := OrderClient.PostOrder(c, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": res.Res,
	})
}

//// for basic authentication
func basicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"team2": "xurde",
	})
}

// for starting the server
func INIT(filename string) {

	fmt.Println("hello from API INIT function")
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v", err)
	}

	defer conn.Close()

	OrderClient = orderProto.NewOrderClient(conn)

	// router := gin.Default()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiRouter := router.Group("/api")
	authRouter := router.Group("/auth")
	authRouter.Use(basicAuth())

	apiRouter.GET("/", GetIndex)

	apiRouter.GET("/orders", GetAllOrders)

	apiRouter.GET("/avg-price", GetAvgPricesOrders)

	authRouter.GET("/top-buyers/:numBuyers", GetTopBuyers)

	authRouter.GET("/top-restaurants/:numRestau", GetTopRestaurants)

	authRouter.POST("/new-order", PostOrder)

	router.Run("localhost:9001")
}
