package main

import (
	CustomerDBInit "Team2CaseStudy1/pkg/Customer/InitDB"
	OrderDBInit "Team2CaseStudy1/pkg/Order/InitDB"
	"fmt"

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
	fmt.Println(db)
	CustomerDBInit.CustomerDBInit(db, "../../assets/customer.csv")
	OrderDBInit.OrderDBInit(db, "../../assets/orders.csv")
}
