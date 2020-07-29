package TopRestauBuyers

import (
	"Team2CaseStudy1/pkg/Err"
	"Team2CaseStudy1/pkg/Models"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"

	"github.com/tidwall/gjson"
)

func FindTopRestaurants(byteValue []byte, numRestau int64) []Models.TopRestaurants {
	TopNRestaurantsList := make([]Models.TopRestaurants, 0)
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

	var TopRestaurantsList []Models.TopRestaurants
	for rests, revs := range restaurantsRevenue {
		TopRestaurantsList = append(TopRestaurantsList, Models.TopRestaurants{Restaurant: rests, Revenue: float32(revs)})
	}
	sort.Slice(TopRestaurantsList, func(i, j int) bool {
		return TopRestaurantsList[i].Revenue > TopRestaurantsList[j].Revenue
	})
	fmt.Println("The top-5 Restaurants having following revenues are:")
	for ind := 0; ind < int(math.Min(float64(numRestau), float64(len(TopRestaurantsList)))); ind++ {
		fmt.Println(TopRestaurantsList[ind])
		TopNRestaurantsList = append(TopNRestaurantsList, TopRestaurantsList[ind])
	}

	return TopNRestaurantsList
}

func FindTopBuyers(byteValue []byte, numCust int64) []Models.TopCustomers {
	TopNCustomersList := make([]Models.TopCustomers, 0)
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

	var TopCustomersList []Models.TopCustomers
	for custs, expd := range customersExpenditure {
		TopCustomersList = append(TopCustomersList, Models.TopCustomers{CustomerID: custs, Expenditure: float32(expd)})
	}
	sort.Slice(TopCustomersList, func(i, j int) bool {
		return TopCustomersList[i].Expenditure > TopCustomersList[j].Expenditure
	})
	fmt.Println("The top-5 Buyers having following expenditures are:")
	for ind := 0; ind < int(math.Min(float64(numCust), float64(len(TopCustomersList)))); ind++ {
		fmt.Println(TopCustomersList[ind])
		TopNCustomersList = append(TopNCustomersList, TopCustomersList[ind])
	}
	return TopNCustomersList
}

func RestauBuyers(filename string) {
	jsonFile, err := os.Open(filename)
	Err.CheckError(err)
	fmt.Println("Successfully Opened orders.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	FindTopRestaurants(byteValue, 5) // number of top restaurants you want
	FindTopBuyers(byteValue, 5)      // number of top customer who buys most
}

func INIT(filename string) {
	fmt.Println("Reading json file" + filename)
	RestauBuyers(filename)
}
