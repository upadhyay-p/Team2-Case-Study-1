package Apisgo

import (
	"AvgPrice"
	"Toprestaubuyers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"structs"

	"github.com/gin-gonic/gin"
)

var data []structs.Order
var byteValue []byte
var fname string

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Team 2": "Hello from Aadithya, Abhishek, Priya, Shashi!",
	})
}

func GetAllOrders(c *gin.Context) {
	c.JSON(http.StatusOK, &data)
}

func GetAvgPrice(c *gin.Context) {
	avgPrices := AvgPrice.INIT(strings.TrimSpace(fname))
	c.JSON(http.StatusOK, &avgPrices)
}

func GetTopBuyers(c *gin.Context) {
	numberOfBuyers, _ := strconv.ParseInt(c.Param("numBuyers"), 10, 64)
	topCustomersList := Toprestaubuyers.FindTopBuyers(byteValue, numberOfBuyers)
	c.JSON(http.StatusOK, &topCustomersList)
}

func GetTopRestaurants(c *gin.Context) {
	numberOfRestaurants, _ := strconv.ParseInt(c.Param("numRestau"), 10, 64)
	topRestaurantsList := Toprestaubuyers.FindTopRestaurants(byteValue, numberOfRestaurants)
	c.JSON(http.StatusOK, &topRestaurantsList)
}

func PostOrder(c *gin.Context) {
	body := c.Request.Body
	byteContent, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Sorry no content found: ", err.Error())
	}

	var NewOrder structs.Order
	_ = json.Unmarshal(byteContent, &NewOrder)
	data = append(data, NewOrder)
	toJSON()

	c.JSON(http.StatusCreated, gin.H{
		"message": NewOrder,
	})
	fmt.Println("New Entry Added")
}

func toJSON() {
	byteValue, _ = json.MarshalIndent(data, "", "	  ")
	err := ioutil.WriteFile(fname, byteValue, 0644)
	if err != nil {
		fmt.Println("Error in writing the file")
	}
	fmt.Println("Output file is stored as: " + fname)
}

func INIT(filename string) {

	fmt.Println("hello from API INIT function")
	fname = filename
	byteValue, _ = ioutil.ReadFile(filename)
	_ = json.Unmarshal(byteValue, &data)

	router := gin.Default()
	apiRouter := router.Group("/api")

	apiRouter.GET("/", GetIndex)

	apiRouter.GET("/orders", GetAllOrders)

	apiRouter.GET("/avg-price", GetAvgPrice)

	apiRouter.GET("/top-buyers/:numBuyers", GetTopBuyers)

	apiRouter.GET("/top-restaurants/:numRestau", GetTopRestaurants)

	apiRouter.POST("/new-order", PostOrder)

	router.Run("localhost:9001")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
