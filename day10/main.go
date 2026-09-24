package main

import (
	"bytes"
	"fmt"
	"io"
)

func writeMessage(w io.Writer, message string) error {
	prepareMessage := []byte(message)
	_, err := w.Write(prepareMessage)
	if err != nil {
		return fmt.Errorf("write message: %w", err)
	}
	return nil
}

func main() {
	var buffer bytes.Buffer
	err := writeMessage(&buffer, "Learn recording in Go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buffer.String())
}
