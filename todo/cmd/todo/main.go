package main

import (
	"flag"
	"fmt"
	"joshua.com/todo"
	"os"
)

const TodoFile = "./todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(),
			"This tool was Developed as training by Joshua")
		fmt.Fprintln(flag.CommandLine.Output(), "Copyright 2023")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage Information")
		flag.PrintDefaults()
	}
	task := flag.String("task", "", "task-name")
	completed := flag.Int("complete", 0, "completed name")
	list := flag.Bool("list", false, "List all task")
	flag.Parse()

	l := &todo.List{}
	if err := l.Get(TodoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	switch {
	case *list:
		fmt.Print(l)
		//for _, item := range *l {
		//	if !item.Done {
		//		fmt.Println(item.Task)
		//	}
		//}
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

		// Concatenate all provided arguments with a space and
		// add to the list as an item
	default:
		// Concatenate all arguments with a space
		fmt.Fprintln(os.Stderr, "Invalid Option")
		os.Exit(1)
	}

}
