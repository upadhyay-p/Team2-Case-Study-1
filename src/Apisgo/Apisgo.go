package Apisgo

import (
	"AvgPrice"
	"Structs"
	"TopRestauBuyers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var data []Structs.Order
var byteValue []byte
var fname string

// HomePage for this webserver
func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Team 2": "Hello from Aadithya, Abhishek, Priya, Shashi!",
	})
}

// To get all the orders
func GetAllOrders(c *gin.Context) {
	c.JSON(http.StatusOK, &data)
}

// To get average order price and average number of orders per customer
func GetAvgPrice(c *gin.Context) {
	avgPrices := AvgPrice.INIT(strings.TrimSpace(fname))
	c.JSON(http.StatusOK, &avgPrices)
}

// To get "n" top-customers based on expenditure
func GetTopBuyers(c *gin.Context) {
	numberOfBuyers, _ := strconv.ParseInt(c.Param("numBuyers"), 10, 64)
	topCustomersList := TopRestauBuyers.FindTopBuyers(byteValue, numberOfBuyers)
	c.JSON(http.StatusOK, &topCustomersList)
}

// To get "n" top-restaurants based on its revenue
func GetTopRestaurants(c *gin.Context) {
	numberOfRestaurants, _ := strconv.ParseInt(c.Param("numRestau"), 10, 64)
	topRestaurantsList := TopRestauBuyers.FindTopRestaurants(byteValue, numberOfRestaurants)
	c.JSON(http.StatusOK, &topRestaurantsList)
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
	_ = json.Unmarshal(byteValue, &data)

	// router := gin.Default()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	apiRouter := router.Group("/api")
	authRouter := router.Group("/auth")
	authRouter.Use(basicAuth())

	apiRouter.GET("/", GetIndex)

	apiRouter.GET("/orders", GetAllOrders)

	apiRouter.GET("/avg-price", GetAvgPrice)

	authRouter.GET("/top-buyers/:numBuyers", GetTopBuyers)

	authRouter.GET("/top-restaurants/:numRestau", GetTopRestaurants)

	authRouter.POST("/new-order", PostOrder)

	router.Run("localhost:9001")
}
