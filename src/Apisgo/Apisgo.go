package Apisgo

import (
	"AvgPrice"
	"Structs"
	"TopRestauBuyers"
	"encoding/json"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"order/orderProto"
	"Err"
	"google.golang.org/grpc"
)

var data []Structs.Order
var byteValue []byte
var fname string
var OrderClient orderProto.OrderClient

type server struct{

}

// HomePage for this webserver
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Team 2": "Hello from Aadithya, Abhishek, Priya, Shashi!",
	})
}

// To get all the orders
func GetAllOrders(c *gin.Context) {
	//c.JSON(http.StatusOK, &data)
	for i := range data {
		var items []*orderProto.OrderStruct_Item
		for j := range data[i].ItemLine {
			items = append(items,&orderProto.OrderStruct_Item{Name:data[i].ItemLine[j].Name, Price:data[i].ItemLine[j].Price, Quantity:data[i].ItemLine[j].Quantity})
		}
		req := &orderProto.OrderRequest{OrdReq:&orderProto.OrderStruct{
			OrderID:    data[i].OrderID,
			CustomerID: data[i].CustomerID,
			Restaurant: data[i].Restaurant,
			ItemLine:   items,
			Price:      data[i].Price,
			Quantity:   data[i].Quantity,
			Discount:   data[i].Discount,
			Date:       data[i].Date,
		},
		}

		res, err := OrderClient.CreateOrder(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"orderID": res.OrdRes.OrderID,
			"customerID": res.OrdRes.CustomerID,
			"rest": res.OrdRes.Restaurant,
			"item": res.OrdRes.ItemLine,
			"Price": res.OrdRes.Price,
			"Quantity": res.OrdRes.Quantity,
			"Discount": res.OrdRes.Discount,
			"Date": res.OrdRes.Date,
		})
	}
}


// To get average order price and average number of orders per customer
func GetAvgPricesOrders(c *gin.Context) {
	avgPrices := AvgPrice.INIT(strings.TrimSpace(fname))
	for i := range avgPrices{
		req := &orderProto.AvgPriceInfoRequest{
			CustomerID: avgPrices[i].CustomerID,
			AvgPrice: float32(avgPrices[i].AvgPrice),
			AvgOrders: avgPrices[i].AvgOrders,
		}
		res, err := OrderClient.GetAvgPricesOrders(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"customerID": res.CustomerID,
			"avg .price": res.AvgPrice,
			"avg. orders": res.AvgOrders,
		})
	}
	//c.JSON(http.StatusOK, &avgPrices)
}


// To get "n" top-customers based on expenditure
func GetTopBuyers(c *gin.Context) {
	numberOfBuyers, _ := strconv.ParseInt(c.Param("numBuyers"), 10, 64)
	topCustomersList := TopRestauBuyers.FindTopBuyers(byteValue, numberOfBuyers)
	for i := range topCustomersList{
		req := &orderProto.TopCustomersRequest{
			CustomerID: topCustomersList[i].CustomerID,
			Expenditure: float32(topCustomersList[i].Expenditure),
		}
		res, err := OrderClient.GetTopCustomers(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"CustomerID": res.CustomerID,
			"Expenditure": res.Expenditure,
		})
	}
	//c.JSON(http.StatusOK, &topCustomersList)
}


// To get "n" top-restaurants based on its revenue
func GetTopRestaurants(c *gin.Context) {
	numberOfRestaurants, _ := strconv.ParseInt(c.Param("numRestau"), 10, 64)
	topRestaurantsList := TopRestauBuyers.FindTopRestaurants(byteValue, numberOfRestaurants)

	for i := range topRestaurantsList{
		req := &orderProto.TopRestaurantsRequest{
			Restaurant: topRestaurantsList[i].Restaurant,
			Revenue: float32(topRestaurantsList[i].Revenue),
		}
		res, err := OrderClient.GetTopRest(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Restaurant": res.Restaurant,
			"Revenue": res.Revenue,
		})
	}
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
	data = append(data, NewOrder)
	toJSON()

	c.JSON(http.StatusCreated, gin.H{
		"message": NewOrder,
	})
	fmt.Println("New Entry Added")
}

// To update the json file  and the byteValue slice
func toJSON() {
	byteValue, _ = json.MarshalIndent(data, "", "	  ")
	err := ioutil.WriteFile(fname, byteValue, 0644)
	if err != nil {
		fmt.Println("Error in writing the file")
	}
	fmt.Println("Output file is stored as: " + fname)
}

// for basic authentication
func basicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"team2": "xurde",
	})
}

// for starting the server
func INIT(filename string) {

	fmt.Println("hello from API INIT function")
	fname = filename
	byteValue, _ = ioutil.ReadFile(filename)
	err := json.Unmarshal(byteValue, &data)
	Err.CheckError(err)

	conn ,err := grpc.Dial("localhost:5051",grpc.WithInsecure())

	if err!=nil {
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
