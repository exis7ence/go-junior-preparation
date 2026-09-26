package main

import (
	"errors"
	"fmt"
	"strings"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

type TaskList struct {
	Tasks  []Task
	LastID int
}

// Add добавляет задачу с очищенным названием и новым ID.
func (t *TaskList) Add(title string) (Task, error) {
	trimmedTitle := strings.TrimSpace(title)
	if trimmedTitle == "" {
		return Task{}, errors.New("title cannot be empty")
	}
	t.LastID++
	task := Task{
		ID:    t.LastID,
		Title: trimmedTitle,
	}
	t.Tasks = append(t.Tasks, task)
	return task, nil
}

// Complete помечает задачу с указанным ID как завершенную.
func (t *TaskList) Complete(id int) error {
	for i := range t.Tasks {
		if t.Tasks[i].ID == id {
			t.Tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("task %d not found", id)
}

func main() {
	var list TaskList
	_, err := list.Add(" Go ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = list.Add("Docker")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = list.Complete(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list.Tasks)
}
