package _62

import (
	"reflect"
	"testing"
)

func TestSearchInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{
			[]int{4, 5, 1, 2, 3},
			1,
			2,
		},
		{
			[]int{4, 5, 1, 2, 3},
			0,
			-1,
		},
		{
			[]int{},
			0,
			-1,
		},
	}

	for _, test := range tests {
		result := searchInRotatedSortedArray(test.arr, test.target)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("searchInRotatedSortedArray(%v, %d) => %d,want %d", test.arr, test.target, result, test.want)
		}
	}
}
