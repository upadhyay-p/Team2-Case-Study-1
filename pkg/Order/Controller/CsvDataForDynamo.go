package Controller

import (
	"Team2CaseStudy1/pkg/Err"
	"Team2CaseStudy1/pkg/Order/Models"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Parse the record into its parameters
func parseRecord(record []string) Models.Order {
	OID := record[0]
	CID := record[1]
	RestID := record[2]
	Discount := record[3]
	itemName := record[4]
	Cost := record[5]
	itemObj := Models.Item{Name: itemName, Price: Cost}
	orderObj := Models.Order{OrderId: OID, CustomerId: CID, RestaurantId: RestID, ItemLine: []Models.Item{itemObj}, Price: Cost, Discount: Discount}
	return orderObj
}

// Club the records in slice of order interface
func clubRecords(records [][]string) []Models.Order {
	var clubbedRecords []Models.Order
	prev := "INF"
	var orderObj Models.Order
	var itemObj Models.Item
	flag := false
	for _, record := range records {
		tempObj := parseRecord(record)
		if record[0] != prev {
			if flag == true {
				clubbedRecords = append(clubbedRecords, orderObj)
			}
			flag = true
			orderObj = tempObj
			prev = record[0]
		} else {
			itemObj = tempObj.ItemLine[0]
			orderObj.ItemLine = append(orderObj.ItemLine, itemObj)
			orderObjPrice, _ := strconv.ParseFloat(orderObj.Price, 64)
			itemObjPrice, _ := strconv.ParseFloat(itemObj.Price, 64)
			orderObjPrice += itemObjPrice
			orderObj.Price = fmt.Sprintf("%.2f", orderObjPrice)

		}
	}
	clubbedRecords = append(clubbedRecords, orderObj)
	fmt.Printf("Records after clubbing: %v\n", len(clubbedRecords))
	return clubbedRecords
}

// Read the csv file and store as a string in records variable
func readCSV(filename string) [][]string {
	csvFile, err := os.Open(filename)
	Err.CheckError(err)
	r := csv.NewReader(csvFile)
	_, _ = r.Read()
	var records [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		Err.CheckError(err)
		records = append(records, record)
	}
	fmt.Printf("Records processed: %v\n", len(records))
	return records
}

// Initialise to convert the csv file to json format (json object)
func CsvDataForDynamoDB(filename string) []Models.Order {
	fmt.Println("Reading " + filename)
	records := readCSV(filename)
	orders := clubRecords(records)
	return orders
}
