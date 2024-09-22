package main

import (
	"os"
	"fmt"
	"log"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: todo [list|add|remove|mark|help]")
	}

	fmt.Println(os.Args[1])
}
