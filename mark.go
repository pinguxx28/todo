package main

import (
	"os"
	"fmt"
)

func reverseBoolStr(str string) string {
	if str == "true" {
		return "false"
	} else if str == "false" {
		return "true"
	} else {
		fmt.Println("Str is neither true or false")
		os.Exit(1)
		return "" // avoid compiler error
	}
}

func mark() {
	tasks := getAllTasks()
	id := atoi(getInputNum("id", 0, 999))

	index := findTaskById(tasks, id)

	// false ->true, true ->false
	tasks[index][4] = reverseBoolStr(tasks[index][4])

	setAllTasks(tasks)
}
