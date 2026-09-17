package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestParseTask проверяет разбор JSON,
// очистку названия и отклонение некорректных данных.
func TestParseTask(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    Task
		wantErr bool
	}{
		{
			name:    "active task",
			raw:     `{"title":"Go","done":false}`,
			want:    Task{Title: "Go"},
			wantErr: false,
		},
		{
			name:    "done task",
			raw:     `{"title":"SQL","done":true}`,
			want:    Task{Title: "SQL", Done: true},
			wantErr: false,
		},
		{
			name:    "incorrect type done",
			raw:     `{"title":"Go","done":"false"}`,
			want:    Task{},
			wantErr: true,
		},
		{
			name:    "unclosed JSON",
			raw:     `{"title":"Go"`,
			want:    Task{},
			wantErr: true,
		},
		{
			name:    "missing done",
			raw:     `{"title":"Go"}`,
			want:    Task{Title: "Go"},
			wantErr: false,
		},
		{
			name:    "empty object",
			raw:     `{}`,
			want:    Task{},
			wantErr: true,
		},
		{
			name:    "empty title",
			raw:     `{"title":""}`,
			want:    Task{},
			wantErr: true,
		},
		{
			name:    "spaces at the edges",
			raw:     `{"title":"   Go   "}`,
			want:    Task{Title: "Go"},
			wantErr: false,
		},
		{
			name:    "only spaces",
			raw:     `{"title":"    "}`,
			want:    Task{},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseTask(test.raw)

			if (err != nil) != test.wantErr {
				t.Errorf(
					"parseTask: got error %v, wantErr %v",
					err,
					test.wantErr,
				)
			}

			if got != test.want {
				t.Errorf(
					"parseTask: got %+v, want %+v",
					got,
					test.want,
				)
			}
		})
	}
}

// TestLoadTask проверяет загрузку задачи из файла
// и удаление пробелов вокруг названия.
func TestLoadTask(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "task.json")

	raw := `{"title":"  Go  ","done":true}`

	err := os.WriteFile(path, []byte(raw), 0600)
	if err != nil {
		t.Fatalf("prepare file: %v", err)
	}

	got, err := loadTask(path)
	if err != nil {
		t.Fatalf("loadTask: %v", err)
	}

	want := Task{Title: "Go", Done: true}
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

// TestLoadTaskMissingFile проверяет ошибку
// и пустую задачу при отсутствии файла.
func TestLoadTaskMissingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "task.json")

	got, err := loadTask(path)
	if err == nil {
		t.Fatal("expected an error for a missing file.")
	}

	want := Task{}
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

// TestLoadTaskInvalidJSON проверяет ошибку
// и пустую задачу при некорректном JSON в файле.
func TestLoadTaskInvalidJSON(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "task.json")

	raw := `{"title":"Go", "done":true`

	err := os.WriteFile(path, []byte(raw), 0600)
	if err != nil {
		t.Fatalf(
			"prepare file: %v",
			err,
		)
	}

	got, err := loadTask(path)
	if err == nil {
		t.Fatal("expected an error for invalid JSON")
	}

	want := Task{}
	if got != want {
		t.Errorf(
			"got %+v, want %+v",
			got,
			want,
		)
	}
}

// TestSaveTask проверяет cоответствие
// записанного JSON исходной задачe.
func TestSaveTask(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "task.json")

	task := Task{
		Title: "Go",
		Done:  true,
	}

	err := saveTask(path, task)
	if err != nil {
		t.Fatalf("saveTask: unexpected error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read saved file: %v", err)
	}

	var got Task

	err = json.Unmarshal(data, &got)
	if err != nil {
		t.Fatalf("decode saved JSON: %v", err)
	}

	if task != got {
		t.Errorf("saved task: got %+v, want %+v", got, task)
	}
}

// TestSaveTaskWriteError проверяет возврат ошибки
// при попытке записи по пути к каталогу.
func TestSaveTaskWriteError(t *testing.T) {
	dir := t.TempDir()

	task := Task{
		Title: "Go",
	}

	err := saveTask(dir, task)
	if err == nil {
		t.Fatal("expected an error when writing to a directory")
	}
}

// TestSaveAndLoadTask проверяет сохранение значений
// полей после записи и загрузки.
func TestSaveAndLoadTask(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "task.json")

	task := Task{
		Title: "Go",
		Done:  true,
	}

	err := saveTask(path, task)
	if err != nil {
		t.Fatalf("saveTask: unexpected error: %v", err)
	}

	got, err := loadTask(path)
	if err != nil {
		t.Fatalf("loadTask: unexpected error: %v", err)
	}

	if got != task {
		t.Errorf("got %+v, want %+v", got, task)
	}
}
