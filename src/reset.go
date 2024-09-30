package src

import (
	"os"
	"fmt"
)

func Reset() {
	file, err := os.OpenFile("tasks.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Couldn't open tasks.csv:", err)
		os.Exit(1)
	}
	file.Close()
}
