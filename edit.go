package main

import (
	"os"
	"fmt"
)

func getEditedPart(partToEdit int) string {
	switch partToEdit {
	case 1:
		return getInputNum("New priority", 0, 10)
	case 2:
		return getInput("New name", true)
	case 3:
		return getInput("New desc", false)
	case 4:
		return getMarking()
	default:
		fmt.Println("PartToEdit is not in the range 1-4")
		os.Exit(1)
		return "" // shuts the compiler up
	}
}

func edit() {
	tasks := getAllTasks()
	id := atoi(getInputNum("id", 0, 999))
	index := findTaskById(tasks, id)

	fmt.Printf("0. Priority: %v\n", tasks[index][1])
	fmt.Printf("1. Name: %v\n", tasks[index][2])
	fmt.Printf("2. Desc: %v\n", tasks[index][3])
	fmt.Printf("3. Marking: %v\n", tasks[index][4])

	partToEdit := atoi(getInputNum("Field", 0, 3)) + 1
	editedPart := getEditedPart(partToEdit)

	tasks[index][partToEdit] = editedPart

	setAllTasks(tasks)
}
