package _64

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		arr1 []int
		m    int
		arr2 []int
		n    int
		want []int
	}{
		{
			[]int{1, 2, 3},
			3,
			[]int{4, 5},
			2,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 5},
			3,
			[]int{3, 4},
			2,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 4},
			3,
			[]int{3, 5},
			2,
			[]int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		before := append([]int{}, test.arr1...)
		mergeSortedArray(&test.arr1, test.m, test.arr2, test.n)
		if !reflect.DeepEqual(test.arr1, test.want) {
			t.Errorf("mergeSortedArray(%v, %d, %v, %d) => %v,want %v", before, test.m, test.arr2, test.n, test.arr1, test.want)
		}
	}
}
