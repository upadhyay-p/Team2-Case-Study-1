package Controller

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CustomerDBInit(db *dynamodb.DynamoDB, filename string) {

	awsParams := &dynamodb.CreateTableInput{
		TableName: aws.String("T2-Customer"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("CustomerId"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("Name"), KeyType: aws.String("RANGE")},
		},

		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("CustomerId"), AttributeType: aws.String("S")},
			{AttributeName: aws.String("Name"), AttributeType: aws.String("S")},
		},

		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}
	resp, err := db.CreateTable(awsParams)
	if err != nil {
		fmt.Println("Sorry, error creating table : ", err)
	} else {
		fmt.Println(resp)
		MigrateCustomerData(db, filename)
	}
}
