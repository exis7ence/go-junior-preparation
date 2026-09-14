package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		sums    int
	}{
		{
			name:    "положительные числа",
			numbers: []int{1, 2, 3},
			sums:    6,
		},
		{
			name:    "отрицательные и положительные числа",
			numbers: []int{-1, 2, -3, 4},
			sums:    2,
		},
		{
			name:    "пустой слайс",
			numbers: []int{},
			sums:    0,
		},
		{
			name:    "отрицательные числа",
			numbers: []int{-5, -10},
			sums:    -15,
		},
		{
			name:    "один элемент положительный",
			numbers: []int{100},
			sums:    100,
		},
		{
			name:    "один элемент отрицательный",
			numbers: []int{-100},
			sums:    -100,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotSums := sum(test.numbers)

			if gotSums != test.sums {
				t.Errorf(
					"got %d, want sums %d",
					gotSums,
					test.sums,
				)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		max     int
		OK      bool
	}{
		{
			name:    "отрицательные и положительные числа",
			numbers: []int{1, -2, 3, -4},
			max:     3,
			OK:      true,
		},
		{
			name:    "пустой слайс",
			numbers: []int{},
			max:     0,
			OK:      false,
		},
		{
			name:    "отрицательные числа",
			numbers: []int{-5, -10},
			max:     -5,
			OK:      true,
		},
		{
			name:    "один элемент отрицательный",
			numbers: []int{-100},
			max:     -100,
			OK:      true,
		},
		{
			name:    "положительные числа",
			numbers: []int{3, 3, 3, 3},
			max:     3,
			OK:      true,
		},
		{
			name:    "один элемент положительный",
			numbers: []int{100},
			max:     100,
			OK:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotMax, gotOK := max(test.numbers)

			if gotOK != test.OK {
				t.Errorf(
					"gotOK %t, OK %t",
					gotOK,
					test.OK,
				)
			}
			if gotMax != test.max {
				t.Errorf(
					"gotMax %d, max %d",
					gotMax,
					test.max,
				)
			}
		})
	}
}

func TestAnalyze(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		min     int
		max     int
		avg     float64
	}{
		{
			name:    "один элемент положительный",
			numbers: []int{100},
			min:     100,
			max:     100,
			avg:     100,
		},
		{
			name:    "один элемент отрицательный",
			numbers: []int{-100},
			min:     -100,
			max:     -100,
			avg:     -100,
		},
		{
			name:    "несколько чисел",
			numbers: []int{10, 20, 30},
			min:     10,
			max:     30,
			avg:     20,
		},
		{
			name:    "отрицательные числа",
			numbers: []int{-10, -20, -30},
			min:     -30,
			max:     -10,
			avg:     -20,
		},
		{
			name:    "отрицательные и положительные числа",
			numbers: []int{-10, 20, -30, 40},
			min:     -30,
			max:     40,
			avg:     5,
		},
		{
			name:    "пустой слайс",
			numbers: []int{},
			min:     0,
			max:     0,
			avg:     0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotMin, gotMax, gotAvg := analyze(test.numbers)

			if gotMin != test.min {
				t.Errorf(
					"gotMin %d, min %d",
					gotMin,
					test.min,
				)
			}
			if gotMax != test.max {
				t.Errorf(
					"gotMax %d, max %d",
					gotMax,
					test.max,
				)
			}
			if gotAvg != test.avg {
				t.Errorf(
					"gotAvg %f, min %f",
					gotAvg,
					test.avg,
				)
			}
		})
	}
}

func TestCountNumbers(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    map[int]int
	}{
		{
			name:    "первый тест",
			numbers: []int{},
			want:    map[int]int{},
		},
		{
			name:    "второй тест",
			numbers: []int{1},
			want:    map[int]int{1: 1},
		},
		{
			name:    "третий тест",
			numbers: []int{1, 1, 2, 3, 2, 2},
			want:    map[int]int{1: 2, 2: 3, 3: 1},
		},
		{
			name:    "четвертый тест",
			numbers: []int{-1, -1, 5},
			want:    map[int]int{-1: 2, 5: 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := countNumbers(test.numbers)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(
					"got %v, want %v",
					got,
					test.want,
				)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		want    []int
	}{
		{
			name:    "первый тест",
			numbers: []int{},
			want:    []int{},
		},
		{
			name:    "второй тест",
			numbers: []int{1},
			want:    []int{1},
		},
		{
			name:    "третий тест",
			numbers: []int{1, 2, 1, 3, 2, 4},
			want:    []int{1, 2, 3, 4},
		},
		{
			name:    "четвертый тест",
			numbers: []int{5, 5, 3, 5, 3, 1},
			want:    []int{5, 3, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := unique(test.numbers)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(
					"got %v, want %v",
					got,
					test.want,
				)
			}
		})
	}
}

func TestRuneCount(t *testing.T) {
	tests := []struct {
		name string
		text string
		want int
	}{
		{
			name: "первый тест",
			text: "hello",
			want: 5,
		},
		{
			name: "второй тест",
			text: "Привет",
			want: 6,
		},
		{
			name: "третий тест",
			text: "",
			want: 0,
		},
		{
			name: "четвертый тест",
			text: "АБ",
			want: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := runeCount(test.text)

			if got != test.want {
				t.Errorf(
					"got %d, want %d",
					got,
					test.want,
				)
			}
		})
	}
}

func TestFirstRune(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		wantRune rune
		wantOK   bool
	}{
		{
			name:     "первый тест",
			text:     "Привет",
			wantRune: 'П',
			wantOK:   true,
		},
		{
			name:     "второй тест",
			text:     "Go",
			wantRune: 'G',
			wantOK:   true,
		},
		{
			name:     "третий тест",
			text:     "",
			wantRune: 0,
			wantOK:   false,
		},
		{
			name:     "четвертый тест",
			text:     "Я",
			wantRune: 'Я',
			wantOK:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotRune, gotOK := firstRune(test.text)

			if gotRune != test.wantRune {
				t.Errorf(
					"gotString %q, wantString %q",
					gotRune,
					test.wantRune,
				)
			}

			if gotOK != test.wantOK {
				t.Errorf(
					"gotOK %t, wantOK %t",
					gotOK,
					test.wantOK,
				)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "первый тест",
			text: "hello",
			want: "olleh",
		},
		{
			name: "второй тест",
			text: "Привет",
			want: "тевирП",
		},
		{
			name: "третий тест",
			text: "GoЯ",
			want: "ЯoG",
		},
		{
			name: "четвертый тест",
			text: "",
			want: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := reverseString(test.text)

			if got != test.want {
				t.Errorf(
					"got %s, want %s",
					got,
					test.want,
				)
			}
		})
	}
}
