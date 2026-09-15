package main

import (
	"reflect"
	"testing"
)

func TestNewTask(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		priority int
		wantTask *Task
		wantErr  bool
	}{
		{
			name:     "first test",
			title:    "Go",
			priority: 2,
			wantTask: &Task{
				Title:    "Go",
				Priority: 2,
			},
		},
		{
			name:     "second test",
			title:    " Go ",
			priority: 1,
			wantTask: &Task{
				Title:    "Go",
				Priority: 1,
			},
		},
		{
			name:     "third test",
			title:    "SQL",
			priority: 3,
			wantTask: &Task{
				Title:    "SQL",
				Priority: 3,
			},
		},
		{
			name:     "fourth task",
			title:    "",
			priority: 2,
			wantErr:  true,
		},
		{
			name:     "fifth task",
			title:    "    ",
			priority: 2,
			wantErr:  true,
		},
		{
			name:     "sixth task",
			title:    "Go",
			priority: 0,
			wantErr:  true,
		},
		{
			name:     "seventh",
			title:    "Go",
			priority: 4,
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotTask, err := newTask(test.title, test.priority)
			if (err != nil) != test.wantErr {
				t.Fatalf(
					"error: got %v, wantErr %v",
					err,
					test.wantErr,
				)
			}
			if !reflect.DeepEqual(gotTask, test.wantTask) {
				t.Fatalf(
					"error: gotTask %v, wantTask %v",
					gotTask,
					test.wantTask,
				)
			}
		})
	}
}

func TestRename(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		wantTitle string
		wantErr   bool
	}{
		{
			name:      "first test",
			title:     "SQL",
			wantTitle: "SQL",
		},
		{
			name:      "second test",
			title:     " Docker ",
			wantTitle: "Docker",
		},
		{
			name:      "third test",
			title:     "",
			wantTitle: "Go",
			wantErr:   true,
		},
		{
			name:      "fourth test",
			title:     "   ",
			wantTitle: "Go",
			wantErr:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			task := Task{Title: "Go", Priority: 2}
			err := task.Rename(test.title)
			if (err != nil) != test.wantErr {
				t.Fatalf(
					"error: got %v, wantErr %v",
					err,
					test.wantErr,
				)
			}
			if task.Title != test.wantTitle {
				t.Fatalf(
					"error: gotTitle %q, wantTitle %q",
					task.Title,
					test.wantTitle,
				)
			}

		})
	}
}

func TestSetPriority(t *testing.T) {
	tests := []struct {
		name         string
		priority     int
		wantPriority int
		wantErr      bool
	}{
		{
			name:         "first test",
			priority:     1,
			wantPriority: 1,
		},
		{
			name:         "second test",
			priority:     2,
			wantPriority: 2,
		},
		{
			name:         "fourth test",
			priority:     3,
			wantPriority: 3,
		},
		{
			name:         "fifth test",
			priority:     0,
			wantPriority: 2,
			wantErr:      true,
		},
		{
			name:         "sixth test",
			priority:     4,
			wantPriority: 2,
			wantErr:      true,
		},
		{
			name:         "seventh test",
			priority:     -1,
			wantPriority: 2,
			wantErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			task := Task{Title: "Go", Priority: 2}
			err := task.SetPriority(test.priority)
			if (err != nil) != test.wantErr {
				t.Fatalf(
					"error: got %v, wantErr %v",
					err,
					test.wantErr,
				)
			}
			if task.Priority != test.wantPriority {
				t.Fatalf(
					"error: gotPririty %d, wantPriority %d",
					task.Priority,
					test.wantPriority,
				)
			}
		})
	}
}

func TestComplete(t *testing.T) {
	tests := []struct {
		name     string
		done     bool
		wantDone bool
	}{
		{
			name:     "first test",
			wantDone: true,
		},
		{
			name:     "second test",
			done:     true,
			wantDone: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			task := Task{Title: "Go", Priority: 2}
			task.Done = test.done

			task.Complete()

			gotDone := task.Done
			if gotDone != test.wantDone {
				t.Fatalf(
					"gotDone %t, wantDone %t",
					gotDone,
					test.wantDone,
				)
			}
		})
	}
}

func TestStatus(t *testing.T) {
	tests := []struct {
		name       string
		done       bool
		wantStatus string
	}{
		{
			name:       "first test",
			wantStatus: "active",
		},
		{
			name:       "second: status",
			done:       true,
			wantStatus: "done",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			task := Task{Title: "Go", Priority: 2}
			task.Done = test.done

			gotStatus := task.Status()

			if gotStatus != test.wantStatus {
				t.Fatalf(
					"gotStatus %q, wantStatus %q",
					gotStatus,
					test.wantStatus,
				)
			}

		})
	}
}
