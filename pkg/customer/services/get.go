package services

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/Models"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/customer"
)

// database handler to get a particular customer
func GetOne(db *dynamodb.DynamoDB, id int64) *customer.Customer {
	condition := expression.Key("ID").Equal(expression.Value(id))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("Name"), expression.Name("Address"), expression.Name("Phone"))
	expr, err := expression.NewBuilder().WithKeyCondition(condition).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression for getting specific Customer")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	params := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Team2-CUSTOMERS"),
	}
	res, err := db.Query(params)
	var customerDetails []Models.Customer
	_ = dynamodbattribute.UnmarshalListOfMaps(res.Items, &customerDetails)
	var cst *customer.Customer
	if len(customerDetails) == 1 {
		customerItem := customerDetails[0]
		cst = &customer.Customer{ID: customerItem.ID, Name: customerItem.Name, Address: customerItem.Address, Phone: customerItem.Phone}
	}
	return cst
}

// database handler to get all customers
func GetAll(db *dynamodb.DynamoDB) []*customer.Customer {
	var allCustomers []*customer.Customer
	// Create the Expression to fill the input struct with.
	filt := expression.Name("ID").GreaterThan(expression.Value(0))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("Name"), expression.Name("Address"), expression.Name("Phone"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression for getting all Customers")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Team2-CUSTOMERS"),
	}
	// Make the DynamoDB Query API call
	result, err := db.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed for Customer table fetched")
		fmt.Println((err.Error()))
		os.Exit(1)
	}
	for _, i := range result.Items {
		cst := Models.Customer{}
		err = dynamodbattribute.UnmarshalMap(i, &cst)
		if err != nil {
			fmt.Println("Got error unmarshalling customer table")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		allCustomers = append(allCustomers, &customer.Customer{ID: cst.ID, Name: cst.Name, Address: cst.Address, Phone: cst.Phone})
	}
	return allCustomers
}
