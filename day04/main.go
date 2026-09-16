package main

import (
	"fmt"
)

type StatusProvider interface {
	Status() string
}

type Payable interface {
	MarkPaid()
	Status() string
}

type Download struct {
	FileName string
	Done     bool
}

type Task struct {
	Title string
	Done  bool
}

type Payment struct {
	Paid bool
}

func (p *Payment) MarkPaid() {
	p.Paid = true
}

func (p *Payment) Status() string {
	if p == nil {
		return "unknown"
	}
	if p.Paid {
		return "paid"
	}
	return "pending"
}

func (d Download) Status() string {
	if d.Done {
		return "downloaded"
	}
	return "downloading"
}

func (t *Task) Status() string {
	if t == nil {
		return "unknown"
	}
	if t.Done {
		return "done"
	}
	return "active"
}

func payAndStatus(item Payable) string {
	item.MarkPaid()
	return item.Status()
}

func printStatus(item StatusProvider) {
	fmt.Println(item.Status())
}

func statusText(item StatusProvider) string {
	if item == nil {
		return "no object"
	}
	return item.Status()
}

func collectStatuses(items []StatusProvider) []string {
	statuses := []string{}
	for _, item := range items {
		status := statusText(item)
		statuses = append(statuses, status)
	}
	return statuses
}

func main() {

	file1 := Download{
		FileName: "first file",
		Done:     true,
	}
	file2 := Download{
		FileName: "second file",
	}
	task1 := Task{
		Title: "first task",
		Done:  true,
	}
	task2 := Task{
		Title: "second task",
	}

	items := []StatusProvider{
		file1,
		file2,
		&task1,
		&task2,
	}

	for _, item := range items {
		printStatus(item)
	}

	download := Download{}
	task := Task{}

	var first StatusProvider = &download
	var second StatusProvider = &task

	download.Done = true
	task.Done = true

	fmt.Println(first.Status())
	fmt.Println(second.Status())

	selected := []StatusProvider{
		&Task{Title: "Оплатить интернет"},
		Download{FileName: "photo.jpg", Done: true},
		nil,
		(*Task)(nil),
	}

	fmt.Println(collectStatuses(selected))

	payment := Payment{}

	fmt.Println(payment.Status())
	fmt.Println(payAndStatus(&payment))
	fmt.Println(payment.Paid)
}
