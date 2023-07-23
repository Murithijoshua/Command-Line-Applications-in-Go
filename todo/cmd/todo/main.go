package main

import (
	"fmt"
	"joshua.com/todo"
	"os"
	"strings"
)

const TodoFile = "./todo.json"

func main() {
	l := &todo.List{}
	if err := l.Get(TodoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	// For no extra arguments, print the list
	case len(os.Args) == 1:
		// List current to do items
		for _, item := range *l {
			fmt.Println(item.Task)
		}
		// Concatenate all provided arguments with a space and
		// add to the list as an item
	default:
		// Concatenate all arguments with a space
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(TodoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
