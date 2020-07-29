package AvgPrice

import (
	"Team2CaseStudy1/pkg/Err"
	"Team2CaseStudy1/pkg/Models"
	"fmt"
	"io/ioutil"

	"github.com/tidwall/gjson"
)

func AvgPriceReport(filename string) []Models.AvgPriceInfo {

	file, err := ioutil.ReadFile(filename)
	Err.CheckError(err)

	ItemPrices := make([]float32, 0, 5)
	price := gjson.GetBytes(file, "#.Price").Array()

	for _, val := range price {
		ItemPrices = append(ItemPrices, float32(val.Float()))
	}

	cids := gjson.GetBytes(file, "#.CustomerID").Array()
	custOrders := make(map[int64]int64)
	custSpend := make(map[int64]float32)

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
	avgPrices := make([]Models.AvgPriceInfo, 0, 50)
	for cid, cost := range custSpend {
		n = custOrders[cid]
		avgP := cost / float32(n)
		obj := Models.AvgPriceInfo{CustomerID: cid, AvgPrice: avgP, AvgOrders: n}
		avgPrices = append(avgPrices, obj)
		//fmt.Printf("Customer ID: %v No of Orders: %v Average Price: %.2f \n", cid, n, cost/float32(n))
	}
	return avgPrices
}

func INIT(filename string) []Models.AvgPriceInfo {
	//	"./orders.json"
	fmt.Println("Reading " + filename)
	res := AvgPriceReport(filename)
	return res
}
