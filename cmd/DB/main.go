package main

import (
	"Team2CaseStudy1/pkg/Customer/Controller"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:8000"),
		Region:   aws.String("us-east-1"),
	}))

	db := dynamodb.New(sess)
	Controller.CustomerDBInit(db, "../../assets/customer.csv")

}
