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
	for _, arg := range(args) {
		if arg[0] == '-' {
			flag = arg[1:]
			break
		}
	}
	
	return strings.Split(flag, "")
}

func getCommand(args []string) string {
	if len(args) == 2 {
		return args[1]
	}

	for i, arg := range(args) {
		if arg[0] != '-' && i > 0 {
			return arg
		}
	}

	return "help" // just return help as default
}

func main() {
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println("Too many or too few arguments")
		fmt.Println("Usage: todo [list|add|remove|mark|edit|help]")
		os.Exit(1)
	}

	flags := getFlags(os.Args)
	command := getCommand(os.Args)

	switch command {
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
		fmt.Println("Usage: todo [list|add|remove|mark|edit|help]")
		os.Exit(1)
	}
}
