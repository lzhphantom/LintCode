package _63

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   bool
	}{
		{
			[]int{},
			1,
			false,
		},
		{
			[]int{3, 4, 4, 5, 7, 0, 1, 2},
			4,
			true,
		},
	}

	for _, test := range tests {
		result := searchInRotatedSortedArrayII(test.arr, test.target)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("searchInRotatedSortedArrayII(%v, %d) => %v,want %v", test.arr, test.target, result, test.want)
		}
	}
}
