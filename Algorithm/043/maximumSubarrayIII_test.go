package _43

import (
	"reflect"
	"testing"
)

func TestMaximumSubarrayIII(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want int
	}{
		{
			[]int{1, 2, 3},
			1,
			6,
		},
		{
			[]int{-1, 4, -2, 3, -2, 3},
			2,
			8,
		},
		{
			[]int{-1, 4, -2, 3, -2, 3},
			3,
			10,
		},
		{
			[]int{-1, 4, -2, 3, -2, 3},
			7,
			0,
		},
		{
			[]int{-1, 4, -2, 3, -2, 3},
			6,
			5,
		},
		{
			[]int{-1, 4, -2, 3, -2, 3},
			5,
			7,
		},
	}

	for _, test := range tests {

		result := maximumSubarrayIII(test.arr, test.k)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("maximumSubarrayIII(%v, %d) => %d,want %d", test.arr, test.k, result, test.want)
		}
	}
}
