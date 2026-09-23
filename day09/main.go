package main

import (
	"fmt"
	"io"
	"os"
)

// readMessage читает данные из io.Reader и возвращает их как строку.
func readMessage(r io.Reader) (string, error) {
	text, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("readMessage: %w", err)
	}
	return string(text), nil
}

func main() {
	file, err := os.Open("day08/message.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	text, err := readMessage(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
}
