package AvgPrice

import (
	"fmt"
	"io/ioutil"
	"structs"

	"github.com/tidwall/gjson"
)

func AvgPriceReport(filename string) []structs.AvgPriceInfo {

	file, err := ioutil.ReadFile(filename)
	CheckError(err)

	ItemPrices := make([]float64, 0, 5)
	price := gjson.GetBytes(file, "#.Price").Array()

	for _, val := range price {
		ItemPrices = append(ItemPrices, val.Float())
	}

	cids := gjson.GetBytes(file, "#.CustomerID").Array()
	custOrders := make(map[int64]int64)
	custSpend := make(map[int64]float64)

	for i, cid := range cids {
		_, f := custOrders[cid.Int()]
		if f == false {
			custOrders[cid.Int()] = 0
			custSpend[cid.Int()] = 0
		}
		custOrders[cid.Int()] += 1
		custSpend[cid.Int()] += ItemPrices[i]
	}
	var n int64
	avgPrices := make([]structs.AvgPriceInfo, 0, 50)
	for cid, cost := range custSpend {
		n = custOrders[cid]
		avgP := cost / float64(n)
		obj := structs.AvgPriceInfo{CustomerID: cid, AvgPrice: avgP, AvgOrders: n}
		avgPrices = append(avgPrices, obj)
		//fmt.Printf("Customer ID: %v No of Orders: %v Average Price: %.2f \n", cid, n, cost/float64(n))
	}
	return avgPrices
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func INIT(filename string) []structs.AvgPriceInfo {
	//	"./orders.json"
	fmt.Println("Reading " + filename)
	res := AvgPriceReport(filename)
	return res
}
