package main

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

// TestReadMessageMissingFile проверяет наличие ошибки отсутствия файла.
func TestReadMessageMissingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missing.txt")

	text, err := readMessage(path)
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected os.ErrNotExist, got %v", err)
	}
	if text != "" {
		t.Errorf("expected empty text, got %q", text)
	}
}

// TestReadMessageSuccess проверяет, что содержимое файла возвращается без изменений.
func TestReadMessageSuccess(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "task.txt")

	raw := `{"меня
	 зовут илья"}`

	err := os.WriteFile(path, []byte(raw), 0o600)
	if err != nil {
		t.Fatalf("prepare file: %v", err)
	}
	text, err := readMessage(path)
	if err != nil {
		t.Fatalf("readMessage: %v", err)
	}
	if text != raw {
		t.Errorf("expected %q, got %q", raw, text)
	}
}

// TestReadMessageEmptyFile проверяет, что чтение пустого файла возвращает пустую строку без ошибки.
func TestReadMessageEmptyFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "task.txt")

	raw := ""

	err := os.WriteFile(path, []byte(raw), 0o600)
	if err != nil {
		t.Fatalf("prepare file: %v", err)
	}
	text, err := readMessage(path)
	if err != nil {
		t.Fatalf("readMessage: %v", err)
	}
	if text != raw {
		t.Errorf("expected %q, got %q", raw, text)
	}
}

// TestReadMessagePathError проверяет, что через обёртку доступны операция и путь исходной файловой ошибки.
func TestReadMessagePathError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "notExist.txt")

	var pathError *os.PathError
	text, err := readMessage(path)
	if !errors.As(err, &pathError) {
		t.Fatalf("expected *os.PathError, got %v", err)
	}
	if pathError.Op != "open" {
		t.Errorf("Op: expected open, got %q", pathError.Op)
	}
	if pathError.Path != path {
		t.Errorf("Path: expected %q, got %q", path, pathError.Path)
	}
	if text != "" {
		t.Errorf("expected empty text, got %q", text)
	}
}
