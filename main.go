package main

import (
	"AvgPrice"
	"Structs"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

var data []Structs.Order
var AvgPrices []Structs.AvgPriceInfo

const fname = "assets/input.json"
const outputfile = "assets/output"

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Team 2": "Hello from Aadithya, Abhishek, Priya, Shashi!",
	})
}

func GetAllOrders(c *gin.Context) {
	c.JSON(http.StatusOK, &data)
}

func GetAvgPrice(c *gin.Context) {
	c.JSON(http.StatusOK, &AvgPrices)
}

func PostOrder(c *gin.Context) {
	body := c.Request.Body

	content, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Sorry no content found: ", err.Error())
	}
	var NewOrder Structs.Order
	stringConent := string(content)
	_ = json.Unmarshal([]byte(stringConent), &NewOrder)

	c.JSON(http.StatusCreated, gin.H{
		"message": NewOrder,
	})

	// Appends the
	data = append(data, NewOrder)
	toJSON(data, outputfile)
	fmt.Println("New entry added!")

}

func toJSON(orders []Structs.Order, filename string) {
	f, err := json.MarshalIndent(orders, "", "	")
	AvgPrice.CheckError(err)
	err = ioutil.WriteFile(filename+".json", f, 0644)
	AvgPrice.CheckError(err)
	fmt.Println("Output file is stored as: " + filename + ".json")
}

func main() {

	file, _ := ioutil.ReadFile(fname)

	AvgPrices = AvgPrice.INIT(strings.TrimSpace(fname))

	_ = json.Unmarshal([]byte(file), &data)

	router := gin.Default()
	api := router.Group("/api")

	api.GET("/", Index)

	api.GET("/orders", GetAllOrders)

	api.GET("/avg-price", GetAvgPrice)

	api.POST("/new-order", PostOrder)

	router.Run("localhost:9001")

}
