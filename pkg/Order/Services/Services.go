package Services

import (
	OrderModels "Team2CaseStudy1/pkg/Order/Models"
	"Team2CaseStudy1/pkg/OrderProto/orderpb"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func FetchOrderTable(db *dynamodb.DynamoDB) []*orderpb.Order {

	var allOrders []*orderpb.Order

	// Create the Expression to fill the input struct with.
	filt := expression.Name("OrderId").GreaterThan(expression.Value(0))

	proj := expression.NamesList(expression.Name("OrderId"), expression.Name("CustomerId"), expression.Name("RestaurantId"), expression.Name("ItemLine"), expression.Name("Price"), expression.Name("Discount"))

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
		TableName:                 aws.String("T2-Order"),
	}

	// Make the DynamoDB Query API call
	result, err := db.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed for Order table fetched")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	for _, i := range result.Items {
		orderItem := OrderModels.Order{}
		var itemLine []*orderpb.Item

		err = dynamodbattribute.UnmarshalMap(i, &orderItem)

		if err != nil {
			fmt.Println("Got error unmarshalling order table")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		for _, item := range orderItem.ItemLine {
			itemLine = append(itemLine, &orderpb.Item{Name: item.Name, Price: item.Price})
		}

		allOrders = append(allOrders, &orderpb.Order{OrderId: orderItem.OrderId, CustomerId: orderItem.CustomerId, RestaurantId: orderItem.RestaurantId, ItemLine: itemLine, Price: orderItem.Price, Discount: orderItem.Discount})
	}

	return allOrders

}

func AddOrderDetails(db *dynamodb.DynamoDB, order OrderModels.Order) {

	orderDynAttr, err := dynamodbattribute.MarshalMap(order)

	if err != nil {
		panic("Cannot map the values given in Order struct for post request...")
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String("T2-Order"),
		Item:      orderDynAttr,
	}

	_, err = db.PutItem(params)

	if err != nil {
		panic("Error in putting the order item")
	}

}
