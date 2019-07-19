package _58

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   [][]int
	}{
		{
			[]int{2, 7, 11, 15},
			3,
			[][]int{},
		},
		{
			[]int{1, 0, -1, 0, -2, 2},
			0,
			[][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			[]int{1, 0, -1, 0, -2, 2, 3},
			0,
			[][]int{
				{-2, -1, 0, 3},
				{-2, -1, 0, 3},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
	}

	for _, test := range tests {
		result := fourSum(test.arr, test.target)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("fourSum(%v, %d) => %v,want %v", test.arr, test.target, result, test.want)
		}
	}
}
