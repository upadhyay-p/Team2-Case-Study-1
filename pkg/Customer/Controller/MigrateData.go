package Controller

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func MigrateCustomerData(db *dynamodb.DynamoDB, filename string) {

	// Get table items from customer.csv
	customerList := CsvDataForDynamoDB(filename)

	for _, customer := range customerList {

		customerItem, err := dynamodbattribute.MarshalMap(customer)

		if err != nil {
			panic("Got error marshalling Customer map: dynamodb")
		}
		// fmt.Println("Customer is: ", customer)

		params := &dynamodb.PutItemInput{
			TableName: aws.String("T1-Customer"),
			Item:      customerItem,
		}

		// fmt.Println("CustomerParams is: ", params)
		_, err = db.PutItem(params)

		if err != nil {
			fmt.Println("Error is: ", err)
			panic("Got error calling PutItem in Customer DB Migration")
		}

	}

	// to check the data
	// query parameters
	customerID := "6"
	name := "Simonette"

	params := &dynamodb.GetItemInput{
		TableName: aws.String("T1-Customer"),
		Key: map[string]*dynamodb.AttributeValue{
			"CustomerId": {
				S: aws.String(customerID),
			},
			"Name": {
				S: aws.String(name),
			},
		},
	}
	resp, err := db.GetItem(params)
	if err != nil {
		fmt.Println("Sorry item not found...")
	}
	fmt.Println(resp)
}
