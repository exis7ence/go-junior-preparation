package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func priorityName(priority int) (string, bool) {
	switch priority {
	case 1:
		return "низкий", true
	case 2:
		return "средний", true
	case 3:
		return "высокий", true
	default:
		return "", false
	}
}

func deadlineText(days int) string {
	if days < 0 {
		return "просрочено"
	}
	if days > 0 {
		return fmt.Sprintf("через %d дн.", days)
	}
	return "сегодня"
}

func printTask(title, days, priority string) {
	fmt.Printf("Задача: %s \nСрок: %s \nПриоритет %s\n", title, days, priority)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Некорректное количество аргументов")
		os.Exit(1)
	}
	title := strings.TrimSpace(os.Args[1])
	if title == "" {
		fmt.Println("Ошибка: пустой заголовок")
		os.Exit(1)
	}
	days, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Ошибка: Количество дней должно быть числом")
		os.Exit(1)
	}
	priority, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Ошибка: Приоритет должен быть числом")
		os.Exit(1)
	}
	priorityText, ok := priorityName(priority)
	if !ok {
		fmt.Println("Ошибка: Приоритет должен быть от 1 до 3")
		os.Exit(1)
	}
	printTask(title, deadlineText(days), priorityText)
}
