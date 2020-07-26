package TopRestauBuyers

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"math"
	"os"
	"sort"
)

type topCustomers struct {
	CustomerID string
	Expenditure float64
}

type topRestaurants struct {
	Restaurant string
	Revenue float64
}

func findTopRestaurants (byteValue []byte, numRestau int) {
	resps := gjson.GetManyBytes(byteValue, "#.Restaurant","#.Price")
	restaurantsRevenue := make(map[string] float64)
	var restaurants []string
	var revenues []float64
	for i,resp := range resps {
		if i==0 {
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

	var TopRestaurants []topRestaurants
	for rests, revs := range restaurantsRevenue {
		TopRestaurants = append(TopRestaurants, topRestaurants{rests, revs})
	}
	sort.Slice(TopRestaurants, func(i, j int) bool {
		return TopRestaurants[i].Revenue > TopRestaurants[j].Revenue
	})
	fmt.Println("The top-5 Restaurants having following revenues are:")
	for ind := 0; ind < int(math.Min(float64(numRestau), float64(len(TopRestaurants)))); ind++ {
		fmt.Println(TopRestaurants[ind])
	}
}

func findTopBuyers (byteValue []byte, numCust int) {
	resps := gjson.GetManyBytes(byteValue, "#.CustomerID","#.Price")
	customersExpenditure := make(map[string] float64)
	var customers []string
	var expenditure []float64
	for i,resp := range resps {
		if i==0 {
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

	var TopCustomers []topCustomers
	for custs, expd := range customersExpenditure {
		TopCustomers = append(TopCustomers, topCustomers{custs, expd})
	}
	sort.Slice(TopCustomers, func(i, j int) bool {
		return TopCustomers[i].Expenditure > TopCustomers[j].Expenditure
	})
	fmt.Println("The top-5 Buyers having following expenditures are:")
	for ind := 0; ind < int(math.Min(float64(numCust), float64(len(TopCustomers)))); ind++ {
		fmt.Println(TopCustomers[ind])
	}
}

func RestauBuyers(filename string) {
	jsonFile, err := os.Open(filename)
	checkError(err)
	fmt.Println("Successfully Opened orders.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	findTopRestaurants(byteValue, 5)      // number of top restaurants you want
	findTopBuyers(byteValue, 5)			   // number of top customer who buys most
}

func INIT(filename string) {
	fmt.Println("Reading json file" + filename)
	RestauBuyers(filename)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
