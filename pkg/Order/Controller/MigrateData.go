package Controller

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func MigrateCustomerData(db *dynamodb.DynamoDB, filename string) {

	// Get table items from customer.csv
	orderList := CsvDataForDynamoDB(filename)

	for _, order := range orderList {

		orderItem, err := dynamodbattribute.MarshalMap(order)

		if err != nil {
			panic("Got error marshalling Customer map: dynamodb")
		}
		// fmt.Println("Customer is: ", customer)

		params := &dynamodb.PutItemInput{
			TableName: aws.String("T2-Order"),
			Item:      orderItem,
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
	orderID := "2"
	customerID := "278"

	params := &dynamodb.GetItemInput{
		TableName: aws.String("T2-Order"),
		Key: map[string]*dynamodb.AttributeValue{
			"OrderId": {
				S: aws.String(orderID),
			},
			"CustomerId": {
				S: aws.String(customerID),
			},
		},
	}
	resp, err := db.GetItem(params)
	if err != nil {
		fmt.Println("Sorry item not found...")
	}
	fmt.Println(resp)
}
