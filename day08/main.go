package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// readMessage читает файл и возвращает его содержимое как строку.
func readMessage(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open message: %w", err)
	}
	defer file.Close()
	text, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("read message: %w", err)
	}
	return string(text), nil
}

func main() {
	message, err := readMessage("day08/message.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("message file not found")
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(message)
}
