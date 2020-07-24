package main

import (
	"CSV2JSON"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		CSV2JSON.INIT(args[0])
	} else {
		println("Please Give File name as an Argument!")
	}
}
