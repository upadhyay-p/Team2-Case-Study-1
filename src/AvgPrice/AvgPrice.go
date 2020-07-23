package AvgPrice

import (
	"../CSV2JSON"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
)


func AvgPriceReport(filename string) {

	file, err := ioutil.ReadFile(filename)
	CSV2JSON.CheckErr(err)

	ItemPrices:= make([]float64, 0, 5)
	price := gjson.GetBytes(file,"#.Price").Array()

	for _, val := range price {
		ItemPrices = append(ItemPrices, val.Float())
	}

	cids := gjson.GetBytes(file,"#.CustomerID").Array()
	custOrders := make(map[int64] int64)
	custSpend := make(map[int64] float64)

	for i,cid := range cids {
		_, f := custOrders[cid.Int()]
		if f == false {
			custOrders[cid.Int()] = 0
			custSpend[cid.Int()] = 0
		}
		custOrders[cid.Int()] += 1
		custSpend[cid.Int()] += ItemPrices[i]
	}
	var n int64
	for cid, cost := range custSpend {
		n = custOrders[cid]
		fmt.Printf("Customer ID: %v No of Orders: %v Average Price: %.2f \n", cid, n, cost/float64(n))
	}
}

func INIT(filename string) {
//	"./orders.json"
	fmt.Println("Reading " + filename)
	AvgPriceReport(filename)
}