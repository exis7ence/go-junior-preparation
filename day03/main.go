package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type Task struct {
	Title    string
	Priority int
	Done     bool
}

func newTask(title string, priority int) (*Task, error) {
	trimmedTitle := strings.TrimSpace(title)
	if len(trimmedTitle) == 0 {
		return nil, errors.New("title cannot be empty")
	}
	switch priority {
	case 1, 2, 3:
		return &Task{
				Title:    trimmedTitle,
				Priority: priority,
			},
			nil
	default:
		return nil, errors.New("priority must be between 1 and 3")
	}
}

func (t Task) Status() string {
	if !t.Done {
		return "active"
	}
	return "done"
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) Rename(title string) error {
	trimmedTitle := strings.TrimSpace(title)
	if len(trimmedTitle) == 0 {
		return errors.New("title cannot be empty")
	}
	t.Title = trimmedTitle
	return nil
}

func (t *Task) SetPriority(priority int) error {
	switch priority {
	case 1, 2, 3:
		t.Priority = priority
		return nil
	default:
		return errors.New("priority must be between 1 and 3")
	}
}

func main() {
	task, err := newTask("Go", 2)
	if err != nil {
		log.Fatalf("ошибка при создании: %v", err)
	}

	fmt.Println(task.Title)
	fmt.Println(task.Status())

	task.Complete()
	err = task.Rename("    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(task.Title)
	fmt.Println(task.Status())

	err = task.SetPriority(3)
	if err != nil {
		fmt.Println(err)
	}
	err = task.SetPriority(10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(task.Priority)
}
