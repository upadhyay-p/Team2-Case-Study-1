package main

import (
	"Apisgo"
	"CSV2JSON"
	"fmt"
)

func main() {
	fmt.Println("Task -1: Convert CSV to JSON format")
	fmt.Println("Task -2: See reports through API")

	outputFile := CSV2JSON.INIT("../assets/data")
	fmt.Println("Task-1 Done.\n")

	fmt.Println("Running Task-2")
	Apisgo.INIT(outputFile)

}