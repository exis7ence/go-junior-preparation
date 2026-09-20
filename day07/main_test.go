package main

import (
	"errors"
	"fmt"
	"testing"
)

// TestValidateTitleEmpty проверяет, что название
// из пробелов вызывает *ValidationError с ожидаемыми полями.
func TestValidateTitleEmpty(t *testing.T) {
	var typeErr *ValidationError
	err := validateTitle("  ")
	if !errors.As(err, &typeErr) {
		t.Fatalf("expected *ValidationError, got %T %v", err, err)
	}
	if typeErr.Field != "title" {
		t.Errorf("Field: expected %q, got %q", "title", typeErr.Field)
	}

	if typeErr.Message != "cannot be empty" {
		t.Errorf("Message: expected %q, got %q", "cannot be empty", typeErr.Message)
	}
}

// TestValidateTitleSuccess проверяет, что корректное
// название не вызывает ошибок.
func TestValidateTitleSuccess(t *testing.T) {
	err := validateTitle("Go")
	if err != nil {
		t.Fatalf("expected nil, got %T %v", err, err)
	}
}

// TestValidateTitleWrappedError проверяет, что errors.As
// находит *ValidationError после оборачивания через %w.
func TestValidateTitleWrappedError(t *testing.T) {
	var typeErr *ValidationError
	err := validateTitle("  ")
	wrapErr := fmt.Errorf("create task: %w", err)
	if !errors.As(wrapErr, &typeErr) {
		t.Fatalf("expected *ValidationError, got %T %v", wrapErr, wrapErr)
	}
	if typeErr.Field != "title" {
		t.Errorf("Field: expected %q, got %q", "title", typeErr.Field)
	}
	if typeErr.Message != "cannot be empty" {
		t.Errorf("Message: expected %q, got %q", "cannot be empty", typeErr.Message)
	}
}

// TestValidateTitleTextError проверяет, что errors.As
// не находит *ValidationError после переноса ее текста в новую ошибку %v.
func TestValidateTitleTextError(t *testing.T) {
	var typeErr *ValidationError
	err := validateTitle("  ")
	wrapErr := fmt.Errorf("create task: %v", err)
	if errors.As(wrapErr, &typeErr) {
		t.Fatalf("unexpected *ValidationError: %v", typeErr)
	}
}

// TestValidationErrorTypedNil проверяет, что интерфейс error
// с типизированным nil-указателем не равен nil.
func TestValidationErrorTypedNil(t *testing.T) {
	var nilErr *ValidationError
	var err error = nilErr
	if nilErr != nil {
		t.Fatalf("expected nil, got %v", nilErr)
	}
	if err == nil {
		t.Fatal("expected non-nil error interface, got nil")
	}
}
