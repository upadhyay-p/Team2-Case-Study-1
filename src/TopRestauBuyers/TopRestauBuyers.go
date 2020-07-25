package Toprestaubuyers

import (
	"fmt"
	"math"
	"sort"
	"structs"

	"github.com/tidwall/gjson"
)

func FindTopRestaurants(byteValue []byte, numRestau int64) []structs.TopRestaurants {
	resps := gjson.GetManyBytes(byteValue, "#.Restaurant", "#.Price")
	restaurantsRevenue := make(map[string]float64)
	var restaurants []string
	var revenues []float64
	for i, resp := range resps {
		if i == 0 {
			for _, rest := range resp.Array() {
				restaurants = append(restaurants, rest.String())
			}
		} else {
			for _, rev := range resp.Array() {
				revenues = append(revenues, rev.Float())
			}
		}
	}
	for ind := 0; ind < len(restaurants); ind++ {
		restaurantsRevenue[restaurants[ind]] += revenues[ind]
	}
	//fmt.Println(len(restaurantsRevenue))

	var topRestaurants []structs.TopRestaurants
	for rests, revs := range restaurantsRevenue {
		topRestaurants = append(topRestaurants, structs.TopRestaurants{rests, revs})
	}
	sort.Slice(topRestaurants, func(i, j int) bool {
		return topRestaurants[i].Revenue > topRestaurants[j].Revenue
	})
	var topRestaurantsList []structs.TopRestaurants
	fmt.Println("The top-5 Restaurants having following revenues are:")
	for ind := 0; ind < int(math.Min(float64(numRestau), float64(len(topRestaurants)))); ind++ {
		fmt.Println(topRestaurants[ind])
		topRestaurantsList = append(topRestaurantsList, topRestaurants[ind])
	}
	return topRestaurantsList
}

func FindTopBuyers(byteValue []byte, numCust int64) []structs.TopCustomers {
	resps := gjson.GetManyBytes(byteValue, "#.CustomerID", "#.Price")
	customersExpenditure := make(map[string]float64)
	var customers []string
	var expenditure []float64
	for i, resp := range resps {
		if i == 0 {
			for _, rest := range resp.Array() {
				customers = append(customers, rest.String())
			}
		} else {
			for _, rev := range resp.Array() {
				expenditure = append(expenditure, rev.Float())
			}
		}
	}
	for ind := 0; ind < len(customers); ind++ {
		customersExpenditure[customers[ind]] += expenditure[ind]
	}
	//fmt.Println(len(customersExpenditure))

	var topCustomers []structs.TopCustomers
	for custs, expd := range customersExpenditure {
		topCustomers = append(topCustomers, structs.TopCustomers{custs, expd})
	}
	sort.Slice(topCustomers, func(i, j int) bool {
		return topCustomers[i].Expenditure > topCustomers[j].Expenditure
	})
	var topCustomersList []structs.TopCustomers
	fmt.Println("The top-5 Buyers having following expenditures are:")
	for ind := 0; ind < int(math.Min(float64(numCust), float64(len(topCustomers)))); ind++ {
		fmt.Println(topCustomers[ind])
		topCustomersList = append(topCustomersList, topCustomers[ind])
	}
	return topCustomersList
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
