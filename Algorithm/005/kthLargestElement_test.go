package _05

import (
	"reflect"
	"testing"
)

func TestKthLargestElement(t *testing.T) {
	tests := []struct {
		n    int
		nums []int
		want int
	}{
		{1, []int{3, 2, 4, 5, 1}, 5},
		{2, []int{3, 8, 9, 1, 2, 7, 4}, 8},
	}

	for _, test := range tests {
		result := kthLargestElement(test.n, test.nums)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("kthLargestElement(%d,%v) => %d,want %d", test.n, test.nums, result, test.want)
		}
	}
}

func TestInsertSort(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 4, 3, 5}, []int{1, 3, 4, 5}},
	}

	for _, test := range tests {
		result := insertSort(test.nums)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("insertSort(%v) => %v,want %v", test.nums, result, test.want)
		}
	}
}
