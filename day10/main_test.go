package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

type errorWriter struct {
	err error
}

func (e errorWriter) Write(p []byte) (n int, err error) {
	return 0, e.err
}

// TestWriteMessageSuccess проверяет, что сообщение записалось без ошибок.
func TestWriteMessageSuccess(t *testing.T) {
	var buffer bytes.Buffer
	message := "Учусь записи в Go"
	err := writeMessage(&buffer, message)
	if err != nil {
		t.Fatalf("writeMessage: %v", err)
	}
	if buffer.String() != message {
		t.Errorf("expected %q, got %q", message, buffer.String())
	}
}

// TestWriteMessageError проверяет, что сообщение записывается в файл без изменений.
func TestWriteMessageError(t *testing.T) {
	expectedErr := errors.New("write failed")
	writer := errorWriter{
		err: expectedErr,
	}
	err := writeMessage(writer, "Учусь гошке")
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}

// TestWriteMessageToFile проверяет верную запись в файл.
func TestWriteMessageToFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "Go.txt")
	message := "Привет, Go\n"
	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("create file: %v", err)
	}
	defer file.Close()
	err = writeMessage(file, message)
	if err != nil {
		t.Fatalf("writeMessage: %v", err)
	}
	text, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read file: %v", err)
	}
	if string(text) != message {
		t.Errorf("expected %q, got %q", message, text)
	}
}
