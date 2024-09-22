package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: todo [list|add|remove|mark|help]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		list()
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
