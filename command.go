package main

import "flag"

type CmdFlags struct {
	Add    string `short:"a" long:"add" description:"Add a new todo"`
	List   bool   `short:"l" long:"list" description:"List all todos"`
	Toggle int    `short:"t" long:"toggle" description:"Toggle todo status"`
	Delete int    `short:"d" long:"delete" description:"Delete a todo"`
	Edit   string `short:"e" long:"edit" description:"Edit a todo"`
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle todo status")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo")
	flag.Parse()
	return &cf
}
