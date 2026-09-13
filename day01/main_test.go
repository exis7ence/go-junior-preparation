package main

import "testing"

func TestPriorityName(t *testing.T) {
	tests := []struct {
		name     string
		priority int
		wantName string
		wantOK   bool
	}{
		{
			name:     "единица",
			priority: 1,
			wantName: "низкий",
			wantOK:   true,
		},
		{
			name:     "двойка",
			priority: 2,
			wantName: "средний",
			wantOK:   true,
		},
		{
			name:     "тройка",
			priority: 3,
			wantName: "высокий",
			wantOK:   true,
		},
		{
			name:     "четверка",
			priority: 4,
			wantName: "",
			wantOK:   false,
		},
		{
			name:     "нолик",
			priority: 0,
			wantName: "",
			wantOK:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotName, gotOK := priorityName(test.priority)

			if gotName != test.wantName {
				t.Errorf(
					"got %q, want %q",
					gotName,
					test.wantName,
				)
			}
			if gotOK != test.wantOK {
				t.Errorf(
					"got %v, want %v",
					gotOK,
					test.wantOK,
				)
			}
		})
	}
}

func TestDeadlineText(t *testing.T) {
	tests := []struct {
		name string
		days int
		want string
	}{
		{
			name: "просроченно на 2 дня",
			days: -2,
			want: "просрочено",
		},
		{
			name: "в день",
			days: 0,
			want: "сегодня",
		},
		{
			name: "1 день",
			days: 1,
			want: "через 1 дн.",
		},
		{
			name: "5 дней",
			days: 5,
			want: "через 5 дн.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := deadlineText(test.days)

			if got != test.want {
				t.Errorf(
					"got %q, want %q",
					got,
					test.want,
				)
			}
		})
	}
}
