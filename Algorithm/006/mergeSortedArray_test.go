package _06

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		arrA, arrB []int
		wantArr    []int
	}{
		{[]int{1, 2}, []int{1, 2}, []int{1, 1, 2, 2}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 4}, []int{2, 6}, []int{1, 2, 4, 6}},
	}

	for _, test := range tests {
		result := mergeSortedArray(test.arrA, test.arrB)
		if !reflect.DeepEqual(result, test.wantArr) {
			t.Errorf("mergeSortedArray(%v,%v) => %v,want %v", test.arrA, test.arrB, result, test.wantArr)
		}
	}
}
