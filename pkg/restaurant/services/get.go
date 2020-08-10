package services

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/Models"
	"github.com/shashijangra22/Team2-Case-Study-1/pkg/restaurant"
)

// db handler to get a particular restaurant
func GetOne(db *dynamodb.DynamoDB, id int64) *restaurant.Restaurant {
	condition := expression.Key("ID").Equal(expression.Value(id))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("Name"), expression.Name("Menu"), expression.Name("Rating"), expression.Name("Category"))
	expr, err := expression.NewBuilder().WithKeyCondition(condition).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression for getting specific Restaurant")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	params := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Team2-RESTAURANTS"),
	}
	res, err := db.Query(params)
	var restaurantDetails []Models.Restaurant
	var rest *restaurant.Restaurant
	_ = dynamodbattribute.UnmarshalListOfMaps(res.Items, &restaurantDetails)

	if len(restaurantDetails) == 1 {
		restItem := restaurantDetails[0]
		var menu []*restaurant.Item
		for _, item := range restItem.Menu {
			menu = append(menu, &restaurant.Item{Name: item.Name, Price: item.Price})
		}
		rest = &restaurant.Restaurant{ID: restItem.ID, Name: restItem.Name, Menu: menu, Online: restItem.Online, Rating: restItem.Rating, Category: restItem.Category}
	}
	return rest
}

// db handler to get all restaurants
func GetAll(db *dynamodb.DynamoDB) []*restaurant.Restaurant {
	var allRestaurants []*restaurant.Restaurant
	// Create the Expression to fill the input struct with.
	filt := expression.Name("ID").GreaterThan(expression.Value(0))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("Name"), expression.Name("Online"), expression.Name("Menu"), expression.Name("Rating"), expression.Name("Category"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression for getting all Restaurants")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Team2-RESTAURANTS"),
	}
	// Make the DynamoDB Query API call
	result, err := db.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed for Restaurant table fetched")
		fmt.Println((err.Error()))
		os.Exit(1)
	}
	for _, i := range result.Items {
		rest := Models.Restaurant{}
		var menu []*restaurant.Item
		err = dynamodbattribute.UnmarshalMap(i, &rest)
		if err != nil {
			fmt.Println("Got error unmarshalling restaurant table")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, item := range rest.Menu {
			menu = append(menu, &restaurant.Item{Name: item.Name, Price: item.Price})
		}
		allRestaurants = append(allRestaurants, &restaurant.Restaurant{ID: rest.ID, Name: rest.Name, Menu: menu, Online: rest.Online, Rating: rest.Rating, Category: rest.Category})
	}
	return allRestaurants
}
