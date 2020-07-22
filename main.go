package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type order struct {
	OrderID    int64
	CustomerID int64
	Restaurant string
	ItemLine   string
	Price      float64
	Quantity   int64
	Discount   int64
	Date       string
}

func toJSON(orders []order) {
	f, err := json.MarshalIndent(orders, "", "	")
	checkErr(err)
	_ = ioutil.WriteFile("output.json", f, 0644)
}

func main() {
	csvFile, err := os.Open("./orders.csv")
	checkErr(err)
	r := csv.NewReader(csvFile)
	var orderObj order
	_, _ = r.Read()
	var orders []order
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		checkErr(err)
		OID, _ := strconv.ParseInt(record[0], 10, 64)
		CID, _ := strconv.ParseInt(record[1], 10, 64)
		Rest := record[2]
		itemLine := record[3]
		Price, _ := strconv.ParseFloat(record[4], 64)
		Quantity, _ := strconv.ParseInt(record[5], 10, 64)
		Discount, _ := strconv.ParseInt(record[6], 10, 64)
		date := record[7]
		orderObj = order{OID, CID, Rest, itemLine, Price, Quantity, Discount, date}
		orders = append(orders, orderObj)
	}
	fmt.Printf("%v records processed!\n", len(orders))
	toJSON(orders)
}
