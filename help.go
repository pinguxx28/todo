package main

import (
	"fmt"
)

func help() {
	fmt.Println("todo - a simple CLI todo list")
	fmt.Println("Usage: todo [list|add|remove|mark|edit|help]")
	fmt.Println("> list: lists all tasks")
	fmt.Println("> add: lets you add a new task")
	fmt.Println("> remove: lets you remove a task")
	fmt.Println("> mark: lets you reverse the marking of a task (completed or uncompleted)")
	fmt.Println("> edit: lets you edit a field in a task")
	fmt.Println("> help: displays this")
}
