package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// parseTask разбирает JSON, очищает название и отклоняет пустое название.
func parseTask(raw string) (Task, error) {
	var task Task

	err := json.Unmarshal([]byte(raw), &task)
	if err != nil {
		return Task{}, err
	}
	task.Title = strings.TrimSpace(task.Title)
	if task.Title == "" {
		return Task{}, errors.New("title cannot be empty")
	}
	return task, nil
}

// loadTask читает файл, разбирает JSON и возвращает проверенную задачу.
func loadTask(path string) (Task, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Task{}, err
	}

	return parseTask(string(data))
}

// saveTask сохраняет задачу в JSON-файл.
func saveTask(path string, task Task) error {
	save, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, save, 0600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	received, err := loadTask("day05/task.json")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(received)
}
