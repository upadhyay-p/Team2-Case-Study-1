package main

import (
	"AvgPrice"
	"CSV2JSON"
	"bufio"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		CSV2JSON.INIT(args[0])
	} else {
		reader :=bufio.NewReader(os.Stdin)
		println("Enter the Json file name for report:")
		fname, err := reader.ReadString('\n')
		CSV2JSON.CheckErr(err)
		AvgPrice.INIT(strings.TrimSpace(fname))
	}
}
