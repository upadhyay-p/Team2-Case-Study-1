package main

import (
	CustomerController "Team2CaseStudy1/pkg/Customer/Controller"
	OrderController "Team2CaseStudy1/pkg/Order/Controller"

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
	CustomerController.CustomerDBInit(db, "../../assets/customer.csv")
	OrderController.OrderDBInit(db, "../../assets/orders.csv")
}
