package main

import (
	"testing"
)

// TestTaskListAddSuccess проверяет, что Add возвращает и сохраняет задачу с очищенным названием.
func TestTaskListAddSuccess(t *testing.T) {
	var list TaskList
	task, err := list.Add(" Go ")
	if err != nil {
		t.Fatalf("add: %v", err)
	}
	if task.Done != false {
		t.Errorf("expected %t, got %t", false, task.Done)
	}
	if task.ID != 1 {
		t.Errorf("expected %d, got %d", 1, task.ID)
	}
	if task.Title != "Go" {
		t.Errorf("expected %q, got %q", "Go", task.Title)
	}
	if list.LastID != 1 {
		t.Errorf("expected %d, got %d", 1, list.LastID)
	}
	if len(list.Tasks) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(list.Tasks))
	}
	if list.Tasks[0] != task {
		t.Errorf("expected %+v, got %+v", task, list.Tasks[0])
	}
}

// TestTaskListAddEmpty проверяет, что пустое название возвращает ошибку и не меняет список.
func TestTaskListAddEmpty(t *testing.T) {
	list := TaskList{
		Tasks:  []Task{{Title: "Go", ID: 1}},
		LastID: 1,
	}
	_, err := list.Add("  ")
	if err == nil {
		t.Fatal("expected error for empty title")
	}
	if list.LastID != 1 {
		t.Errorf("expected %d, got %d", 1, list.LastID)
	}
	if len(list.Tasks) != 1 {
		t.Fatalf("expected %d, got %d", 1, len(list.Tasks))
	}
}

// TestCompleteSuccess проверяет, что задача с указанным ID становится завершенной.
func TestCompleteSuccess(t *testing.T) {
	list := TaskList{
		Tasks:  []Task{{Title: "Go", ID: 1}},
		LastID: 1,
	}
	err := list.Complete(1)
	if err != nil {
		t.Fatalf("complete failed: %v", err)
	}
	if list.Tasks[0].Done != true {
		t.Errorf("expected %t, got %t", true, list.Tasks[0].Done)
	}
}

// TestCompleteNotFound проверяет возврат ошибки на отсутствующий ID.
func TestCompleteNotFound(t *testing.T) {
	list := TaskList{
		Tasks:  []Task{{Title: "Go", ID: 1}},
		LastID: 1,
	}
	err := list.Complete(2)
	if err == nil {
		t.Fatal("expected error for task not found")
	}
	if list.Tasks[0].Done != false {
		t.Errorf("expected %t, got %t", false, list.Tasks[0].Done)
	}
}

// TestCompleteTwice проверяет, что повторный вызов Complete возвращает nil и оставляет Done равным true.
func TestCompleteTwice(t *testing.T) {
	list := TaskList{
		Tasks:  []Task{{Title: "Go", ID: 1}},
		LastID: 1,
	}
	err := list.Complete(1)
	if err != nil {
		t.Fatalf("complete failed: %v", err)
	}
	err = list.Complete(1)
	if err != nil {
		t.Fatalf("twice complete failed: %v", err)
	}
	if list.Tasks[0].Done != true {
		t.Errorf("expected %t, got %t", true, list.Tasks[0].Done)
	}
}
