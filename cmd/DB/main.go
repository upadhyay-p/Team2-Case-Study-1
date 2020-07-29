package main

import (
	"Team2CaseStudy1/pkg/CSV2JSON"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		fmt.Println("Converting CSV to JSON...")
		CSV2JSON.INIT(args[0])
		fmt.Println("Conversion Complete!")
	} else {
		fmt.Println("Please give a CSV file as argument!")
	}
}
