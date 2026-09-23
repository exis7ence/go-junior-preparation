package main

import (
	"errors"
	"strings"
	"testing"
)

type errorReader struct {
	err error
}

type partialErrorReader struct {
	err error
}

func (e errorReader) Read(p []byte) (n int, err error) {
	return 0, e.err
}

func (part partialErrorReader) Read(p []byte) (n int, err error) {
	return copy(p, "Go"), part.err
}

// TestReadMessageSuccess проверяет, что текст возвращается без изменений.
func TestReadMessageSuccess(t *testing.T) {
	text := "я \nИлья"
	r := strings.NewReader(text)
	readText, err := readMessage(r)
	if err != nil {
		t.Fatalf("read message: %v", err)
	}
	if readText != text {
		t.Errorf("expected %q, got %q", text, readText)
	}
}

// TestReadMessageEmpty проверяет, что пустой источник возвращает пустую строку без изменений.
func TestReadMessageEmpty(t *testing.T) {
	text := ""
	r := strings.NewReader(text)
	readText, err := readMessage(r)
	if err != nil {
		t.Fatalf("readMessage: %v", err)
	}
	if readText != text {
		t.Errorf("expected empty, got %q", readText)
	}
}

// TestReadMessageError проверяет, что ошибка источника сохраняется в обёртке, а возвращённая строка пуста.
func TestReadMessageError(t *testing.T) {
	expectedErr := errors.New("read failed")
	reader := errorReader{
		err: expectedErr,
	}

	text, err := readMessage(reader)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
	if text != "" {
		t.Errorf("expected empty, got %q", text)
	}
}

// TestReadMessageTwice проверяет, что повторное чтение исчерпанного источника возвращает пустую строку без ошибки.
func TestReadMessageTwice(t *testing.T) {
	r := strings.NewReader("Go")
	text, err := readMessage(r)
	if err != nil {
		t.Fatalf("first readMessage: %v", err)
	}
	if text != "Go" {
		t.Errorf("first read: expected \"Go\", got %q", text)
	}
	textTwice, err := readMessage(r)
	if err != nil {
		t.Fatalf("second readMessage: %v", err)
	}
	if textTwice != "" {
		t.Errorf("second read: expected empty, got %q", textTwice)
	}
}

// TestReadMessagePartialError проверяет, что при ошибке чтения функция readMessage отбрасывает частичные данные и сохраняет причину ошибки.
func TestReadMessagePartialError(t *testing.T) {
	expectedErr := errors.New("connection lost")
	reader := partialErrorReader{
		err: expectedErr,
	}
	text, err := readMessage(reader)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
	if text != "" {
		t.Errorf("expected empty, got %q", text)
	}
}
