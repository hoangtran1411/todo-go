package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}

	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}

	isCompleted := (*todos)[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		(*todos)[index].CompletedAt = &completionTime
	}

	(*todos)[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}

	(*todos)[index].Title = title
	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *todos {
		completed := "❌"
		completedAt := ""
		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)

	}

	table.Render()
}
