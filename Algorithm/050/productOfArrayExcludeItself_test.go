package _50

import (
	"reflect"
	"testing"
)

func TestProductOfArrayExcludeItself(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{
			[]int{1, 2, 3},
			[]int{6, 3, 2},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{3628800, 1814400, 1209600, 907200, 725760, 604800, 518400, 453600, 403200, 362880},
		},
		{
			[]int{1, 3, 5, 7, 9},
			[]int{945, 315, 189, 135, 105},
		},
	}

	for _, test := range tests {
		result := productOfArrayExcludeItself(test.arr)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("productOfArrayExcludeItself(%v) => %v,want %v", test.arr, result, test.want)
		}
	}
}
