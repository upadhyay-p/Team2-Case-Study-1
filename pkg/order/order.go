package order

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/Models"
)

var OSC OrderServiceClient

// client call to get all the orders
func GetAll(c *gin.Context) {
	req := &NoParamRequest{}
	res, err := OSC.GetOrders(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"_": res.Orders})
}

// client call to get a particular order
func GetOne(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	req := &IDRequest{ID: id}
	res, err := OSC.GetOrder(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"_": res})

}

// client call to add new order
func Add(c *gin.Context) {
	byteContent, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Sorry no content found: ", err.Error())
	}
	var ord Models.Order
	_ = json.Unmarshal(byteContent, &ord)
	var itemline []*Item
	for i := range ord.ItemLine {
		itemline = append(itemline, &Item{
			Name:  ord.ItemLine[i].Name,
			Price: ord.ItemLine[i].Price,
		})
	}
	req := &Order{ID: ord.ID, C_ID: ord.C_ID, R_ID: ord.R_ID, ItemLine: itemline, Price: ord.Price, Discount: ord.Discount}
	res, err := OSC.AddOrder(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"_": res})
}
