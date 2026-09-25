package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

type errorWriter struct {
	err error
}

func (e errorWriter) Write(p []byte) (n int, err error) {
	return 0, e.err
}

// TestCopyMessageSuccess проверяет число скопированных байтов и содержимое буфера.
func TestCopyMessageSuccess(t *testing.T) {
	text := "ЯGo"
	reader := strings.NewReader(text)
	var buffer bytes.Buffer
	written, err := copyMessage(&buffer, reader)
	if err != nil {
		t.Errorf("copy failed: %v", err)
	}
	if written != int64(len(text)) {
		t.Errorf("expected %d, got %d", int64(len(text)), written)
	}
	if buffer.String() != text {
		t.Errorf("expected %q, got %q", text, buffer.String())
	}
}

// TestCopyMessageWriteError проверяет исходную ошибку записи и нулевое число скопированных байтов.
func TestCopyMessageWriteError(t *testing.T) {
	expectedErr := errors.New("copy failed")
	writer := errorWriter{
		err: expectedErr,
	}
	text := "Go"
	reader := strings.NewReader(text)
	written, err := copyMessage(&writer, reader)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
	if written != 0 {
		t.Errorf("expected 0, got %d", written)
	}
}
