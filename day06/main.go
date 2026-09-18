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

var ErrEmptyTitle error = errors.New("title cannot be empty")

// parseTask разбирает JSON, очищает название и отклоняет пустое название.
func parseTask(raw string) (Task, error) {
	var task Task

	err := json.Unmarshal([]byte(raw), &task)
	if err != nil {
		return Task{}, err
	}
	task.Title = strings.TrimSpace(task.Title)
	if task.Title == "" {
		return Task{}, ErrEmptyTitle
	}
	return task, nil
}

// loadTask читает файл, разбирает JSON и возвращает проверенную задачу.
func loadTask(path string) (Task, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Task{}, fmt.Errorf("read task file: %w", err)
	}

	task, err := parseTask(string(data))
	if err != nil {
		return Task{}, fmt.Errorf("parse task file: %w", err)
	}
	return task, nil
}

// saveTask сохраняет задачу в JSON-файл.
func saveTask(path string, task Task) error {
	save, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, save, 0o600)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	received, err := loadTask("day06/task.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("task file not found")
			return
		}
		if errors.Is(err, ErrEmptyTitle) {
			fmt.Println("task title must not be empty")
			return
		}
		var typeErr *json.UnmarshalTypeError
		if errors.As(err, &typeErr) {
			fmt.Printf("problem field %v, expected type %v\n", typeErr.Field, typeErr.Type)
			return
		}
		fmt.Println("error:", err)
		return
	}
	fmt.Println(received)
}
