package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		array    []int
		expected []int
	}{
		{array: []int{42}, expected: []int{42}},
		{array: []int{3, 2}, expected: []int{2, 3}},
		{array: []int{40, 20, 30}, expected: []int{20, 30, 40}},
		{array: []int{1, 2, 3, 4}, expected: []int{1, 2, 3, 4}},
		{array: []int{4, 3, 2, 1}, expected: []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		BubbleSort(test.array)
		if !reflect.DeepEqual(test.array, test.expected) {
			t.Errorf("got %v want %v", test.array, test.expected)
		}
	}
}

func TestShellSort(t *testing.T) {
	tests := []struct {
		array    []int
		expected []int
	}{
		{array: []int{42}, expected: []int{42}},
		{array: []int{3, 2}, expected: []int{2, 3}},
		{array: []int{40, 20, 30}, expected: []int{20, 30, 40}},
		{array: []int{1, 2, 3, 4}, expected: []int{1, 2, 3, 4}},
		{array: []int{4, 3, 2, 1}, expected: []int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		ShellSort(test.array)
		if !reflect.DeepEqual(test.array, test.expected) {
			t.Errorf("got %v want %v", test.array, test.expected)
		}
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		array    []int
		expected []int
	}{
		{array: []int{42}, expected: []int{42}},
		{array: []int{3, 2}, expected: []int{2, 3}},
		{array: []int{40, 20, 30}, expected: []int{20, 30, 40}},
		{array: []int{1, 2, 3, 4}, expected: []int{1, 2, 3, 4}},
		{array: []int{4, 3, 2, 1}, expected: []int{1, 2, 3, 4}},
		{array: []int{4, 2, 3, 1}, expected: []int{1, 2, 3, 4}},
		{array: []int{1, 3, 2, 4}, expected: []int{1, 2, 3, 4}},
		{array: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		QuickSort(test.array)
		if !reflect.DeepEqual(test.array, test.expected) {
			t.Errorf("got %v want %v", test.array, test.expected)
		}
	}
}
