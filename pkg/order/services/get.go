package services

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/Models"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/order"
)

// database handler to get a particular order
func GetOne(db *dynamodb.DynamoDB, id int64) *order.Order {
	condition := expression.Key("ID").Equal(expression.Value(id))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("C_ID"), expression.Name("R_ID"), expression.Name("ItemLine"), expression.Name("Price"), expression.Name("Discount"))
	expr, err := expression.NewBuilder().WithKeyCondition(condition).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression for getting specific Order")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	params := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Team2-ORDERS"),
	}
	res, err := db.Query(params)
	var orderDetails []Models.Order
	var ord *order.Order
	_ = dynamodbattribute.UnmarshalListOfMaps(res.Items, &orderDetails)

	if len(orderDetails) == 1 {
		orderItem := orderDetails[0]
		var itemline []*order.Item
		for _, item := range orderItem.ItemLine {
			itemline = append(itemline, &order.Item{Name: item.Name, Price: item.Price})
		}
		ord = &order.Order{ID: orderItem.ID, C_ID: orderItem.C_ID, R_ID: orderItem.R_ID, ItemLine: itemline, Price: orderItem.Price, Discount: orderItem.Discount}
	}
	return ord
}

// database handler to get all orders
func GetAll(db *dynamodb.DynamoDB) []*order.Order {
	var allOrders []*order.Order
	// Create the Expression to fill the input struct with.
	filt := expression.Name("ID").GreaterThan(expression.Value(0))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("C_ID"), expression.Name("R_ID"), expression.Name("ItemLine"), expression.Name("Price"), expression.Name("Discount"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression for getting all Orders")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Team2-ORDERS"),
	}
	// Make the DynamoDB Query API call
	result, err := db.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed for Order table fetched")
		fmt.Println((err.Error()))
		os.Exit(1)
	}
	for _, i := range result.Items {
		orderItem := Models.Order{}
		var itemLine []*order.Item
		err = dynamodbattribute.UnmarshalMap(i, &orderItem)
		if err != nil {
			fmt.Println("Got error unmarshalling order table")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, item := range orderItem.ItemLine {
			itemLine = append(itemLine, &order.Item{Name: item.Name, Price: item.Price})
		}
		allOrders = append(allOrders, &order.Order{ID: orderItem.ID, C_ID: orderItem.C_ID, R_ID: orderItem.R_ID, ItemLine: itemLine, Price: orderItem.Price, Discount: orderItem.Discount})
	}
	return allOrders
}
