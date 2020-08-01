package Services

import (
	CustomerModels "Team2CaseStudy1/pkg/Customer/Models"
	"Team2CaseStudy1/pkg/OrderProto/orderpb"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func FetchCustomerTable(db *dynamodb.DynamoDB) []*orderpb.Customer {

	var allCustomers []*orderpb.Customer

	// Create the Expression to fill the input struct with.
	filt := expression.Name("CustomerId").GreaterThan(expression.Value(0))

	proj := expression.NamesList(expression.Name("CustomerId"), expression.Name("Name"), expression.Name("Address"), expression.Name("Phone"))

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
		TableName:                 aws.String("T2-Customer"),
	}

	// Make the DynamoDB Query API call
	result, err := db.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed for Customer table fetched")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	for _, i := range result.Items {
		customerItem := CustomerModels.Customer{}

		err = dynamodbattribute.UnmarshalMap(i, &customerItem)

		if err != nil {
			fmt.Println("Got error unmarshalling customer table")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		allCustomers = append(allCustomers, &orderpb.Customer{CustomerId: customerItem.CustomerId, Name: customerItem.Name, Address: customerItem.Address, Phone: customerItem.Phone})
	}

	return allCustomers

}

func AddCustomerDetails(db *dynamodb.DynamoDB, customer CustomerModels.Customer) {

	customerDynAttr, err := dynamodbattribute.MarshalMap(customer)

	if err != nil {
		panic("Cannot map the values given in Customer struct for post request...")
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String("T2-Customer"),
		Item:      customerDynAttr,
	}

	_, err = db.PutItem(params)

	if err != nil {
		panic("Error in putting the customer item")
	}

}
