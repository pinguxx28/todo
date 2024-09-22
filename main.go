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
		fmt.Println("listttt")
	case "add":
		fmt.Println("addddd")
	case "remove":
		fmt.Println("removeeee")
	case "mark":
		fmt.Println("markkkkk")
	case "help":
		fmt.Println("helppppp")
	default:
		fmt.Println("Unkown command:", os.Args[1])
		fmt.Println("Usage: todo [list|add|remove|mark|help]")
		os.Exit(1)
	}
}
