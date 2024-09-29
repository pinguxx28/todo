package main

import (
	"os"
	"fmt"
	"strings"
)

func getFlags(args []string) []string {
	if len(args) == 2 {
		return []string{}
	}

	var flag string
	for _, arg := range(os.Args) {
		if arg[0] == '-' {
			flag = arg[1:]
			break
		}
	}
	
	return strings.Split(flag, "")
}

func main() {
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println("Usage: todo [list|add|remove|mark|help]")
		os.Exit(1)
	}

	flags := getFlags(os.Args)

	switch os.Args[1] {
	case "list":
		list(flags)
	case "add":
		add()
	case "mark":
		mark()
	case "edit":
		edit()
	case "remove":
		remove()
	case "help":
		help()
	default:
		fmt.Println("Unkown command:", os.Args[1])
		fmt.Println("Usage: todo [list|add|remove|mark|help]")
		os.Exit(1)
	}
}
