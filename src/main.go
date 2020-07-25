package main

import (
	"Apisgo"
	"CSV2JSON"
	"os"
)

// This is the main function executed at the beginning
func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		CSV2JSON.INIT(args[0])
		Apisgo.INIT(args[0] + ".json")
	} else {
		println("Please Give File name as an Argument!")
	}
}
