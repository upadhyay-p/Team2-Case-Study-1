package populate

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/Err"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/Models"
)

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

// populates the DB with sample customer data
func Customers(db *dynamodb.DynamoDB, filename string, tableName string) {
	fmt.Println("Populating Customer table with " + filename)

	// Get table items from customers.csv
	records := readCSV(filename)
	customerList := clubCustomers(records)[:100]

	for _, customer := range customerList {
		customerItem, err := dynamodbattribute.MarshalMap(customer)
		Err.CheckError(err)
		params := &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      customerItem,
		}
		_, err = db.PutItem(params)

		Err.CheckError(err)
	}
	fmt.Println("Done")
}

// populates the DB with sample order data
func Orders(db *dynamodb.DynamoDB, filename string, tableName string) {
	fmt.Println("Populating Orders table with " + filename)

	// Get table items from orders.csv
	records := readCSV(filename)
	// just taking first 100 orders (reason: very slow operation)
	orderList := clubOrders(records)[:100]

	for _, order := range orderList {
		orderItem, err := dynamodbattribute.MarshalMap(order)
		Err.CheckError(err)
		params := &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      orderItem,
		}
		_, err = db.PutItem(params)
		Err.CheckError(err)
	}

	fmt.Println("Done")
}

// populates the DB with sample restaurant data
func Restaurants(db *dynamodb.DynamoDB, filename string, tableName string) {
	fmt.Println("Populating Restaurant table with " + filename)

	// Get table items from restaurants.csv
	records := readCSV(filename)
	restaurantList := clubRestaurants(records)

	for _, restaurant := range restaurantList {
		restaurantItem, err := dynamodbattribute.MarshalMap(restaurant)
		Err.CheckError(err)
		params := &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      restaurantItem,
		}
		_, err = db.PutItem(params)
		Err.CheckError(err)
	}

	fmt.Println("Done")
}

// Parse the record into its parameters
func parseCustomer(record []string) Models.Customer {
	CID, _ := strconv.ParseInt(record[0], 10, 64)
	CName := record[1]
	CAddress := record[2]
	CPhone := record[3]
	customerObj := Models.Customer{ID: CID, Name: CName, Address: CAddress, Phone: CPhone}
	return customerObj
}

// Club the records in slice of order interface
func clubCustomers(records [][]string) []Models.Customer {
	var clubbedRecords []Models.Customer
	var customerObj Models.Customer
	for _, record := range records {
		customerObj = parseCustomer(record)
		clubbedRecords = append(clubbedRecords, customerObj)
	}
	fmt.Printf("Records after clubbing: %v\n", len(clubbedRecords))
	return clubbedRecords
}

// Parse the record into its parameters
func parseOrder(record []string) Models.Order {
	OID, _ := strconv.ParseInt(record[0], 10, 64)
	CID, _ := strconv.ParseInt(record[1], 10, 64)
	RestID, _ := strconv.ParseInt(record[2], 10, 64)
	Discount, _ := strconv.ParseInt(record[3], 10, 64)
	itemName := record[4]
	Cost, _ := strconv.ParseFloat(record[5], 64)
	itemObj := Models.Item{Name: itemName, Price: float32(Cost)}
	orderObj := Models.Order{ID: OID, C_ID: CID, R_ID: RestID, ItemLine: []Models.Item{itemObj}, Price: float32(Cost), Discount: Discount}
	return orderObj
}

// Club the records in slice of order interface
func clubOrders(records [][]string) []Models.Order {
	var clubbedRecords []Models.Order
	prev := "INF"
	var orderObj Models.Order
	var itemObj Models.Item
	flag := false
	for _, record := range records {
		tempObj := parseOrder(record)
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

func parseRestaurant(record []string) Models.Restaurant {
	ID, _ := strconv.ParseInt(record[0], 10, 64)
	Name := record[1]
	Online := true
	if record[2] == "Close" {
		Online = false
	}
	itemName := record[3]
	Price, _ := strconv.ParseFloat(record[4], 64)
	Rating, _ := strconv.ParseFloat(record[5], 64)
	Category := record[6]
	itemObj := Models.Item{Name: itemName, Price: float32(Price)}
	restObj := Models.Restaurant{ID: ID, Name: Name, Online: Online, Menu: []Models.Item{itemObj}, Rating: float32(Rating), Category: Category}
	return restObj
}

func clubRestaurants(records [][]string) []Models.Restaurant {
	var clubbedRecords []Models.Restaurant
	prev := "INF"
	var restObj Models.Restaurant
	var itemObj Models.Item
	flag := false
	for _, record := range records {
		tempObj := parseRestaurant(record)
		if record[0] != prev {
			if flag == true {
				clubbedRecords = append(clubbedRecords, restObj)
			}
			flag = true
			restObj = tempObj
			prev = record[0]
		} else {
			itemObj = tempObj.Menu[0]
			restObj.Menu = append(restObj.Menu, itemObj)
		}
	}
	clubbedRecords = append(clubbedRecords, restObj)
	fmt.Printf("Records after clubbing: %v\n", len(clubbedRecords))
	return clubbedRecords
}
