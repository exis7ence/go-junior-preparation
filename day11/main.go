package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// copyMessage копирует данные из src в dst и возвращает число скопированных байтов.
func copyMessage(dst io.Writer, src io.Reader) (int64, error) {
	written, err := io.Copy(dst, src)
	if err != nil {
		return written, fmt.Errorf("copy message: %w", err)
	}
	return written, nil
}

func main() {
	reader := strings.NewReader("ЯGo")
	var buffer bytes.Buffer
	written, err := copyMessage(&buffer, reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("written %d\nbuffer %q\n", written, buffer.String())
}
