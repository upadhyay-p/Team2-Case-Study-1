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
	ItemLine   []string
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

func parseRecord(record []string) order {
	OID, _ := strconv.ParseInt(record[0], 10, 64)
	CID, _ := strconv.ParseInt(record[1], 10, 64)
	Rest := record[2]
	item := record[3]
	Price, _ := strconv.ParseFloat(record[4], 64)
	Quantity, _ := strconv.ParseInt(record[5], 10, 64)
	Discount, _ := strconv.ParseInt(record[6], 10, 64)
	date := record[7]
	orderObj := order{OID, CID, Rest, []string{item}, Price, Quantity, Discount, date}
	return orderObj
}

func clubRecords(records [][]string) []order {
	var clubbedRecords []order
	prev := "INF"
	var orderObj order
	flag := false
	for _, record := range records {
		if record[0] != prev {
			if flag == true {
				clubbedRecords = append(clubbedRecords, orderObj)
			}
			flag = true
			orderObj = parseRecord(record)
			prev = record[0]
		} else {
			orderObj.ItemLine = append(orderObj.ItemLine, record[3])
		}
	}
	return clubbedRecords
}

func main() {
	csvFile, err := os.Open("./orders2.csv")
	checkErr(err)
	r := csv.NewReader(csvFile)
	_, _ = r.Read()
	var records [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		checkErr(err)
		records = append(records, record)
	}
	fmt.Printf("%v records processed!\n", len(records))
	orders := clubRecords(records)
	fmt.Printf("Records after clubbing: %v\n", len(orders))
	toJSON(orders)
}
