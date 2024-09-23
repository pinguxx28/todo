package main

import (
	"os"
	"fmt"
	"strconv"
	"math/rand"
	"encoding/csv"
)

func add() {

	// get name
	// get priority
	// get desc?
	// get marking?

	// get all tasks
	file, err := os.OpenFile("tasks.csv", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Couldn't open file tasks.csv:", err)
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Couldn't read from tasks.csv:", err)
		os.Exit(1)
	}

	// create random id
	var id int

	for sameId := true; sameId; {
		id = rand.Intn(1000)
		sameId = false

		for _, task := range(tasks) {
			taskId, err := strconv.Atoi(task[0])
			if err != nil {
				fmt.Println("Couldn't convert string to int:",err)
				os.Exit(1)
			}

			if taskId == id {
				sameId = true
				break
			}
		}
	}

	fmt.Println(id)

	// add task with id, name, priority, desc and marking

	// set all tasks
}
