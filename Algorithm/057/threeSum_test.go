package _57

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		arr  []int
		want [][]int
	}{
		{
			[]int{2, 7, 11, 15},
			[][]int{},
		},
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, 0, 1},
				{-1, -1, 2},
			},
		},
	}

	for _, test := range tests {
		result := threeSum(test.arr)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("threeSum(%v) => %v,want %v", test.arr, result, test.want)
		}
	}
}
