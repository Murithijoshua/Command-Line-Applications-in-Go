package main

import (
	"flag"
	"fmt"
	"joshua.com/todo"
	"os"
	"strings"
)

const TodoFile = "./todo.json"

func main() {
	TaskName := flag.String("task", "", "task-name")
	completed := flag.Int("complete", 0, "completed name")
	list := flag.Bool("list", false, "List all task")
	fmt.Println(TaskName, completed)
	l := &todo.List{}
	if err := l.Get(TodoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	case *list:
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *completed > 0:
		if err := l.Complete(*completed); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(TodoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)

		if err := l.Save(TodoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

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
