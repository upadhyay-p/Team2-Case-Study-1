package DBInit

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func OrderDBInit(db *dynamodb.DynamoDB, filename string) {

	awsParams := &dynamodb.CreateTableInput{
		TableName: aws.String("T2-Order"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("OrderId"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("CustomerId"), KeyType: aws.String("RANGE")},
		},

		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("OrderId"), AttributeType: aws.String("N")},
			{AttributeName: aws.String("CustomerId"), AttributeType: aws.String("N")},
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
		MigrateOrderData(db, filename)
	}
}
