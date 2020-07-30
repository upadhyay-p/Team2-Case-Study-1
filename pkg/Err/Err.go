package Err

import "os"

func CheckError(err error) {
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}
