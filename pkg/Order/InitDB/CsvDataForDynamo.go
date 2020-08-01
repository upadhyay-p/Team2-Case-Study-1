package DBInit

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
	OID, _ := strconv.ParseInt(record[0], 10, 64)
	CID, _ := strconv.ParseInt(record[1], 10, 64)
	RestID, _ := strconv.ParseInt(record[2], 10, 64)
	Discount, _ := strconv.ParseInt(record[3], 10, 64)
	itemName := record[4]
	Cost, _ := strconv.ParseFloat(record[5], 64)
	itemObj := Models.Item{Name: itemName, Price: float32(Cost)}
	orderObj := Models.Order{OrderId: OID, CustomerId: CID, RestaurantId: RestID, ItemLine: []Models.Item{itemObj}, Price: float32(Cost), Discount: Discount}
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
			orderObj.Price += itemObj.Price
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
