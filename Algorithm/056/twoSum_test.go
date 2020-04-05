package _56

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{15, 2, 7, 11},
			9,
			[]int{1, 2},
		},
		{
			[]int{15, 3, 2, 7, 11},
			9,
			[]int{2, 3},
		},
		{
			[]int{1, 0, -1},
			0,
			[]int{0, 2},
		},
	}

	for _, test := range tests {
		result := twoSum(test.arr, test.target)
		if !reflect.DeepEqual(test.want, result) {
			t.Errorf("twoSum(%v, %d) => %v,want %v", test.arr, test.target, result, test.want)
		}
	}
}
