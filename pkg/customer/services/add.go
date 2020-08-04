package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/Models"
)

func AddOne(db *dynamodb.DynamoDB, cst Models.Customer) {
	cstItem, err := dynamodbattribute.MarshalMap(cst)
	if err != nil {
		panic("Cannot map the values given in Customer struct for post request...")
	}
	params := &dynamodb.PutItemInput{
		TableName: aws.String("Team2-CUSTOMERS"),
		Item:      cstItem,
	}
	_, err = db.PutItem(params)
	if err != nil {
		panic("Error in putting the customer item")
	}
}
