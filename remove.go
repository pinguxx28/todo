package main

import (
	"os"
	"fmt"
	"slices"
	"encoding/csv"
)

func findTaskById(tasks [][]string, id int) int {
	for i, task := range(tasks) {
		taskId := atoi(task[0])
		if id == taskId {
			return i
		}
	}

	fmt.Printf("Couldn't find task with ID: %v\n", id)
	os.Exit(1)

	return -1 // Error if this line is not here
}

func setAllTasks(tasks [][]string) {
	file, err := os.OpenFile("tasks.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Couldn't open file tasks.csv:", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.WriteAll(tasks)

	if err := writer.Error(); err != nil {
		fmt.Println("Couldn't write CSV to tasks.csv:", err)
		os.Exit(1)
	}
}

func remove() { 
	tasks := getAllTasks()
	id := atoi(getInputNum("id", 0, 999))

	index := findTaskById(tasks, id)

	tasks = slices.Delete(tasks, index, index + 1)
	setAllTasks(tasks)
}
