package customer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shashijangra22/bootcamp-project/pkg/Models"
)

var CSC CustomerServiceClient

// client call to get all the customers
func GetAll(c *gin.Context) {
	req := &NoParamRequest{}
	res, err := CSC.GetCustomers(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"_": res.Customers})
}

// client call to get a particular customer
func GetOne(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	req := &IDRequest{ID: id}
	res, err := CSC.GetCustomer(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"_": res})
}

// client call to add new customer
func Add(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("No Data recieved", err.Error())
	}
	var cst Models.Customer
	_ = json.Unmarshal(bytes, &cst)
	req := &Customer{ID: cst.ID, Name: cst.Name, Address: cst.Address, Phone: cst.Phone}
	res, err := CSC.AddCustomer(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"_": res})
}
