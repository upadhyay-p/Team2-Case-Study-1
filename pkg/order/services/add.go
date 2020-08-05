package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/Models"
)

// database handler to add an order
func Add(db *dynamodb.DynamoDB, ord Models.Order) {
	orderDynAttr, err := dynamodbattribute.MarshalMap(ord)
	if err != nil {
		panic("Cannot map the values given in Order struct for post request...")
	}
	params := &dynamodb.PutItemInput{
		TableName: aws.String("Team2-ORDERS"),
		Item:      orderDynAttr,
	}
	_, err = db.PutItem(params)
	if err != nil {
		panic("Error in putting the order item")
	}
}
