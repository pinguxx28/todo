package src

import (
	"os"
	"fmt"
	"sort"
	"slices"
	"strings"
	"strconv"
	"text/tabwriter"
)

func sortTasks(tasks [][]string) {
	sort.Slice(tasks, func(i, j int) bool {
		// sort by priority if they are different
		priority1, _ := strconv.Atoi(tasks[i][1])
		priority2, _ := strconv.Atoi(tasks[j][1])
		if priority1 != priority2 {
			return priority1 < priority2
		}

		// sort by name if priorities are the same
		name1 := strings.ToLower(tasks[i][2])
		name2 := strings.ToLower(tasks[j][2])
		return strings.Compare(name1, name2) < 0
	})
}

func printTasks(tasks [][]string, flags []string) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.AlignRight)
	listAll := slices.Contains(flags, "a")
	listEverything := slices.Contains(flags, "e")

	listRange := 3
	if listEverything {
		listRange = 5
	}

	for _, task := range(tasks) {
		// don't list completed tasks
		if task[4] == "true" && !listAll {
			continue
		}

		for i := 0; i < listRange; i++ {
			fmt.Fprintf(writer, "%v\t", task[i])
		}

		fmt.Fprintln(writer, "")
	}

	writer.Flush()
}

func List(flags []string) {
	tasks := getAllTasks()
	sortTasks(tasks)
	printTasks(tasks, flags)
}
