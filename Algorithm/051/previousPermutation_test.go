package _51

import (
	"reflect"
	"testing"
)

func TestPerviousPermutation(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 3, 2, 3},
			[]int{1, 2, 3, 3},
		},
		{
			[]int{3, 1, 2, 4},
			[]int{2, 4, 3, 1},
		},
		{
			[]int{1, 2, 3, 4},
			[]int{4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		perviousPermutation(test.arr)
		if !reflect.DeepEqual(test.want, test.arr) {
			t.Errorf("perviousPermutation => %v,want %v", test.arr, test.want)
		}
	}
}
