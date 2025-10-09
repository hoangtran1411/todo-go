package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string `short:"a" long:"add" description:"Add a new todo"`
	List   bool   `short:"l" long:"list" description:"List all todos"`
	Toggle int    `short:"t" long:"toggle" description:"Toggle todo status"`
	Del    int    `short:"d" long:"delete" description:"Delete a todo"`
	Edit   string `short:"e" long:"edit" description:"Edit a todo"`
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle todo status")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo")
	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit. Please use index:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit.")
			os.Exit(1)

		}
		todos.Edit(index, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)

	case cf.Del != -1:
		todos.Delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
