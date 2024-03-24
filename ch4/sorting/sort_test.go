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
		output := BubbleSort(test.array)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("got %v want %v", output, test.expected)
		}
	}

}
