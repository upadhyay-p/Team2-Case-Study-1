package main

import (
	"Team2CaseStudy1/pkg/Apisgo"
	"Team2CaseStudy1/pkg/CSV2JSON"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		fmt.Println("Task -1: Converting CSV to JSON format...")
		outputFile := CSV2JSON.INIT(args[0])
		fmt.Println("Task-1 Done.")
		fmt.Println("Task -2: Initialising the API...")
		Apisgo.INIT(outputFile)
	} else {
		println("Please Give File name as an Argument!")
	}
	//fmt.Println("Task -1: Convert CSV to JSON format")
	//fmt.Println("Task -2: See reports through API")
	//
	//outputFile := CSV2JSON.INIT("../assets/data")
	//fmt.Println("Task-1 Done.\n")
	//
	//fmt.Println("Running Task-2")
	//Apisgo.INIT(outputFile)

}