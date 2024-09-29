package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"math/rand"
	"encoding/csv"
)

func getInput(prompt string, mandatory bool) string {
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%v of the task: ", prompt)
		value, err := input.ReadString('\n')

		if err != nil {
			fmt.Println("Please end your input with a newline character!")
			continue
		}

		value = strings.TrimSpace(value)

		if mandatory && value == "" {
			fmt.Println("Field is mandatory!")
			continue
		}

		return value
	}
}

func getInputNum(prompt string, low, high int) string {
	newPrompt := fmt.Sprintf("%v (%v-%v)", prompt, low, high)

	for {
		input := getInput(newPrompt, true)
		inputNum, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Couldn't convert string to int:", err)
			continue
		}

		if (inputNum >= low && inputNum <= high) {
			return input
		}

		fmt.Printf("Please enter a number between %v and %v\n", low, high)
	}
}

func getMarking() string {
	for {
		marking := getInput("Marking", false)

		if marking == "true" {
			return "true"
		} else {
			return "false"
		}
	}
}

func getAllTasks() [][]string {
	file, err := os.OpenFile("tasks.csv", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Couldn't open file tasks.csv:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Couldn't read from tasks.csv:", err)
		os.Exit(1)
	}

	return tasks
}

func createRandomId(tasks [][]string) string {
	var id int

	for sameId := true; sameId; {
		id = rand.Intn(1000)
		sameId = false

		for _, task := range(tasks) {
			taskId, err := strconv.Atoi(task[0])
			if err != nil {
				fmt.Println("Couldn't convert string to int:", err)
				os.Exit(1)
			}

			if taskId == id {
				sameId = true
				break
			}
		}
	}

	return strconv.Itoa(id)
}

func appendTask(task []string) {
	file, err := os.OpenFile("tasks.csv", os.O_WRONLY|os.O_APPEND, 0644) 
	if err != nil {
		fmt.Println("Couldn't open file tasks.csv:", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write(task)
	if err != nil {
		fmt.Println("Couldn't write to writer:", err)
		os.Exit(1)
	}

	writer.Flush()
	if writer.Error() != nil {
		fmt.Println("Couldn't write to the underlaying buffer:", err)
		os.Exit(1)
	}
}


func add() {
	name     := getInput("Name", true)
	priority := getInputNum("Priority", 0, 10)
	desc     := getInput("Desc", false)
	marking  := getMarking()

	tasks := getAllTasks()
	id := createRandomId(tasks)

	task := []string{id, priority, name, desc, marking}
	appendTask(task)
}
