package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// TestLoadTaskMissingFile проверяет сохранение причины
// отсутствия файла после оборачивания.
func TestLoadTaskMissingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missing.json")

	task, err := loadTask(path)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}
	if (task != Task{}) {
		t.Fatalf("expected empty task, got %+v", task)
	}
}

// TestLoadTaskEmptyTitle проверяет сохранение причины
// пустого заголовка после оборачивания.
func TestLoadTaskEmptyTitle(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "task.json")

	raw := `{"title":"  "}`

	err := os.WriteFile(path, []byte(raw), 0o600)
	if err != nil {
		t.Fatalf("prepare file: %v", err)
	}

	task, err := loadTask(path)
	if !errors.Is(err, ErrEmptyTitle) {
		t.Fatalf("expected ErrEmptyTitle, got %v", err)
	}

	if (task != Task{}) {
		t.Fatalf("expected empty task, got %+v", task)
	}
}

// TestLoadTaskSuccess проверяет успешную
// загрузку задачи и очистку названия.
func TestLoadTaskSuccess(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "task.json")

	raw := `{"title":" Go ","done":true}`

	err := os.WriteFile(path, []byte(raw), 0o600)
	if err != nil {
		t.Fatalf("prepare file: %v", err)
	}

	task, err := loadTask(path)
	if err != nil {
		t.Fatalf("loadTask: %v", err)
	}

	want := Task{Title: "Go", Done: true}
	if task != want {
		t.Fatalf("expected %+v, got %+v", want, task)
	}
}

// TestLoadTaskInvalidType проверяет распознавание ошибки
// типа JSON после оборачивания.
func TestLoadTaskInvalidType(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "task.json")

	raw := `{"title":" Go ","done":"yes"}`

	err := os.WriteFile(path, []byte(raw), 0o600)
	if err != nil {
		t.Fatalf("prepare file: %v", err)
	}

	task, err := loadTask(path)
	var typeErr *json.UnmarshalTypeError
	if !errors.As(err, &typeErr) {
		t.Fatalf("expected *json.UnmarshalTypeError, got %v", err)
	}

	if typeErr.Value != "string" {
		t.Fatalf("expected Value %q, got %q", "string", typeErr.Value)
	}

	if typeErr.Field != "done" {
		t.Fatalf("expected field done, got %v", typeErr.Field)
	}

	want := Task{}
	if task != want {
		t.Fatalf("expected %+v, got %+v", want, task)
	}
}

// TestLoadTaskMissingFileIsNotTypeError проверяет, что ошибка
// отсутствующего файла не распознаётся как ошибка несовпадения типов JSON.
func TestLoadTaskMissingFileIsNotTypeError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missing.json")

	task, err := loadTask(path)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}

	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &typeErr) {
		t.Fatalf("unexpected JSON type error: %v", typeErr)
	}

	if typeErr != nil {
		t.Fatalf("expected nil, got %v", typeErr)
	}

	want := Task{}
	if task != want {
		t.Fatalf("expected %+v, got %+v", want, task)
	}
}
