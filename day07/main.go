package main

import (
	"errors"
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}

// validateTitle отклоняет пустое название
// и название из пробелов.
func validateTitle(title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return &ValidationError{
			Field:   "title",
			Message: "cannot be empty",
		}
	}
	return nil
}

func main() {
	emptyTitle := ValidationError{
		Field:   "title",
		Message: "cannot be empty",
	}
	var err error
	err = &emptyTitle
	fmt.Println(err)

	var errType *ValidationError
	err = validateTitle("Go")
	if errors.As(err, &errType) {
		fmt.Println(errType.Field)
		fmt.Println(errType.Message)
	}
	fmt.Println(err)
}
