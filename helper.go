package main

import (
	"os"
	"fmt"
	"strconv"
)

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Couldn't convert int(%v) to string %v\n", str, err)
		os.Exit(1)
	}

	return i
}
