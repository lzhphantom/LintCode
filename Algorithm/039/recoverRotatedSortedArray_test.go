package _39

import (
	"reflect"
	"testing"
)

func TestRecoverRotatedSortedArray(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			[]int{2, 3, 4, 1},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{4, 1, 2, 3},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{6, 8, 9, 1, 2},
			[]int{1, 2, 6, 8, 9},
		},
	}

	for _, test := range tests {
		result := recoverRotatedSortedArray(test.arr)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("recoverRotatedSortedArray(%v) => %v, want %d", test.arr, result, test.want)
		}
	}
}
