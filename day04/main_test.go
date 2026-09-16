package main

import (
	"reflect"
	"testing"
)

func TestStatus(t *testing.T) {
	tests := []struct {
		name string
		item StatusProvider
		want string
	}{
		{
			name: "incomplete download",
			item: Download{},
			want: "downloading",
		},
		{
			name: "complete download",
			item: Download{Done: true},
			want: "downloaded",
		},
		{
			name: "incomplete task",
			item: &Task{},
			want: "active",
		},
		{
			name: "complete task",
			item: &Task{Done: true},
			want: "done",
		},
		{
			name: "missing task",
			item: (*Task)(nil),
			want: "unknown",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.item.Status()

			if got != test.want {
				t.Fatalf(
					"error: got %q, want %q",
					got,
					test.want,
				)
			}
		})
	}
}

func TestStatusText(t *testing.T) {
	tests := []struct {
		name string
		item StatusProvider
		want string
	}{
		{
			name: "first test",
			item: nil,
			want: "no object",
		},
		{
			name: "second test",
			item: (*Task)(nil),
			want: "unknown",
		},
		{
			name: "third test",
			item: &Task{Done: true},
			want: "done",
		},
		{
			name: "fourth test",
			item: Download{},
			want: "downloading",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := statusText(test.item)

			if got != test.want {
				t.Fatalf(
					"error: got %q, want %q",
					got,
					test.want,
				)
			}
		})
	}
}

func TestCollectStatuses(t *testing.T) {
	tests := []struct {
		name  string
		items []StatusProvider
		want  []string
	}{
		{
			name:  "empty slice",
			items: []StatusProvider{},
			want:  []string{},
		},
		{
			name: "mixed objects",
			items: []StatusProvider{
				&Task{Title: "Оплатить интернет"},
				Download{FileName: "photo.jpg", Done: true},
				nil,
				(*Task)(nil),
			},
			want: []string{
				"active",
				"downloaded",
				"no object",
				"unknown",
			},
		},
		{
			name:  "nil slice",
			items: nil,
			want:  []string{},
		},
		{
			name:  "one nil interface",
			items: []StatusProvider{nil},
			want: []string{
				"no object",
			},
		},
		{
			name:  "two nil interface",
			items: []StatusProvider{nil, nil},
			want: []string{
				"no object",
				"no object",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := collectStatuses(test.items)

			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf(
					"error: got %q, want %q",
					got,
					test.want,
				)
			}
		})
	}
}

func TestPayAndStatus(t *testing.T) {
	tests := []struct {
		name        string
		initialPaid bool
		want        string
	}{
		{
			name:        "paid",
			initialPaid: true,
			want:        "paid",
		},
		{
			name:        "pending",
			initialPaid: false,
			want:        "paid",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			payment := Payment{Paid: test.initialPaid}

			got := payAndStatus(&payment)

			if got != test.want {
				t.Errorf(
					"error: got %q, want %q",
					got,
					test.want,
				)
			}

			if !payment.Paid {
				t.Error("payment must be paid after payAndStatus")
			}
		})
	}
}
