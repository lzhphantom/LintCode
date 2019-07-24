package _65

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		arr1, arr2 []int
		want       interface{}
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{2, 3, 4, 5},
			3.5,
		},
		{
			[]int{1, 2, 3},
			[]int{4, 5},
			3,
		},
	}

	for _, test := range tests {
		result := medianOf2SortedArrays(test.arr1, test.arr2)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("medianOf2SortedArrays(%v, %v) => %v,want %v", test.arr1, test.arr2, result, test.want)
		}
	}
}
