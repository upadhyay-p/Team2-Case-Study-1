package Controller

import (
	"Team2CaseStudy1/pkg/Customer/Models"
	"Team2CaseStudy1/pkg/Err"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// Parse the record into its parameters
func parseRecord(record []string) Models.Customer {
	CID := record[0]
	CName := record[1]
	CAddress := record[2]
	CPhone := record[3]
	customerObj := Models.Customer{CustomerId: CID, Name: CName, Address: CAddress, Phone: CPhone}
	return customerObj
}

// Club the records in slice of order interface
func clubRecords(records [][]string) []Models.Customer {
	var clubbedRecords []Models.Customer
	var customerObj Models.Customer
	for _, record := range records {
		customerObj = parseRecord(record)
		clubbedRecords = append(clubbedRecords, customerObj)
	}
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
func CsvDataForDynamoDB(filename string) []Models.Customer {
	fmt.Println("Reading " + filename)
	records := readCSV(filename)
	customers := clubRecords(records)
	return customers
}
