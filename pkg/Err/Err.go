package Err

import "os"

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}
}
